package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type WarehouseRepositoryInterface interface {
	GetWarehouses() map[int]models.Warehouse
	GetWarehouseById(id int) (wh models.Warehouse, err error)
	DeleteWarehouse(id int) (err error)
	CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error)
}
