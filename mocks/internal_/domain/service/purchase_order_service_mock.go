package service_mock

import (
	"time"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockPurchaseOrderService struct {
	mock.Mock
}

func (m *MockPurchaseOrderService) GetById(id int) (models.PurchaseOrder, error) {
	args := m.Called(id)
	return args.Get(0).(models.PurchaseOrder), args.Error(1)
}

func (m *MockPurchaseOrderService) Create(
	purchaseOrder models.PurchaseOrderAttributesFKs,
	utcNow time.Time,
	newUUID string,
) (models.PurchaseOrder, error) {
	args := m.Called(purchaseOrder)
	return args.Get(0).(models.PurchaseOrder), args.Error(1)
}
