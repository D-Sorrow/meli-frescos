package service_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockBuyerService struct {
	mock.Mock
}

func (m *MockBuyerService) GetAll() ([]models.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]models.Buyer), args.Error(1)
}

func (m *MockBuyerService) GetById(id int) (models.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *MockBuyerService) Create(buyer models.BuyerAttributes) (models.Buyer, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *MockBuyerService) Patch(
	id int,
	buyerToPatch models.BuyerAttributes,
) (models.Buyer, error) {
	args := m.Called(id, buyerToPatch)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *MockBuyerService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBuyerService) GetReportPurchaseOrders(
	buyerID *int,
) ([]models.ReportPurchaseOrders, error) {
	args := m.Called(buyerID)
	return args.Get(0).([]models.ReportPurchaseOrders), args.Error(1)
}
