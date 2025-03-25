package repository_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type LocalityRepositoryMock struct {
	mock.Mock
}

func NewLocalityRepositoryMock() *LocalityRepositoryMock {
	return &LocalityRepositoryMock{}
}

func (_r *LocalityRepositoryMock) CreateLocality(locality models.Locality) (l models.Locality, err error) {
	return
}

func (_r *LocalityRepositoryMock) GetSellersByLocality(localityId int) (ls models.LocalitySellers, err error) {
	return
}

func (_r *LocalityRepositoryMock) GetCarriersByAllLocalities() ([]models.LocalityCarriers, error) {
	args := _r.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.LocalityCarriers), args.Error(1)
}

func (_r *LocalityRepositoryMock) GetCarriersByLocality(id int) (models.LocalityCarriers, error) {
	args := _r.Called(id)
	return args.Get(0).(models.LocalityCarriers), args.Error(1)
}
