package repository

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	repoErrors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type CarrierRepository struct {
	db *sql.DB
}

func NewCarrierRepository(db *sql.DB) *CarrierRepository {
	return &CarrierRepository{db: db}
}

func (cr *CarrierRepository) GetAllCarriers() ([]models.Carrier, error) {
	var carriers []models.Carrier

	query := `SELECT id,
					cid,
					company_name,
					address,
					telephone,
					locality_id
				FROM carriers`
	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, repoErrors.HandleRepositoryError(repository.ErrCarrierDataBase, err)
	}
	defer rows.Close()

	for rows.Next() {
		var carrier models.Carrier
		err = rows.Scan(&carrier.Id,
			&carrier.Cid,
			&carrier.CompanyName,
			&carrier.Address,
			&carrier.Telephone,
			&carrier.LocalityId)
		if err != nil {
			return nil, repoErrors.HandleRepositoryError(repository.ErrCarrierDataBase, err)
		}
		carriers = append(carriers, carrier)
	}
	return carriers, nil
}

func (cr *CarrierRepository) GetCarrierById(id int) (models.Carrier, error) {
	var carrier models.Carrier
	query := `SELECT id,
					cid,
					company_name,
					address,
					telephone,
					locality_id
					FROM carriers
				WHERE id = ?`
	row := cr.db.QueryRow(query, id)
	err := row.Scan(&carrier.Id,
		&carrier.Cid,
		&carrier.CompanyName,
		&carrier.Address,
		&carrier.Telephone,
		&carrier.LocalityId)
	if err != nil {
		return models.Carrier{}, repoErrors.HandleRepositoryError(repository.ErrCarrierNotFound, err)
	}

	return carrier, nil
}

func (cr *CarrierRepository) CreateCarrier(carrier models.Carrier) (models.Carrier, error) {
	query := `INSERT INTO carriers 
				(cid,
				company_name,
				address,
				telephone,
				locality_id)
			VALUES (?, ?, ?, ?, ?);`
	result, err := cr.db.Exec(query,
		carrier.Cid,
		carrier.CompanyName,
		carrier.Address,
		carrier.Telephone,
		carrier.LocalityId)
	if err != nil {
		return models.Carrier{}, repoErrors.HandleRepositoryError(repoErrors.HandleCarrierRepositoryError(err), err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Carrier{}, repoErrors.HandleRepositoryError(repository.ErrCarrierGetUpdatedOrCreatedItem, err)
	}
	newCarrier, _ := cr.GetCarrierById(int(id))
	return newCarrier, nil
}
