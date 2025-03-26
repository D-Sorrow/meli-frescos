package service_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type CarrierServiceMock struct {
	mock.Mock
}

func NewCarryServiceMock() *CarrierServiceMock {
	return &CarrierServiceMock{}
}

func (_s *CarrierServiceMock) GetAllCarriers() ([]models.Carrier, error) {
	args := _s.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Carrier), args.Error(1)
}

func (_s *CarrierServiceMock) CreateCarrier(carrier models.Carrier) (models.Carrier, error) {
	args := _s.Called(carrier)
	return args.Get(0).(models.Carrier), args.Error(1)
}
