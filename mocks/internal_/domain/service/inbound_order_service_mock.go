package service_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockInboundOrderService struct {
	mock.Mock
}

func (m *MockInboundOrderService) CreateInboundOrder(inboundOrder *models.InboundOrder) error {
	args := m.Called(inboundOrder)
	return args.Error(0)
}
