package repository_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
	"github.com/stretchr/testify/mock"
)

type MockPurchaseOrderRepository struct {
	mock.Mock
}

func (m *MockPurchaseOrderRepository) GetById(id int) (entities.PurchaseOrderEntity, error) {
	args := m.Called(id)
	return args.Get(0).(entities.PurchaseOrderEntity), args.Error(1)
}

func (m *MockPurchaseOrderRepository) Create(
	purchaseOrder entities.PurchaseOrderEntity,
) (entities.PurchaseOrderEntity, error) {
	args := m.Called(purchaseOrder)
	return args.Get(0).(entities.PurchaseOrderEntity), args.Error(1)
}
