package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type WarehouseRepositoryMock struct {
	mock.Mock
}

func (_w *WarehouseRepositoryMock) GetWarehouses() (map[int]models.Warehouse, error) {
	args := _w.Called()
	return args.Get(0).(map[int]models.Warehouse), args.Error(1)
}
