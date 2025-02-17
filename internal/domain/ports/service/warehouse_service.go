package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type WarehouseServiceInterface interface {
	GetWarehouses() (map[int]models.Warehouse, error)
	GetWarehouseById(id int) (models.Warehouse, error)
	DeleteWarehouse(id int) error
	CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error)
	PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error)
}
