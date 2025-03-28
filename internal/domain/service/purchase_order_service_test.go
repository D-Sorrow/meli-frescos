package service_test

import (
	"testing"
	"time"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	serviceImpl "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/helpers"
	repository_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/infrastructure/repository"
)

func switchPurchaseOrderTest(
	t *testing.T,
	test helpers.ServiceTestStruct,
	purchaseOrderService *serviceImpl.PurchaseOrderService,
	testDatetime string,
	testUUID string,
) {
	t.Helper()

	switch test.RepositoryMethod {
	case "GetById":
		result, err := purchaseOrderService.GetById(test.MockParams[0].(int))
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	case "Create":
		param, ok := test.ServiceParams[0].(models.PurchaseOrderAttributesFKs)
		if !ok {
			t.Fatalf(
				"Invalid parameter type for Create method. Expected models.PurchaseOrderAttributesFKs, got %T",
				test.ServiceParams[0],
			)
		}

		dt, err := time.Parse("2006-01-02 15:04:05", testDatetime)
		if err != nil {
			t.Fatalf("Failed to parse date: %v", err)
		}

		result, err := purchaseOrderService.Create(
			param,
			dt.UTC(),
			testUUID,
		)
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	}
}

func assertPurchaseOrderService(
	t *testing.T,
	test helpers.ServiceTestStruct,
	mockRepository *repository_mock.MockPurchaseOrderRepository,
	testDatetime string,
	testUUID string,
) {
	t.Helper()

	purchaseOrderService := serviceImpl.NewPurchaseOrderService(mockRepository)
	switchPurchaseOrderTest(t, test, purchaseOrderService, testDatetime, testUUID)
}

func TestPurchaseOrderGetByIdService(t *testing.T) {
	const testDatetime string = "2025-02-27 13:08:52"
	const testUUID string = "d07a0493-0882-4702-b3aa-2dfe9ae108fe"
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[GetById] OK Get purchase order by ID",
			RepositoryMethod: "GetById",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse: entities.PurchaseOrderEntity{
				ID:            1,
				OrderNumber:   "OR0001",
				OrderDate:     testDatetime,
				TrackingCode:  testUUID,
				BuyerID:       1,
				CarrierID:     1,
				OrderStatusID: 1,
				WarehouseID:   1,
			},
			ExpectedOutput: models.PurchaseOrder{
				ID: 1,
				PurchaseOrderAttributes: models.PurchaseOrderAttributes{
					OrderNumber:  "OR0001",
					OrderDate:    testDatetime,
					TrackingCode: testUUID,
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
			Name:             "[GetById] Error PurchaseOrder not found",
			RepositoryMethod: "GetById",
			ServiceParams:    []interface{}{99},
			MockParams:       []interface{}{99},
			MockResponse:     entities.PurchaseOrderEntity{},
			MockError:        repository.ErrPurchaseOrderNotFoundWithID,
			ExpectedOutput:   models.PurchaseOrder{},
			ExpectedErr:      service.ErrPurchaseOrderDoesNotExist,
		},
		{
			Name:             "[GetById] Error Unexpected error",
			RepositoryMethod: "GetById",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse:     entities.PurchaseOrderEntity{},
			MockError:        repository.ErrPurchaseOrderUnexpectedError,
			ExpectedOutput:   models.PurchaseOrder{},
			ExpectedErr:      service.ErrPurchaseOrderUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockPurchaseOrderRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertPurchaseOrderService(t, test, mockRepository, testDatetime, testUUID)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestPurchaseOrderCreateService(t *testing.T) {
	const testDatetime string = "2025-02-27 13:08:52"
	const testUUID string = "d07a0493-0882-4702-b3aa-2dfe9ae108fe"
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[Create] OK Create new purchase order",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
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
			MockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     testDatetime,
					TrackingCode:  testUUID,
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			MockResponse: entities.PurchaseOrderEntity{
				ID:            2,
				OrderNumber:   "OR0002",
				OrderDate:     testDatetime,
				TrackingCode:  testUUID,
				BuyerID:       1,
				CarrierID:     1,
				OrderStatusID: 1,
				WarehouseID:   1,
			},
			ExpectedOutput: models.PurchaseOrder{
				ID: 2,
				PurchaseOrderAttributes: models.PurchaseOrderAttributes{
					OrderNumber:  "OR0002",
					OrderDate:    testDatetime,
					TrackingCode: testUUID,
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
			Name:             "[Create] Error FK ware house ID not valid",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
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
			MockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     testDatetime,
					TrackingCode:  testUUID,
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   0,
				},
			},
			MockResponse:   entities.PurchaseOrderEntity{},
			MockError:      repository.ErrPurchaseOrderFKWareHouseIdNotValid,
			ExpectedOutput: models.PurchaseOrder{},
			ExpectedErr:    service.ErrPurchaseOrderFKWareHouseIdNotValid,
		},
		{
			Name:             "[Create] Error FK buyer ID not valid",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
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
			MockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     testDatetime,
					TrackingCode:  testUUID,
					BuyerID:       0,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			MockResponse:   entities.PurchaseOrderEntity{},
			MockError:      repository.ErrPurchaseOrderFKBuyerIdNotValid,
			ExpectedOutput: models.PurchaseOrder{},
			ExpectedErr:    service.ErrPurchaseOrderFKBuyerIdNotValid,
		},
		{
			Name:             "[Create] Error FK order status ID not valid",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
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
			MockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     testDatetime,
					TrackingCode:  testUUID,
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 0,
					WarehouseID:   1,
				},
			},
			MockResponse:   entities.PurchaseOrderEntity{},
			MockError:      repository.ErrPurchaseOrderFKOrderStatusIdNotValid,
			ExpectedOutput: models.PurchaseOrder{},
			ExpectedErr:    service.ErrPurchaseOrderFKOrderStatusIdNotValid,
		},
		{
			Name:             "[Create] Error FK carrier ID not valid",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
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
			MockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     testDatetime,
					TrackingCode:  testUUID,
					BuyerID:       1,
					CarrierID:     0,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			MockResponse:   entities.PurchaseOrderEntity{},
			MockError:      repository.ErrPurchaseOrderFKCarrierIdNotValid,
			ExpectedOutput: models.PurchaseOrder{},
			ExpectedErr:    service.ErrPurchaseOrderFKCarrierIdNotValid,
		},
		{
			Name:             "[Create] Error Unexpected error",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
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
			MockParams: []interface{}{
				entities.PurchaseOrderEntity{
					OrderDate:     testDatetime,
					TrackingCode:  testUUID,
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			MockResponse:   entities.PurchaseOrderEntity{},
			MockError:      service.ErrPurchaseOrderUnexpectedError,
			ExpectedOutput: models.PurchaseOrder{},
			ExpectedErr:    service.ErrPurchaseOrderUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockPurchaseOrderRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertPurchaseOrderService(t, test, mockRepository, testDatetime, testUUID)

			mockRepository.AssertExpectations(t)
		})
	}
}
