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
