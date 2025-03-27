package repository_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type WarehouseRepositoryMock struct {
	mock.Mock
}

func NewWarehouseRepositoryMock() *WarehouseRepositoryMock {
	return &WarehouseRepositoryMock{}
}

func (_w *WarehouseRepositoryMock) GetWarehouses() (map[int]models.Warehouse, error) {
	args := _w.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[int]models.Warehouse), args.Error(1)
}

func (_w *WarehouseRepositoryMock) GetWarehouseById(id int) (models.Warehouse, error) {
	args := _w.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (_w *WarehouseRepositoryMock) DeleteWarehouse(id int) error {
	args := _w.Called(id)
	return args.Error(0)
}

func (_w *WarehouseRepositoryMock) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	args := _w.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (_w *WarehouseRepositoryMock) PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error) {
	args := _w.Called(id, data)
	return args.Get(0).(models.Warehouse), args.Error(1)
}
