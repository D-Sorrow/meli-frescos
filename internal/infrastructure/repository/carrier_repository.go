package repository

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repoErros "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
	"github.com/go-sql-driver/mysql"
)

type CarrierRepository struct {
	db *sql.DB
}

func NewCarrierRepository(db *sql.DB) *CarrierRepository {
	return &CarrierRepository{db: db}
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
		return models.Carrier{}, repoErros.ErrIdNotFound
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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				return models.Carrier{}, repoErros.ErrCarrierCidDuplicate
			case 1452:
				return models.Carrier{}, repoErros.ErrLocalityId
			case 1451:
				return models.Carrier{}, repoErros.ErrFKConstraintFail
			default:
				return models.Carrier{}, repoErros.ErrDataBase
			}

		}
	}

	id, _ := result.LastInsertId()
	newCarrier, _ := cr.GetCarrierById(int(id))
	return newCarrier, nil
}
