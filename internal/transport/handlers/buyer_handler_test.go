package handlers_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/mocks/helpers"
	service_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
)

func switchBuyerTest(
	t *testing.T,
	test helpers.HandlerTestStruct,
	rt *chi.Mux,
	buyerHandler *handlers.BuyerHandler,
) {
	t.Helper()
	ctx := context.Background()

	switch test.ServiceMethod {
	case "GetAll":
		rt.Get("/api/v1/buyers", buyerHandler.GetAll(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, nil)
		helpers.CheckHandlerResponse[[]dto.BuyerDTO](t, test, result)
	case "GetById":
		rt.Get("/api/v1/buyers/{id}", buyerHandler.GetById(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, nil)
		helpers.CheckHandlerResponse[*dto.BuyerDTO](t, test, result)
	case "Create":
		rt.Post("/api/v1/buyers", buyerHandler.Create(&ctx))
		result := helpers.GetResult(
			t,
			rt,
			test.HttpPath,
			test.HttpMethod,
			test.HttpBody,
		)
		helpers.CheckHandlerResponse[*dto.BuyerDTO](t, test, result)
	case "Patch":
		rt.Patch("/api/v1/buyers/{id}", buyerHandler.Patch(&ctx))
		result := helpers.GetResult(
			t,
			rt,
			test.HttpPath,
			test.HttpMethod,
			test.HttpBody,
		)
		helpers.CheckHandlerResponse[*dto.BuyerDTO](t, test, result)
	case "Delete":
		rt.Delete("/api/v1/buyers/{id}", buyerHandler.Delete(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, nil)
		helpers.CheckHandlerResponse[*dto.BuyerDTO](t, test, result)
	case "GetReportPurchaseOrders":
		rt.Get("/api/v1/reportPurchaseOrders", buyerHandler.GetReportPurchaseOrders(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, nil)
		helpers.CheckHandlerResponse[[]dto.ReportPurchaseOrdersDTO](t, test, result)
	}
}

func assertBuyerHandler(
	t *testing.T,
	test helpers.HandlerTestStruct,
	mockService *service_mock.MockBuyerService,
) {
	t.Helper()

	rt := chi.NewRouter()
	buyerHandler := handlers.NewBuyerHandler(mockService)
	switchBuyerTest(t, test, rt, buyerHandler)
}

func TestBuyerGetAllHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[GetAll] OK Get all buyers",
			ServiceMethod: "GetAll",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "GET",
			MockResponse: []models.Buyer{
				{
					ID: 1,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: helpers.PtrStr("M1234567890"),
						FirstName:    helpers.PtrStr("John"),
						LastName:     helpers.PtrStr("Doe"),
					},
				},
				{
					ID: 2,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: helpers.PtrStr("F1098765432"),
						FirstName:    helpers.PtrStr("Jane"),
						LastName:     helpers.PtrStr("Doe"),
					},
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[[]dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "Get all buyers successful",
				Data: []dto.BuyerDTO{
					{
						ID:           1,
						CardNumberID: "M1234567890",
						FirstName:    "John",
						LastName:     "Doe",
					},
					{
						ID:           2,
						CardNumberID: "F1098765432",
						FirstName:    "Jane",
						LastName:     "Doe",
					},
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetAll] Error No buyers registered yet",
			ServiceMethod: "GetAll",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "GET",
			MockResponse:  make([]models.Buyer, 0),
			MockError:     service.ErrBuyerNoRegisteredBuyersYet,
			ExpectedOutput: helpers.TestGenericStruct[[]dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "ERR: No registered buyers yet",
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetAll] Error Unexpected error",
			ServiceMethod: "GetAll",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "GET",
			MockResponse:  make([]models.Buyer, 0),
			MockError:     service.ErrBuyerUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[[]dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)
			helpers.InitServiceMock(t, test, &mockService)
			assertBuyerHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}

func TestBuyerGetByIdHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[GetById] OK Get buyer by ID",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "GET",
			MockParams:    []interface{}{1},
			MockResponse: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "Get buyer by ID successful",
				Data: &dto.BuyerDTO{
					ID:           1,
					CardNumberID: "M1234567890",
					FirstName:    "John",
					LastName:     "Doe",
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetById] Error Invalid ID",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/buyers/InvalidID",
			HttpMethod:    "GET",
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer ID format",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[GetById] Error Buyer not found",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/buyers/99",
			HttpMethod:    "GET",
			MockParams:    []interface{}{99},
			MockResponse:  models.Buyer{},
			MockError:     service.ErrBuyerDoesNotExist,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested buyer does not exist in the database for the ID: 99",
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetById] Error Unexpected error",
			ServiceMethod: "GetById",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "GET",
			MockParams:    []interface{}{1},
			MockResponse:  models.Buyer{},
			MockError:     service.ErrBuyerUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)
			helpers.InitServiceMock(t, test, &mockService)
			assertBuyerHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}

func TestBuyerCreateHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[Create] OK Create new buyer",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"card_number_id": "M1234543210",
				"first_name":     "Baby",
				"last_name":      "Doe"
			}`),
			MockParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234543210"),
					FirstName:    helpers.PtrStr("Baby"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockResponse: models.Buyer{
				ID: 3,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234543210"),
					FirstName:    helpers.PtrStr("Baby"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusCreated,
				Msg:  "Create buyer successful",
				Data: &dto.BuyerDTO{
					ID:           3,
					CardNumberID: "M1234543210",
					FirstName:    "Baby",
					LastName:     "Doe",
				},
			},
			ExpectedStatusCode: http.StatusCreated,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error Buyer already exists",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"card_number_id": "M1234567890",
				"first_name":     "John",
				"last_name":      "Doe"
			}`),
			MockParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockResponse: models.Buyer{},
			MockError:    service.ErrBuyerAlreadyExists,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: A buyer already exists in the database with the card number ID: M1234567890",
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Create] Error Invalid JSON",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "POST",
			HttpBody:      []byte("invalid_json"),
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer JSON format",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[Create] Error Invalid buyer Create DTO",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"card_number_id": "invalid_card_number"
			}`),
			ExpectedOutput: helpers.TestGenericStruct[map[string]interface{}]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer Create DTO",
				Data: map[string]interface{}{
					"card_number_id": "Must start with a letter and be followed by 10 digits",
					"first_name":     "Field is required",
					"last_name":      "Field is required",
				},
			},
			IsExpectedOutputMapType: true,
			ExpectedStatusCode:      http.StatusBadRequest,
			ExpectedCalls:           0,
		},
		{
			Name:          "[Create] Error Unexpected error",
			ServiceMethod: "Create",
			HttpPath:      "/api/v1/buyers",
			HttpMethod:    "POST",
			HttpBody: []byte(`{
				"card_number_id": "M1234567890",
				"first_name":     "John",
				"last_name":      "Doe"
			}`),
			MockParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockResponse: models.Buyer{},
			MockError:    service.ErrBuyerUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)
			helpers.InitServiceMock(t, test, &mockService)
			assertBuyerHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}

func TestBuyerPatchHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[Patch] OK Update buyer",
			ServiceMethod: "Patch",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "PATCH",
			HttpBody: []byte(`{
				"card_number_id": "M1234567876"
			}`),
			MockParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockResponse: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "Update buyer successful",
				Data: &dto.BuyerDTO{
					ID:           1,
					CardNumberID: "M1234567876",
					FirstName:    "John",
					LastName:     "Doe",
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Patch] Error Invalid ID",
			ServiceMethod: "Patch",
			HttpPath:      "/api/v1/buyers/InvalidID",
			HttpMethod:    "PATCH",
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer ID format",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[Patch] Error Invalid JSON",
			ServiceMethod: "Patch",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "PATCH",
			HttpBody:      []byte("invalid_json"),
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer JSON format",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[Patch] Error Invalid buyer Patch DTO",
			ServiceMethod: "Patch",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "PATCH",
			HttpBody: []byte(`{
				"card_number_id": "invalid_card_number"
			}`),
			ExpectedOutput: helpers.TestGenericStruct[map[string]interface{}]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer Patch DTO",
				Data: map[string]interface{}{
					"card_number_id": "Must start with a letter and be followed by 10 digits",
				},
			},
			IsExpectedOutputMapType: true,
			ExpectedStatusCode:      http.StatusBadRequest,
			ExpectedCalls:           0,
		},
		{
			Name:          "[Patch] Error Buyer not found",
			ServiceMethod: "Patch",
			HttpPath:      "/api/v1/buyers/99",
			HttpMethod:    "PATCH",
			HttpBody: []byte(`{
				"card_number_id": "M1234567876"
			}`),
			MockParams: []interface{}{
				99,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockResponse: models.Buyer{},
			MockError:    service.ErrBuyerDoesNotExist,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested buyer does not exist in the database for the ID: 99",
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Patch] Error Buyer already exists",
			ServiceMethod: "Patch",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "PATCH",
			HttpBody: []byte(`{
				"card_number_id": "F1098765432"
			}`),
			MockParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("F1098765432"),
				},
			},
			MockResponse: models.Buyer{},
			MockError:    service.ErrBuyerAlreadyExists,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: A buyer already exists in the database with the card number ID: F1098765432",
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)
			helpers.InitServiceMock(t, test, &mockService)
			assertBuyerHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}

func TestBuyerDeleteHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:               "[Delete] OK Delete buyer",
			ServiceMethod:      "Delete",
			HttpPath:           "/api/v1/buyers/1",
			HttpMethod:         "DELETE",
			MockParams:         []interface{}{1},
			ExpectedStatusCode: http.StatusNoContent,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Delete] Error Buyer not found",
			ServiceMethod: "Delete",
			HttpPath:      "/api/v1/buyers/99",
			HttpMethod:    "DELETE",
			MockParams:    []interface{}{99},
			MockError:     service.ErrBuyerDoesNotExist,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested buyer does not exist in the database for the ID: 99",
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Delete] Error Invalid ID",
			ServiceMethod: "Delete",
			HttpPath:      "/api/v1/buyers/InvalidID",
			HttpMethod:    "DELETE",
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer ID format",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[Delete] Error Cannot delete buyer with orders",
			ServiceMethod: "Delete",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "DELETE",
			MockParams:    []interface{}{1},
			MockError:     service.ErrBuyerCannotDeleteBuyerWithOrders,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: Cannot delete buyer with orders. Please delete the related purchase orders first",
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedCalls:      1,
		},
		{
			Name:          "[Delete] Error Unexpected error",
			ServiceMethod: "Delete",
			HttpPath:      "/api/v1/buyers/1",
			HttpMethod:    "DELETE",
			MockParams:    []interface{}{1},
			MockError:     service.ErrBuyerUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)
			helpers.InitServiceMock(t, test, &mockService)
			assertBuyerHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}

func TestBuyerGetReportPurchaseOrdersHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[GetReportPurchaseOrders] Ok Get purchase order report",
			ServiceMethod: "GetReportPurchaseOrders",
			HttpPath:      "/api/v1/reportPurchaseOrders",
			HttpMethod:    "GET",
			MockParams:    []interface{}{helpers.PtrIntNil()},
			MockResponse: []models.ReportPurchaseOrders{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
				{
					ID:                  2,
					CardNumberID:        "F1098765432",
					FirstName:           "Jane",
					LastName:            "Doe",
					PurchaseOrdersCount: 27,
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[[]dto.ReportPurchaseOrdersDTO]{
				Code: http.StatusOK,
				Msg:  "Get all orders",
				Data: []dto.ReportPurchaseOrdersDTO{
					{
						ID:                  1,
						CardNumberID:        "M1234567890",
						FirstName:           "John",
						LastName:            "Doe",
						PurchaseOrdersCount: 4,
					},
					{
						ID:                  2,
						CardNumberID:        "F1098765432",
						FirstName:           "Jane",
						LastName:            "Doe",
						PurchaseOrdersCount: 27,
					},
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetReportPurchaseOrders] Error Invalid ID",
			ServiceMethod: "GetReportPurchaseOrders",
			HttpPath:      "/api/v1/reportPurchaseOrders?id=InvalidID",
			HttpMethod:    "GET",
			ExpectedOutput: helpers.TestGenericStruct[[]dto.ReportPurchaseOrdersDTO]{
				Code: http.StatusBadRequest,
				Msg:  "ERR: Invalid buyer ID format",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedCalls:      0,
		},
		{
			Name:          "[GetReportPurchaseOrders] Error Unexpected error",
			ServiceMethod: "GetReportPurchaseOrders",
			HttpPath:      "/api/v1/reportPurchaseOrders?id=1",
			HttpMethod:    "GET",
			MockParams:    []interface{}{helpers.PtrInt(1)},
			MockResponse:  make([]models.ReportPurchaseOrders, 0),
			MockError:     service.ErrBuyerUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[[]dto.ReportPurchaseOrdersDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetReportPurchaseOrders] Ok Get purchase order report by ID",
			ServiceMethod: "GetReportPurchaseOrders",
			HttpPath:      "/api/v1/reportPurchaseOrders?id=1",
			HttpMethod:    "GET",
			MockParams:    []interface{}{helpers.PtrInt(1)},
			MockResponse: []models.ReportPurchaseOrders{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[[]dto.ReportPurchaseOrdersDTO]{
				Code: http.StatusOK,
				Msg:  "Get all orders",
				Data: []dto.ReportPurchaseOrdersDTO{
					{
						ID:                  1,
						CardNumberID:        "M1234567890",
						FirstName:           "John",
						LastName:            "Doe",
						PurchaseOrdersCount: 4,
					},
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)
			helpers.InitServiceMock(t, test, &mockService)
			assertBuyerHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}
