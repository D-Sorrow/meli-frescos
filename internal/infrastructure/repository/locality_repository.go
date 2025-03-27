package repository

import (
	"database/sql"
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type LocalityRepository struct {
	db *sql.DB
}

func NewLocalityRepository(db *sql.DB) *LocalityRepository {
	return &LocalityRepository{db: db}
}

func (repo *LocalityRepository) CreateLocality(locality models.Locality) (models.Locality, error) {

	var provinceId int
	row := repo.db.QueryRow("select p.id from countries c join provinces p on c.id = p.id_country_fk where country_name = ? and province_name = ?", locality.CountryName, locality.ProvinceName)
	if err := row.Err(); err != nil {
		return models.Locality{}, error_management.HandleRepositoryError(repository.ErrProvinceNotFound, err)
	}
	err := row.Scan(&provinceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Locality{}, error_management.HandleRepositoryError(repository.ErrProvinceNotFound, err)
		}
		return models.Locality{}, error_management.HandleRepositoryError(repository.ErrLocalityRepositoryGeneric, err)
	}

	result, err := repo.db.Exec("INSERT INTO localities (locality_name,province_id,zip_code) values (?,?,?)", locality.Name, provinceId, locality.ZipCode)

	if err != nil {
		return models.Locality{}, error_management.HandleRepositoryError(error_management.HandleLocalityRepositoryError(err), err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return models.Locality{}, error_management.HandleRepositoryError(repository.ErrLocalityRepositoryGeneric, err)
	}

	locality.Id = int(lastInsertId)
	return locality, nil
}

func (repo *LocalityRepository) GetSellersByLocality(localityId int) (models.LocalitySellers, error) {
	var localitySellers models.LocalitySellers

	row := repo.db.QueryRow("SELECT l.id, l.zip_code, l.locality_name, count(s.id) as seller_count FROM localities l LEFT JOIN sellers s ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.zip_code, l.locality_name", localityId)
	if err := row.Err(); err != nil {
		return models.LocalitySellers{}, error_management.HandleRepositoryError(error_management.HandleLocalityRepositoryError(err), err)
	}

	err := row.Scan(&localitySellers.LocalityId, &localitySellers.ZipCode, &localitySellers.Name, &localitySellers.SellersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.LocalitySellers{}, error_management.HandleRepositoryError(repository.ErrLocalityNotFound, err)
		}
		return models.LocalitySellers{}, error_management.HandleRepositoryError(repository.ErrLocalityRepositoryGeneric, err)

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
		return nil, error_management.HandleRepositoryError(repository.ErrGetAllLocalities, err)
	}
	defer rows.Close()

	for rows.Next() {
		var carrierLocality models.LocalityCarriers
		err = rows.Scan(&carrierLocality.LocalityId,
			&carrierLocality.ZipCode,
			&carrierLocality.Name,
			&carrierLocality.CarriersCount)
		if err != nil {
			return nil, error_management.HandleRepositoryError(repository.ErrGetAllLocalities, err)
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
		return models.LocalityCarriers{}, error_management.HandleRepositoryError(repository.ErrLocalityNotFound, err)
	}

	err := row.Scan(&localityCarriers.LocalityId,
		&localityCarriers.ZipCode,
		&localityCarriers.Name,
		&localityCarriers.CarriersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.LocalityCarriers{}, error_management.HandleRepositoryError(repository.ErrLocalityNotFound, err)
		}
		return models.LocalityCarriers{}, error_management.HandleRepositoryError(repository.ErrLocalityRepositoryGeneric, err)
	}
	return localityCarriers, nil
}
