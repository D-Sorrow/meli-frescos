package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repoErros "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type WarehouseRepository struct {
	db map[int]models.Warehouse
}

func NewWarehouseRepository(db map[int]models.Warehouse) *WarehouseRepository {
	defaultDb := make(map[int]models.Warehouse)
	if db != nil {
		defaultDb = db
	}
	return &WarehouseRepository{db: defaultDb}
}

func (wr *WarehouseRepository) GetWarehouses() map[int]models.Warehouse {
	return wr.db
}

func (wr *WarehouseRepository) GetWarehouseById(id int) (wh models.Warehouse, err error) {
	for _, warehouse := range wr.db {
		if warehouse.Id == id {
			wh = warehouse
			return
		}
	}

	err = repoErros.ErrIdNotFound
	return
}

func (wr *WarehouseRepository) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	id := len(wr.db) + 1

	for _, wh := range wr.db {
		switch {
		case wh.Id == id:
			return models.Warehouse{}, repoErros.ErrIdDuplicate
		case wh.WarehouseCode == warehouse.WarehouseCode:
			return models.Warehouse{}, repoErros.ErrWarehouseCodeDuplicate
		}
	}

	warehouse.Id = id
	wr.db[id] = warehouse
	return wr.db[id], nil
}

func (wr *WarehouseRepository) PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error) {
	warehouse := wr.db[id]

	for key, value := range data {
		switch key {
		case "warehouse_code":
			for _, wh := range wr.db {
				if wh.WarehouseCode == value && wh.Id != id {
					return models.Warehouse{}, repoErros.ErrWarehouseCodeDuplicate
				}
			}

			warehouse.WarehouseCode = value.(string)
		case "address":
			warehouse.Address = value.(string)
		case "telephone":
			warehouse.Telephone = value.(string)
		case "minimun_capacity":
			warehouse.MinimunCapacity = int(value.(float64))
		case "minimun_temperature":
			warehouse.MinimunTemperature = int(value.(float64))
		}
	}

	wr.db[id] = warehouse
	return warehouse, nil
}

func (wr *WarehouseRepository) DeleteWarehouse(id int) (err error) {
	var warehouseExist bool
	for _, warehouse := range wr.db {
		if warehouse.Id == id {
			warehouseExist = true
			break
		}
	}

	if warehouseExist {
		delete(wr.db, id)
		return
	}
	err = repoErros.ErrIdNotFound
	return
}
