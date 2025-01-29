package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type WarehouseRepositoryInterface interface {
	GetWarehouses() map[int]models.Warehouse
}
