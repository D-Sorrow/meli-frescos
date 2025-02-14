package repository

import (
	"database/sql"
	"log"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repository_errors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
	"github.com/go-sql-driver/mysql"
)

type LocalityRepository struct {
	db *sql.DB
}

func NewLocalityRepository(db *sql.DB) *LocalityRepository {
	return &LocalityRepository{db: db}
}

func (repo *LocalityRepository) CreateLocality(locality models.Locality) (models.Locality, error) {

	result, err := repo.db.Exec("INSERT INTO localities (locality_name,province_id,zip_code) values (?,?,?)", locality.Name, locality.ProvinceId, locality.ZipCode)

	if err != nil {
		log.Print(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return models.Locality{}, repository_errors.ErrLocalityAlreadyExists
		}
		return models.Locality{}, repository_errors.ErrLocalityCannotBeCreated
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return models.Locality{}, repository_errors.ErrLocalityCannotBeCreated
	}

	locality.Id = int(lastInsertId)
	return locality, nil
}

func (repo *LocalityRepository) GetSellersByLocality(localityId int) (models.LocalitySellers, error) {
	//TODO
	return models.LocalitySellers{}, nil
}
