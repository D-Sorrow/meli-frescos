package repository_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
	"github.com/stretchr/testify/mock"
)

type MockBuyerRepository struct {
	mock.Mock
}

func (m *MockBuyerRepository) GetAll() ([]entities.BuyerEntity, error) {
	args := m.Called()
	return args.Get(0).([]entities.BuyerEntity), args.Error(1)
}

func (m *MockBuyerRepository) GetById(id int) (entities.BuyerEntity, error) {
	args := m.Called(id)
	return args.Get(0).(entities.BuyerEntity), args.Error(1)
}

func (m *MockBuyerRepository) Create(buyer entities.BuyerEntity) (entities.BuyerEntity, error) {
	args := m.Called(buyer)
	return args.Get(0).(entities.BuyerEntity), args.Error(1)
}

func (m *MockBuyerRepository) Patch(
	id int,
	buyerToPatch entities.BuyerEntity,
) (entities.BuyerEntity, error) {
	args := m.Called(id, buyerToPatch)
	return args.Get(0).(entities.BuyerEntity), args.Error(1)
}

func (m *MockBuyerRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBuyerRepository) GetReportPurchaseOrders(
	buyerID *int,
) ([]entities.ReportPurchaseOrdersEntity, error) {
	args := m.Called(buyerID)
	return args.Get(0).([]entities.ReportPurchaseOrdersEntity), args.Error(1)
}
