package service_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type LocalityServiceMock struct {
	mock.Mock
}

func NewLocalityServiceMock() *LocalityServiceMock {
	return &LocalityServiceMock{}
}

func (_l *LocalityServiceMock) CreateLocality(locality models.Locality) (l models.Locality, err error) {
	return
}

func (_l *LocalityServiceMock) GetSellersByLocality(localityId int) (ls models.LocalitySellers, err error) {
	return
}

func (_l *LocalityServiceMock) GetCarriersByAllLocalities() ([]models.LocalityCarriers, error) {
	args := _l.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.LocalityCarriers), args.Error(1)
}

func (_l *LocalityServiceMock) GetCarriersByLocality(id int) (models.LocalityCarriers, error) {
	args := _l.Called(id)
	return args.Get(0).(models.LocalityCarriers), args.Error(1)
}
