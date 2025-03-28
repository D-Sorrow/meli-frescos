package service_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type WarehouseServiceMock struct {
	mock.Mock
}

func NewWarehouseServiceMock() *WarehouseServiceMock {
	return &WarehouseServiceMock{}
}

func (_w *WarehouseServiceMock) GetWarehouses() (map[int]models.Warehouse, error) {
	args := _w.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[int]models.Warehouse), args.Error(1)
}

func (_w *WarehouseServiceMock) GetWarehouseById(id int) (models.Warehouse, error) {
	args := _w.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (_w *WarehouseServiceMock) DeleteWarehouse(id int) error {
	args := _w.Called(id)
	return args.Error(0)
}

func (_w *WarehouseServiceMock) CreateWarehouse(
	warehouse models.Warehouse,
) (models.Warehouse, error) {
	args := _w.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (_w *WarehouseServiceMock) PatchWarehouse(
	id int,
	data map[string]interface{},
) (models.Warehouse, error) {
	args := _w.Called(id, data)
	return args.Get(0).(models.Warehouse), args.Error(1)
}
