package service_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockOrderStatusService struct {
	mock.Mock
}

func (m *MockOrderStatusService) GetAll() ([]models.OrderStatus, error) {
	args := m.Called()
	return args.Get(0).([]models.OrderStatus), args.Error(1)
}
