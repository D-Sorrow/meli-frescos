package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type WarehouseRepositoryInterface interface {
	GetWarehouses() (map[int]models.Warehouse, error)
	GetWarehouseById(id int) (models.Warehouse, error)
	DeleteWarehouse(id int) error
	CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error)
	PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error)
}

var (
	ErrWarehouseNotFound                = errors.New("warehouse id not found")
	ErrWarehouseIdDuplicate             = errors.New("warehouse id duplicate")
	ErrWarehouseCodeDuplicate           = errors.New("warehouse code duplicate")
	ErrWarehouseDataBase                = errors.New("database error")
	ErrWarehouseLocalityIdNotFound      = errors.New("locality id not found")
	ErrWarehouseUpdateBySameData        = errors.New("enter different values to update")
	ErrWarehouseFKConstraintFail        = errors.New("foreign key constraint fails")
	ErrWarehouseLastInsertId            = errors.New("error with last insert id")
	ErrWarehouseGetUpdatedOrCreatedItem = errors.New("error getting updated or created warehouse")
)
