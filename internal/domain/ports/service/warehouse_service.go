package service

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type WarehouseServiceInterface interface {
	GetWarehouses() (map[int]models.Warehouse, error)
	GetWarehouseById(id int) (models.Warehouse, error)
	DeleteWarehouse(id int) error
	CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error)
	PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error)
}

var (
	ErrWarehouseServiceDefault          = errors.New("internal server error")
	ErrWarehouseNotFound                = errors.New("warehouse id not found")
	ErrWarehouseIdDuplicate             = errors.New("id already exists")
	ErrWarehouseCodeDuplicate           = errors.New("warehouse code already exists")
	ErrWarehouseLocalityIdNotFound      = errors.New("locality id not found")
	ErrWarehouseUpdateBySameData        = errors.New("enter different data to update")
	ErrWarehouseFKConstraintFail        = errors.New("foreign key constraint fails")
	ErrWarehouseGetUpdatedOrCreatedItem = errors.New(
		"the warehouse was created or updated but could not be displayed",
	)
)
