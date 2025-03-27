package repository_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (_m *ProductRepositoryMock) GetProducts() (map[int]models.Product, error) {
	args := _m.Called()
	return args.Get(0).(map[int]models.Product), args.Error(1)
}

func (_m *ProductRepositoryMock) GetProductByID(id int) (models.Product, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (_m *ProductRepositoryMock) UpdateProduct(id int, attributes map[string]any) error {
	args := _m.Called(id, attributes)
	return args.Error(0)
}

func (_m *ProductRepositoryMock) DeleteProduct(id int) error {
	args := _m.Called(id)
	return args.Error(0)
}

func (_m *ProductRepositoryMock) SaveProduct(productSave models.Product) error {
	args := _m.Called(productSave)
	return args.Error(0)
}
