package repository_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type CarrierRepositoryMock struct {
	mock.Mock
}

func NewCarrierRepositoryMock() *CarrierRepositoryMock {
	return &CarrierRepositoryMock{}
}

func (_r *CarrierRepositoryMock) GetAllCarriers() (carriers []models.Carrier, err error) {
	args := _r.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Carrier), args.Error(1)
}

func (_r *CarrierRepositoryMock) GetCarrierById(id int) (carrier models.Carrier, err error) {
	args := _r.Called(id)
	return args.Get(0).(models.Carrier), args.Error(1)
}

func (_r *CarrierRepositoryMock) CreateCarrier(carrier models.Carrier) (models.Carrier, error) {
	args := _r.Called(carrier)
	return args.Get(0).(models.Carrier), args.Error(1)
}
