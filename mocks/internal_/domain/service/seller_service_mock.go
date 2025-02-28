package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type SellerServiceMock struct {
	mock.Mock
}

func (_m *SellerServiceMock) GetSellers() (map[int]models.Seller, error) {
	args := _m.Called()
	return args.Get(0).(map[int]models.Seller), args.Error(1)
}

func (_m *SellerServiceMock) GetSellerById(id int) (models.Seller, error) {
	args := _m.Called(id)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (_m *SellerServiceMock) CreateSeller(seller models.Seller) (models.Seller, error) {
	args := _m.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (_m *SellerServiceMock) UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error) {
	args := _m.Called(id, seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (_m *SellerServiceMock) DeleteSeller(id int) error {
	args := _m.Called(id)
	return args.Error(0)
}
