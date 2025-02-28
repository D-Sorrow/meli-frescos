package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	service_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
)

type BuyerGenericStruct[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data,omitempty"`
}

func ptrStr(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}

func ptrIntNil() *int {
	return nil
}

func getBuyerResult(
	t *testing.T,
	rt *chi.Mux,
	path string,
	httpMethod string,
	httpBody map[string]interface{},
) (result *http.Response) {
	t.Helper()

	var req *http.Request

	if httpBody != nil {
		jsonBody, err := json.Marshal(httpBody)
		if err != nil {
			t.Fatalf("JSON serialization error: %v", err)
		}

		req = httptest.NewRequest(httpMethod, path, bytes.NewBuffer(jsonBody))
	} else {
		req = httptest.NewRequest(httpMethod, path, nil)
	}

	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, req)
	result = rec.Result()
	return
}

func getBuyerGenericStruct[T interface{}](
	t *testing.T,
	body io.ReadCloser,
) (responseDTOGenericStruct T) {
	t.Helper()

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		t.Fatal("Body reading error:", err)
	}
	defer body.Close()

	if err := json.Unmarshal(bodyBytes, &responseDTOGenericStruct); err != nil {
		t.Fatal("JSON parsing error:", err)
	}

	return
}

func assertBuyerResponse[T interface{}](
	t *testing.T,
	responseDTO BuyerGenericStruct[T],
	responseStatusCode int,
	expectedStatusCode int,
	expectedOutput interface{},
) {
	t.Helper()

	assert.Equal(t, expectedStatusCode, responseStatusCode, "HTTP Status Code mismatch")

	if expectedOutput != nil {
		expectedDTO, ok := expectedOutput.(BuyerGenericStruct[T])
		if !ok {
			t.Fatal("Failed to convert expectedOutput to BuyerGenericStruct")
		}

		assert.Equal(t, expectedDTO.Code, responseDTO.Code, "ResponseDTO Code mismatch")
		assert.Equal(t, expectedDTO.Msg, responseDTO.Msg, "ResponseDTO Message mismatch")
		assert.Equal(t, expectedDTO.Data, responseDTO.Data, "ResponseDTO Data mismatch")
	}
}

func TestBuyerHandler(t *testing.T) {
	tests := []struct {
		name               string
		serviceMethod      string
		httpPath           string
		httpMethod         string
		httpBody           map[string]interface{}
		mockParams         []interface{}
		mockResponse       interface{}
		mockError          error
		expectedOutput     interface{}
		expectedStatusCode int
	}{
		{
			name:          "[GetAll] OK Get all buyers",
			serviceMethod: "GetAll",
			httpPath:      "/api/v1/buyers",
			httpMethod:    "GET",
			mockResponse: []models.Buyer{
				{
					ID: 1,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: ptrStr("M1234567890"),
						FirstName:    ptrStr("John"),
						LastName:     ptrStr("Doe"),
					},
				},
				{
					ID: 2,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: ptrStr("F1098765432"),
						FirstName:    ptrStr("Jane"),
						LastName:     ptrStr("Doe"),
					},
				},
			},
			expectedOutput: BuyerGenericStruct[[]dto.BuyerDTO]{
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
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetAll] Error No buyers registered yet",
			serviceMethod: "GetAll",
			httpPath:      "/api/v1/buyers",
			httpMethod:    "GET",
			mockResponse:  make([]models.Buyer, 0),
			mockError:     service.ErrBuyerNoRegisteredBuyersYet,
			expectedOutput: BuyerGenericStruct[[]dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "ERR: No registered buyers yet",
				Data: nil,
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetAll] Error Unexpected error",
			serviceMethod: "GetAll",
			httpPath:      "/api/v1/buyers",
			httpMethod:    "GET",
			mockResponse:  make([]models.Buyer, 0),
			mockError:     service.ErrBuyerUnexpectedError,
			expectedOutput: BuyerGenericStruct[[]dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:          "[GetById] OK Get buyer by ID",
			serviceMethod: "GetById",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "GET",
			mockParams:    []interface{}{1},
			mockResponse: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "Get buyer by ID successful",
				Data: &dto.BuyerDTO{
					ID:           1,
					CardNumberID: "M1234567890",
					FirstName:    "John",
					LastName:     "Doe",
				},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetById] Error Buyer not found",
			serviceMethod: "GetById",
			httpPath:      "/api/v1/buyers/99",
			httpMethod:    "GET",
			mockParams:    []interface{}{99},
			mockResponse:  models.Buyer{},
			mockError:     service.ErrBuyerDoesNotExist,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested buyer does not exist in the database for the ID: 99",
				Data: nil,
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "[GetById] Error Unexpected error",
			serviceMethod: "GetById",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "GET",
			mockParams:    []interface{}{1},
			mockResponse:  models.Buyer{},
			mockError:     service.ErrBuyerUnexpectedError,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:          "[Create] OK Create new buyer",
			serviceMethod: "Create",
			httpPath:      "/api/v1/buyers",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"card_number_id": "M1234543210",
				"first_name":     "Baby",
				"last_name":      "Doe",
			},
			mockParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234543210"),
					FirstName:    ptrStr("Baby"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockResponse: models.Buyer{
				ID: 3,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: ptrStr("M1234543210"),
					FirstName:    ptrStr("Baby"),
					LastName:     ptrStr("Doe"),
				},
			},
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusCreated,
				Msg:  "Create buyer successful",
				Data: &dto.BuyerDTO{
					ID:           3,
					CardNumberID: "M1234543210",
					FirstName:    "Baby",
					LastName:     "Doe",
				},
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:          "[Create] Error Buyer already exists",
			serviceMethod: "Create",
			httpPath:      "/api/v1/buyers",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"card_number_id": "M1234567890",
				"first_name":     "John",
				"last_name":      "Doe",
			},
			mockParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockResponse: models.Buyer{},
			mockError:    service.ErrBuyerAlreadyExists,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: A buyer already exists in the database with the card number ID: M1234567890",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Create] Error Unexpected error",
			serviceMethod: "Create",
			httpPath:      "/api/v1/buyers",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"card_number_id": "M1234567890",
				"first_name":     "John",
				"last_name":      "Doe",
			},
			mockParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockResponse: models.Buyer{},
			mockError:    service.ErrBuyerUnexpectedError,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:          "[Patch] OK Update buyer",
			serviceMethod: "Patch",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "PATCH",
			httpBody: map[string]interface{}{
				"card_number_id": "M1234567876",
			},
			mockParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockResponse: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusOK,
				Msg:  "Update buyer successful",
				Data: &dto.BuyerDTO{
					ID:           1,
					CardNumberID: "M1234567876",
					FirstName:    "John",
					LastName:     "Doe",
				},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[Patch] Error Buyer not found",
			serviceMethod: "Patch",
			httpPath:      "/api/v1/buyers/99",
			httpMethod:    "PATCH",
			httpBody: map[string]interface{}{
				"card_number_id": "M1234567876",
			},
			mockParams: []interface{}{
				99,
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockResponse: models.Buyer{},
			mockError:    service.ErrBuyerDoesNotExist,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested buyer does not exist in the database for the ID: 99",
				Data: nil,
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "[Patch] Error Buyer already exists",
			serviceMethod: "Patch",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "PATCH",
			httpBody: map[string]interface{}{
				"card_number_id": "F1098765432",
			},
			mockParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: ptrStr("F1098765432"),
				},
			},
			mockResponse: models.Buyer{},
			mockError:    service.ErrBuyerAlreadyExists,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: A buyer already exists in the database with the card number ID: F1098765432",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Patch] Error Unexpected error",
			serviceMethod: "Patch",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "PATCH",
			httpBody: map[string]interface{}{
				"card_number_id": "M1234567876",
			},
			mockParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockResponse: models.Buyer{},
			mockError:    service.ErrBuyerUnexpectedError,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "[Delete] OK Delete buyer",
			serviceMethod:      "Delete",
			httpPath:           "/api/v1/buyers/1",
			httpMethod:         "DELETE",
			mockParams:         []interface{}{1},
			mockResponse:       nil,
			mockError:          nil,
			expectedOutput:     nil,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:          "[Delete] Error Buyer not found",
			serviceMethod: "Delete",
			httpPath:      "/api/v1/buyers/99",
			httpMethod:    "DELETE",
			mockParams:    []interface{}{99},
			mockResponse:  nil,
			mockError:     service.ErrBuyerDoesNotExist,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested buyer does not exist in the database for the ID: 99",
				Data: nil,
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "[Delete] Error Cannot delete buyer with orders",
			serviceMethod: "Delete",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "DELETE",
			mockParams:    []interface{}{1},
			mockResponse:  nil,
			mockError:     service.ErrBuyerCannotDeleteBuyerWithOrders,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: Cannot delete buyer with orders. Please delete the related purchase orders first",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Delete] Error Unexpected error",
			serviceMethod: "Delete",
			httpPath:      "/api/v1/buyers/1",
			httpMethod:    "DELETE",
			mockParams:    []interface{}{1},
			mockResponse:  nil,
			mockError:     service.ErrBuyerUnexpectedError,
			expectedOutput: BuyerGenericStruct[*dto.BuyerDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested buyer, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:          "[GetReportPurchaseOrders] Ok Get purchase order report",
			serviceMethod: "GetReportPurchaseOrders",
			httpPath:      "/api/v1/reportPurchaseOrders",
			httpMethod:    "GET",
			mockParams:    []interface{}{ptrIntNil()},
			mockResponse: []models.ReportPurchaseOrders{
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
			expectedOutput: BuyerGenericStruct[[]dto.ReportPurchaseOrdersDTO]{
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
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetReportPurchaseOrders] Ok Get purchase order report by ID",
			serviceMethod: "GetReportPurchaseOrders",
			httpPath:      "/api/v1/reportPurchaseOrders?id=1",
			httpMethod:    "GET",
			mockParams:    []interface{}{ptrInt(1)},
			mockResponse: []models.ReportPurchaseOrders{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
			},
			expectedOutput: BuyerGenericStruct[[]dto.ReportPurchaseOrdersDTO]{
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
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service_mock.MockBuyerService)

			if tt.mockResponse != nil {
				mockService.On(tt.serviceMethod, tt.mockParams...).
					Return(tt.mockResponse, tt.mockError)
			} else {
				mockService.On(tt.serviceMethod, tt.mockParams...).Return(tt.mockError)
			}

			rt := chi.NewRouter()
			buyerHandler := handlers.NewBuyerHandler(mockService)

			switch tt.serviceMethod {
			case "GetAll":
				rt.Get("/api/v1/buyers", buyerHandler.GetAll())

				result := getBuyerResult(t, rt, tt.httpPath, tt.httpMethod, nil)
				responseDTO := getBuyerGenericStruct[BuyerGenericStruct[[]dto.BuyerDTO]](
					t,
					result.Body,
				)
				assertBuyerResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)

			case "GetById":
				rt.Get("/api/v1/buyers/{id}", buyerHandler.GetById())

				result := getBuyerResult(t, rt, tt.httpPath, tt.httpMethod, nil)
				responseDTO := getBuyerGenericStruct[BuyerGenericStruct[*dto.BuyerDTO]](
					t,
					result.Body,
				)
				assertBuyerResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)
			case "Create":
				rt.Post("/api/v1/buyers", buyerHandler.Create())

				result := getBuyerResult(t, rt, tt.httpPath, tt.httpMethod, tt.httpBody)
				responseDTO := getBuyerGenericStruct[BuyerGenericStruct[*dto.BuyerDTO]](
					t,
					result.Body,
				)
				assertBuyerResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)
			case "Patch":
				rt.Patch("/api/v1/buyers/{id}", buyerHandler.Patch())

				result := getBuyerResult(t, rt, tt.httpPath, tt.httpMethod, tt.httpBody)
				responseDTO := getBuyerGenericStruct[BuyerGenericStruct[*dto.BuyerDTO]](
					t,
					result.Body,
				)
				assertBuyerResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)
			case "Delete":
				rt.Delete("/api/v1/buyers/{id}", buyerHandler.Delete())

				result := getBuyerResult(t, rt, tt.httpPath, tt.httpMethod, nil)
				responseDTO := getBuyerGenericStruct[BuyerGenericStruct[*dto.BuyerDTO]](
					t,
					result.Body,
				)
				assertBuyerResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)
			case "GetReportPurchaseOrders":
				rt.Get("/api/v1/reportPurchaseOrders", buyerHandler.GetReportPurchaseOrders())

				result := getBuyerResult(t, rt, tt.httpPath, tt.httpMethod, nil)
				responseDTO := getBuyerGenericStruct[BuyerGenericStruct[[]dto.ReportPurchaseOrdersDTO]](
					t,
					result.Body,
				)
				assertBuyerResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)
			}

			mockService.AssertExpectations(t)
		})
	}
}
