package repository_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	"github.com/stretchr/testify/mock"
)

type MockOrderStatusRepository struct {
	mock.Mock
}

func (m *MockOrderStatusRepository) GetAll() ([]entities.OrderStatusEntity, error) {
	args := m.Called()
	return args.Get(0).([]entities.OrderStatusEntity), args.Error(1)
}
