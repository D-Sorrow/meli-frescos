package repository_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

// MockInboundOrderRepository es un mock de InboundOrderRepository
type MockInboundOrderRepository struct {
	mock.Mock
}

// CreateInboundOrder es un método mock que simula la creación de un pedido de entrada
func (m *MockInboundOrderRepository) CreateInboundOrder(inboundOrder *models.InboundOrder) error {
	args := m.Called(inboundOrder)
	return args.Error(0)
}
