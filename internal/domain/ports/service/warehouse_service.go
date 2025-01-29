package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type WarehouseServiceInterface interface {
	GetWarehouses() map[int]models.Warehouse
}
