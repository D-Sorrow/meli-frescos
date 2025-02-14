package repository

import (
	"database/sql"
	"errors"
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
	var localitySellers models.LocalitySellers

	row := repo.db.QueryRow("SELECT l.id, l.zip_code, l.locality_name, count(s.id) as seller_count FROM localities l LEFT JOIN sellers s ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.zip_code, l.locality_name", localityId)
	if err := row.Err(); err != nil {
		log.Print(err)
		return models.LocalitySellers{}, repository_errors.ErrLocalityNotFound
	}

	err := row.Scan(&localitySellers.LocalityId, &localitySellers.ZipCode, &localitySellers.Name, &localitySellers.SellersCount)
	if errors.Is(err, sql.ErrNoRows) {
		log.Print(err)
		return models.LocalitySellers{}, repository_errors.ErrLocalityNotFound
	}

	return localitySellers, nil

}

func (repo *LocalityRepository) GetCarriersByAllLocalities() ([]models.LocalityCarriers, error) {
	carriersByLocalities := []models.LocalityCarriers{}

	query := `SELECT l.id, l.zip_code, l.locality_name, COUNT(c.id) AS carriers_count
				FROM localities l LEFT JOIN carriers c ON l.id = c.locality_id
				GROUP BY l.id`

	rows, err := repo.db.Query(query)
	if err != nil {
		log.Print(err)
		return nil, repository_errors.ErrGetAllLocalities
	}
	defer rows.Close()

	for rows.Next() {
		var carrierLocality models.LocalityCarriers
		err = rows.Scan(&carrierLocality.LocalityId,
			&carrierLocality.ZipCode,
			&carrierLocality.Name,
			&carrierLocality.CarriersCount)
		if err != nil {
			return nil, repository_errors.ErrGetAllLocalities
		}
		carriersByLocalities = append(carriersByLocalities, carrierLocality)
	}
	return carriersByLocalities, nil
}

func (repo *LocalityRepository) GetCarriersByLocality(id int) (models.LocalityCarriers, error) {
	var localityCarriers models.LocalityCarriers

	query := `SELECT l.id, l.zip_code, l.locality_name, COUNT(c.id) AS carriers_count
				FROM localities l LEFT JOIN carriers c ON l.id = c.locality_id
				WHERE l.id = ? GROUP BY l.id`
	row := repo.db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		log.Print(err)
		return models.LocalityCarriers{}, repository_errors.ErrLocalityNotFound
	}

	err := row.Scan(&localityCarriers.LocalityId,
		&localityCarriers.ZipCode,
		&localityCarriers.Name,
		&localityCarriers.CarriersCount)
	if errors.Is(err, sql.ErrNoRows) {
		log.Print(err)
		return models.LocalityCarriers{}, repository_errors.ErrLocalityNotFound
	}

	return localityCarriers, nil
}
