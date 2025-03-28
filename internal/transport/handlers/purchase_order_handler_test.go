package handlers_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/melisource/fury_go-core/pkg/web"
	"github.com/melisource/fury_go-platform/pkg/fury"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/helpers"
	service_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/service"
)

func switchPurchaseOrderTest(
	t *testing.T,
	test helpers.HandlerTestStruct,
	rt *web.Router,
	purchaseOrderHandler *handlers.PurchaseOrderHandler,
) {
	t.Helper()
	ctx := context.Background()

	switch test.ServiceMethod {
	case "GetById":
		rt.Get("/api/v1/purchaseOrders/{id}", purchaseOrderHandler.GetById(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, nil)
		helpers.CheckHandlerResponse[*dto.PurchaseOrderDTO](t, test, result)
	case "Create":
		rt.Post("/api/v1/purchaseOrders", purchaseOrderHandler.Create(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, test.HttpBody)
		helpers.CheckHandlerResponse[*dto.PurchaseOrderDTO](t, test, result)
	}
}

func assertPurchaseOrderHandler(
	t *testing.T,
	test helpers.HandlerTestStruct,
	mockService *service_mock.MockPurchaseOrderService,
) {
	t.Helper()

	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	rt := app.Router
	purchaseOrderHandler := handlers.NewPurchaseOrderHandler(mockService)
	switchPurchaseOrderTest(t, test, rt, purchaseOrderHandler)
}

func TestPurchaseOrderGetByIdHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[GetById] OK Get purchase order by ID",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/purchaseOrders/1",
			HttpMethod:    "GET",
			MockParams:    []interface{}{1},
			MockResponse: models.PurchaseOrder{
				ID: 1,
				PurchaseOrderAttributes: models.PurchaseOrderAttributes{
					OrderNumber:  "OR0001",
					OrderDate:    "2025-02-15T04:08:52Z",
					TrackingCode: "d2af4044-3c92-4d0b-9de6-9678c755f6c1",
				},
				PurchaseOrderFKs: models.PurchaseOrderFKs{
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusOK,
				Msg:  "Get purchase order by ID successful",
				Data: &dto.PurchaseOrderDTO{
					ID:            1,
					OrderNumber:   "OR0001",
					OrderDate:     "2025-02-15T04:08:52Z",
					TrackingCode:  "d2af4044-3c92-4d0b-9de6-9678c755f6c1",
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetById] Error Purchase order not found",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/purchaseOrders/99",
			HttpMethod:    "GET",
			MockParams:    []interface{}{99},
			MockResponse:  models.PurchaseOrder{},
			MockError:     service.ErrPurchaseOrderDoesNotExist,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested purchase order does not exist in the database for the ID: 99",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetById] Error Unexpected error",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/purchaseOrders/1",
			HttpMethod:    "GET",
			MockParams:    []interface{}{1},
			MockResponse:  models.PurchaseOrder{},
			MockError:     service.ErrPurchaseOrderUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested purchase order, please try again later",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetById] Error Invalid ID",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/purchaseOrders/InvalidID",
			HttpMethod:    "GET",
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid purchase order ID format",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockPurchaseOrderService)
			helpers.InitServiceMock(t, test, &mockService)
			assertPurchaseOrderHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}

func TestPurchaseOrderCreateHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[Create] Error Invalid JSON",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody:      []byte("invalid_json"),
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid purchase order JSON format",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[Create] OK Create new purchase order",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    1
			}`),
			MockParams: []interface{}{
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
			MockResponse: models.PurchaseOrder{
				ID: 2,
				PurchaseOrderAttributes: models.PurchaseOrderAttributes{
					OrderNumber:  "OR0002",
					OrderDate:    "2025-02-27T13:08:52Z",
					TrackingCode: "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
				},
				PurchaseOrderFKs: models.PurchaseOrderFKs{
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusCreated,
				Msg:  "Create purchase order successful",
				Data: &dto.PurchaseOrderDTO{
					ID:            2,
					OrderNumber:   "OR0002",
					OrderDate:     "2025-02-27T13:08:52Z",
					TrackingCode:  "d07a0493-0882-4702-b3aa-2dfe9ae108fe",
					BuyerID:       1,
					CarrierID:     1,
					OrderStatusID: 1,
					WarehouseID:   1,
				},
			},
			ExpectedStatusCode: http.StatusCreated,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error FK ware house ID not valid",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    0
			}`),
			MockParams: []interface{}{
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
			MockResponse: models.PurchaseOrder{},
			MockError:    service.ErrPurchaseOrderFKWareHouseIdNotValid,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the warehouse ID of the requested purchase order is not valid",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error FK buyer ID not valid",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"buyer_id":        0,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    1
			}`),
			MockParams: []interface{}{
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
			MockResponse: models.PurchaseOrder{},
			MockError:    service.ErrPurchaseOrderFKBuyerIdNotValid,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the buyer ID of the requested purchase order is not valid",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error FK order status ID not valid",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 0,
				"warehouse_id":    1
			}`),
			MockParams: []interface{}{
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
			MockResponse: models.PurchaseOrder{},
			MockError:    service.ErrPurchaseOrderFKOrderStatusIdNotValid,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the order status ID of the requested purchase order is not valid",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error FK carrier ID not valid",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"buyer_id":        1,
				"carrier_id":      0,
				"order_status_id": 1,
				"warehouse_id":    1
			}`),
			MockParams: []interface{}{
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
			MockResponse: models.PurchaseOrder{},
			MockError:    service.ErrPurchaseOrderFKCarrierIdNotValid,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the carrier ID of the requested purchase order is not valid",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error Unexpected error",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/purchaseOrders",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    1
			}`),
			MockParams: []interface{}{
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
			MockResponse: models.PurchaseOrder{},
			MockError:    service.ErrPurchaseOrderUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested purchase order, please try again later",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockPurchaseOrderService)
			helpers.InitServiceMock(t, test, &mockService)
			assertPurchaseOrderHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}
