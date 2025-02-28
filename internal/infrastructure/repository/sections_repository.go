package repository

import (
	"database/sql"
	"log"

	er "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/mappers"
	"github.com/go-sql-driver/mysql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type SectionsRepository struct {
	db *sql.DB
}

func NewSectionsRepository(db *sql.DB) *SectionsRepository {
	return &SectionsRepository{db: db}
}

func (s SectionsRepository) GetSections() map[int]models.Sections {
	sectionsMap := make(map[int]models.Sections)
	rows, err := s.db.Query("SELECT id,section_number,current_temperature,minimum_temperature,current_capacity,minimum_capacity,maximum_capacity,warehouse_id,product_type_id FROM melifresh.sections")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var sections models.Sections
		err := rows.Scan(&sections.Id, &sections.Section_number, &sections.Current_temperature, &sections.Minimum_temperature, &sections.Current_capacity, &sections.Minimum_capacity, &sections.Maximum_capacity, &sections.Warehouse_id, &sections.Product_type_id)
		if err != nil {
			panic(err)
		}
		sectionsMap[sections.Id] = sections
	}
	return sectionsMap
}

func (s SectionsRepository) GetSectionsById(id int) (models.Sections, error) {
	var sections models.Sections

	row := s.db.QueryRow("SELECT id,section_number,current_temperature,minimum_temperature,current_capacity,minimum_capacity,maximum_capacity,warehouse_id,product_type_id FROM melifresh.sections WHERE  id = ?", id)
	err := row.Scan(&sections.Id, &sections.Section_number, &sections.Current_temperature, &sections.Minimum_temperature, &sections.Current_capacity, &sections.Minimum_capacity, &sections.Maximum_capacity, &sections.Warehouse_id, &sections.Product_type_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Sections{}, err
		}
	}
	return sections, nil

}

func (s SectionsRepository) SaveSections(sectionsSave models.Sections) error {

	sectionsEntity := mappers.ToSectionsEntity(sectionsSave)

	row := "INSERT INTO melifresh.sections" +
		"(section_number,current_temperature,minimum_temperature,current_capacity,minimum_capacity,maximum_capacity,warehouse_id,product_type_id)" +
		"VALUES (?,?,?,?,?,?,?,?)"
	_, err := s.db.Exec(row, sectionsEntity.Section_number, sectionsEntity.Current_temperature, sectionsEntity.Minimum_temperature, sectionsEntity.Current_temperature, sectionsEntity.Minimum_capacity, sectionsEntity.Maximum_capacity, sectionsEntity.Warehouse_id, sectionsEntity.Product_type_id)

	if err != nil {
		return err
	}
	return nil

}

func (s SectionsRepository) DeleteSections(id int) error {
	query := "DELETE FROM melifresh.sections WHERE id = ?"

	err := s.db.QueryRow(query, id)

	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func (s *SectionsRepository) UpdateSections(id int, sections models.SectionsPatch) (models.Sections, error) {
	value, err := s.GetSectionsById(id)

	if err != nil {
		return models.Sections{}, er.CodeSectionsIsExistErr
	}

	if sections.Warehouse_id != nil {
		value.Warehouse_id = *sections.Warehouse_id
	}
	if sections.Product_type_id != nil {
		value.Product_type_id = *sections.Product_type_id
	}
	if sections.Section_number != nil {
		value.Section_number = *sections.Section_number
	}
	if sections.Current_temperature != nil {
		value.Current_temperature = *sections.Current_temperature
	}
	if sections.Minimum_temperature != nil {
		value.Minimum_temperature = *sections.Minimum_temperature
	}
	if sections.Current_capacity != nil {
		value.Current_capacity = *sections.Current_capacity
	}
	if sections.Minimum_capacity != nil {
		value.Minimum_capacity = *sections.Minimum_capacity
	}
	if sections.Maximum_capacity != nil {
		value.Maximum_capacity = *sections.Maximum_capacity
	}

	_, err = s.db.Exec("UPDATE sections SET section_number=?,current_temperature=?,minimum_temperature=?,current_capacity =?,minimum_capacity =?,maximum_capacity =?,warehouse_id =?, product_type_id=?", value.Section_number, value.Current_temperature, value.Minimum_temperature, value.Current_capacity, value.Minimum_capacity, value.Maximum_capacity, value.Warehouse_id, value.Product_type_id, value.Id)

	if err != nil {
		log.Print(err)
		if mysqlError, ok := err.(*mysql.MySQLError); ok && mysqlError.Number == 1062 {
			return models.Sections{}, er.CodeSectionsIsExistErr
		}
		return models.Sections{}, er.CodeSectionsNotUpdate
	}

	return value, nil
}
