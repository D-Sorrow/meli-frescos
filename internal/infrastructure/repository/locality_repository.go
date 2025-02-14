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
