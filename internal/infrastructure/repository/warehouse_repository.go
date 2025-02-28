package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	repoErrors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type WarehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

func (wr *WarehouseRepository) GetWarehouses() (map[int]models.Warehouse, error) {
	warehousesMap := make(map[int]models.Warehouse)

	query := `SELECT id,
					warehouse_code,
					address,
					telephone,
					minimum_capacity,
					minimum_temperature,
					locality_id
				FROM warehouses`
	rows, err := wr.db.Query(query)
	if err != nil {
		return nil, repoErrors.HandleRepositoryError(repository.ErrWarehouseDataBase, err)
	}
	defer rows.Close()

	for rows.Next() {
		var warehouse models.Warehouse
		err = rows.Scan(&warehouse.Id,
			&warehouse.WarehouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimunCapacity,
			&warehouse.MinimunTemperature,
			&warehouse.LocalityId)
		if err != nil {
			return nil, repoErrors.HandleRepositoryError(repository.ErrWarehouseDataBase, err)
		}
		warehousesMap[warehouse.Id] = warehouse
	}
	return warehousesMap, nil
}

func (wr *WarehouseRepository) GetWarehouseById(id int) (models.Warehouse, error) {
	var warehouse models.Warehouse
	query := `SELECT id,
					warehouse_code,
					address,
					telephone,
					minimum_capacity,
					minimum_temperature,
					locality_id
				FROM warehouses
				WHERE id = ?`
	row := wr.db.QueryRow(query, id)
	err := row.Scan(&warehouse.Id,
		&warehouse.WarehouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimunCapacity,
		&warehouse.MinimunTemperature,
		&warehouse.LocalityId)
	if err != nil {
		return warehouse, repository.ErrWarehouseNotFound
	}

	return warehouse, nil
}

func (wr *WarehouseRepository) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	query := `INSERT INTO warehouses 
				(warehouse_code,
				address,
				telephone,
				minimum_capacity,
				minimum_temperature,
				locality_id)
			VALUES (?, ?, ?, ?, ?, ?);`
	result, err := wr.db.Exec(query,
		warehouse.WarehouseCode,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.MinimunCapacity,
		warehouse.MinimunTemperature,
		warehouse.LocalityId)
	if err != nil {
		return models.Warehouse{}, repoErrors.HandleWarehouseRepositoryError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Warehouse{}, repository.ErrWarehouseLastInsertId
	}
	NewWarehouse, err := wr.GetWarehouseById(int(id))
	if err != nil {
		return models.Warehouse{}, repository.ErrWarehouseGetUpdatedOrCreatedItem
	}
	return NewWarehouse, nil

}

func (wr *WarehouseRepository) PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error) {

	idExists, err := wr.verifyIdExist(id)
	if err != nil {
		return models.Warehouse{}, repoErrors.HandleRepositoryError(repository.ErrWarehouseDataBase, err)
	}

	if !idExists {
		return models.Warehouse{}, repository.ErrWarehouseNotFound
	}

	// query parts
	setClauses := make([]string, 0)
	args := make([]interface{}, 0)

	// get clauses and arguments to build the query
	for column, value := range data {
		switch v := value.(type) {
		case string:
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
			args = append(args, v)
		case int:
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
			args = append(args, v)
		case float64:
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
			args = append(args, int(v))
		}
	}
	fmt.Println(setClauses)
	fmt.Println(args...)

	query := fmt.Sprintf("UPDATE warehouses SET %s WHERE id = ?", strings.Join(setClauses, ","))
	args = append(args, id)

	result, err := wr.db.Exec(query, args...)
	if err != nil {
		return models.Warehouse{}, repoErrors.HandleRepositoryError(repoErrors.HandleWarehouseRepositoryError(err), err)
	}

	rowsAfected, err := result.RowsAffected()
	if err != nil {
		return models.Warehouse{}, repoErrors.HandleRepositoryError(repository.ErrWarehouseDataBase, err)
	}
	if rowsAfected == 0 {
		return models.Warehouse{}, repository.ErrWarehouseUpdateBySameData
	}

	UpdatedWarehouse, err := wr.GetWarehouseById(id)
	if err != nil {
		return models.Warehouse{}, repoErrors.HandleRepositoryError(repository.ErrWarehouseGetUpdatedOrCreatedItem, err)
	}
	return UpdatedWarehouse, nil

}

func (wr *WarehouseRepository) DeleteWarehouse(id int) error {
	idExists, err := wr.verifyIdExist(id)
	if err != nil {
		return repoErrors.HandleRepositoryError(repository.ErrWarehouseDataBase, err)
	}

	if !idExists {
		return repository.ErrWarehouseNotFound
	}

	query := "DELETE FROM warehouses WHERE id = ?"

	_, err = wr.db.Exec(query, id)
	if err != nil {
		return repoErrors.HandleRepositoryError(repoErrors.HandleWarehouseRepositoryError(err), err)
	}

	return nil
}

func (wh *WarehouseRepository) verifyIdExist(id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM warehouses WHERE id = ?)"

	err := wh.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
