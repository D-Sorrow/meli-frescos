package repository

import (
	"fmt"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
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

	return wh, fmt.Errorf("id does not exists")
}
