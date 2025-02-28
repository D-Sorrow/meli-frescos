package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (_m *ProductServiceMock) GetProducts() (map[int]models.Product, error) {
	args := _m.Called()
	return args.Get(0).(map[int]models.Product), args.Error(1)
}

func (_m *ProductServiceMock) GetProductByID(id int) (models.Product, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (_m *ProductServiceMock) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	args := _m.Called(id, attributes)
	return args.Get(0).(models.Product), args.Error(1)
}

func (_m *ProductServiceMock) DeleteProduct(id int) error {
	args := _m.Called(id)
	return args.Error(0)
}

func (_m *ProductServiceMock) SaveProduct(productSave models.Product) error {
	args := _m.Called(productSave)
	return args.Error(0)
}
