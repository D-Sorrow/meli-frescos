package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type WarehouseServiceInterface interface {
	GetWarehouses() map[int]models.Warehouse
	GetWarehouseById(id int) (wh models.Warehouse, err error)
	DeleteWarehouse(id int) error
}
