package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	serviceImpl "github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	repository_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
)

func assertPurchaseOrderResponse(
	t *testing.T,
	responseErr error,
	expectedErr error,
	responseOutput interface{},
	expectedOutput interface{},
) {
	t.Helper()

	assert.True(t, errors.Is(responseErr, expectedErr), "Error mismatch")

	response, ok1 := responseOutput.(models.PurchaseOrder)
	expected, ok2 := expectedOutput.(models.PurchaseOrder)

	require.True(t, ok1 && ok2, "Output is not of type models.PurchaseOrder")

	assert.Equal(t, expected.BuyerID, response.BuyerID, "BuyerID mismatch")

	if responseErr == nil {
		assert.NotEmpty(t, response.OrderDate, "OrderDate should not be empty")
		assert.NotEmpty(t, response.TrackingCode, "TrackingCode should not be empty")

		_, err := time.Parse("2006-01-02 15:04:05", response.OrderDate)
		assert.NoError(t, err, "OrderDate is not in expected format")
	}
}

func TestPurchaseOrderService(t *testing.T) {
	tests := []struct {
		name             string
		repositoryMethod string
		serviceParams    []interface{}
		mockParams       []interface{}
		mockResponse     interface{}
		mockError        error
		expectedOutput   interface{}
		expectedErr      error
	}{
		{
			name:             "[GetById] OK Get purchase order by ID",
			repositoryMethod: "GetById",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse: entities.PurchaseOrderEntity{
				ID:            1,
				OrderNumber:   "OR0001",
				OrderDate:     "2025-02-15 04:08:52",
				TrackingCode:  "d2af4044-3c92-4d0b-9de6-9678c755f6c1",
				BuyerID:       1,
				CarrierID:     1,
				OrderStatusID: 1,
				WarehouseID:   1,
			},
			expectedOutput: models.PurchaseOrder{
				ID: 1,
				PurchaseOrderAttributes: models.PurchaseOrderAttributes{
					OrderNumber:  "OR0001",
					OrderDate:    "2025-02-15 04:08:52",
					TrackingCode: "d2af4044-3c92-4d0b-9de6-9678c755f6c1",
				},
				PurchaseOrderFKs: models.PurchaseOrderFKs{
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
		},
		{
			name:             "[GetById] Error PurchaseOrder not found",
			repositoryMethod: "GetById",
			serviceParams:    []interface{}{99},
			mockParams:       []interface{}{99},
			mockResponse:     entities.PurchaseOrderEntity{},
			mockError:        repository.ErrPurchaseOrderNotFoundWithID,
			expectedOutput:   models.PurchaseOrder{},
			expectedErr:      service.ErrPurchaseOrderDoesNotExist,
		},
		{
			name:             "[GetById] Error Unexpected error",
			repositoryMethod: "GetById",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse:     entities.PurchaseOrderEntity{},
			mockError:        repository.ErrPurchaseOrderUnexpectedError,
			expectedOutput:   models.PurchaseOrder{},
			expectedErr:      service.ErrPurchaseOrderUnexpectedError,
		},
		{
			name:             "[Create] OK Create new purchase order",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.PurchaseOrderAttributesFKs{
					PurchaseOrderAttributes: models.PurchaseOrderAttributes{},
					PurchaseOrderFKs: models.PurchaseOrderFKs{
						BuyerID:       1,
						CarrierID:     1,
						OrderStatusID: 1,
						WarehouseID:   1,
					},
				},
			},
			mockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     "2025-02-27 13:08:52",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			mockResponse: entities.PurchaseOrderEntity{
				ID:            2,
				OrderNumber:   "OR0002",
				OrderDate:     "2025-02-27 13:08:52",
				TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
				BuyerID:       1,
				CarrierID:     1,
				OrderStatusID: 1,
				WarehouseID:   1,
			},
			expectedOutput: models.PurchaseOrder{
				ID: 2,
				PurchaseOrderAttributes: models.PurchaseOrderAttributes{
					OrderNumber:  "OR0002",
					OrderDate:    "2025-02-27 13:08:52",
					TrackingCode: "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
				},
				PurchaseOrderFKs: models.PurchaseOrderFKs{
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
		},
		{
			name:             "[Create] Error FK ware house ID not valid",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.PurchaseOrderAttributesFKs{
					PurchaseOrderAttributes: models.PurchaseOrderAttributes{},
					PurchaseOrderFKs: models.PurchaseOrderFKs{
						BuyerID:       1,
						CarrierID:     1,
						OrderStatusID: 1,
						WarehouseID:   0,
					},
				},
			},
			mockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     "2025-02-27 13:08:52",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   0,
				},
			},
			mockResponse:   entities.PurchaseOrderEntity{},
			mockError:      repository.ErrPurchaseOrderFKWareHouseIdNotValid,
			expectedOutput: models.PurchaseOrder{},
			expectedErr:    service.ErrPurchaseOrderFKWareHouseIdNotValid,
		},
		{
			name:             "[Create] Error FK buyer ID not valid",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.PurchaseOrderAttributesFKs{
					PurchaseOrderAttributes: models.PurchaseOrderAttributes{},
					PurchaseOrderFKs: models.PurchaseOrderFKs{
						BuyerID:       0,
						CarrierID:     1,
						OrderStatusID: 1,
						WarehouseID:   1,
					},
				},
			},
			mockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     "2025-02-27 13:08:52",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       0,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			mockResponse:   entities.PurchaseOrderEntity{},
			mockError:      repository.ErrPurchaseOrderFKBuyerIdNotValid,
			expectedOutput: models.PurchaseOrder{},
			expectedErr:    service.ErrPurchaseOrderFKBuyerIdNotValid,
		},
		{
			name:             "[Create] Error FK order status ID not valid",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.PurchaseOrderAttributesFKs{
					PurchaseOrderAttributes: models.PurchaseOrderAttributes{},
					PurchaseOrderFKs: models.PurchaseOrderFKs{
						BuyerID:       1,
						CarrierID:     1,
						OrderStatusID: 0,
						WarehouseID:   1,
					},
				},
			},
			mockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     "2025-02-27 13:08:52",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 0,
					WarehouseID:   1,
				},
			},
			mockResponse:   entities.PurchaseOrderEntity{},
			mockError:      repository.ErrPurchaseOrderFKOrderStatusIdNotValid,
			expectedOutput: models.PurchaseOrder{},
			expectedErr:    service.ErrPurchaseOrderFKOrderStatusIdNotValid,
		},
		{
			name:             "[Create] Error FK carrier ID not valid",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.PurchaseOrderAttributesFKs{
					PurchaseOrderAttributes: models.PurchaseOrderAttributes{},
					PurchaseOrderFKs: models.PurchaseOrderFKs{
						BuyerID:       1,
						CarrierID:     0,
						OrderStatusID: 1,
						WarehouseID:   1,
					},
				},
			},
			mockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     "2025-02-27 13:08:52",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       1,
					CarrierID:     0,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			mockResponse:   entities.PurchaseOrderEntity{},
			mockError:      repository.ErrPurchaseOrderFKCarrierIdNotValid,
			expectedOutput: models.PurchaseOrder{},
			expectedErr:    service.ErrPurchaseOrderFKCarrierIdNotValid,
		},
		{
			name:             "[Create] Error Unexpected error",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.PurchaseOrderAttributesFKs{
					PurchaseOrderAttributes: models.PurchaseOrderAttributes{},
					PurchaseOrderFKs: models.PurchaseOrderFKs{
						BuyerID:       1,
						CarrierID:     1,
						OrderStatusID: 1,
						WarehouseID:   1,
					},
				},
			},
			mockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     "2025-02-27 13:08:52",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			mockResponse:   entities.PurchaseOrderEntity{},
			mockError:      service.ErrPurchaseOrderUnexpectedError,
			expectedOutput: models.PurchaseOrder{},
			expectedErr:    service.ErrPurchaseOrderUnexpectedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockPurchaseOrderRepository)

			if tt.mockResponse != nil {
				mockRepository.On(tt.repositoryMethod, tt.mockParams...).
					Return(tt.mockResponse, tt.mockError)
			} else {
				mockRepository.On(tt.repositoryMethod, tt.mockParams...).Return(tt.mockError)
			}

			purchaseOrderService := serviceImpl.NewPurchaseOrderService(mockRepository)

			switch tt.repositoryMethod {
			case "GetById":
				result, err := purchaseOrderService.GetById(tt.mockParams[0].(int))
				assertPurchaseOrderResponse(
					t,
					err,
					tt.expectedErr,
					result,
					tt.expectedOutput,
				)
			}

			mockRepository.AssertExpectations(t)
		})
	}
}
