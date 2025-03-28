package service_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockInboundOrderService struct {
	mock.Mock
}

func (m *MockInboundOrderService) CreateInboundOrder(inboundOrder *models.InboundOrder) error {
	args := m.Called(inboundOrder)
	return args.Error(0)
}
