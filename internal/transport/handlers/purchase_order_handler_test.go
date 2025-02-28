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

type PurchaseOrderGenericStruct[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data,omitempty"`
}

func getPurchaseOrderResult(
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

func getPurchaseOrderGenericStruct[T interface{}](
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

func assertPurchaseOrderResponse[T interface{}](
	t *testing.T,
	responseDTO PurchaseOrderGenericStruct[T],
	responseStatusCode int,
	expectedStatusCode int,
	expectedOutput interface{},
) {
	t.Helper()

	assert.Equal(t, expectedStatusCode, responseStatusCode, "HTTP Status Code mismatch")

	if expectedOutput != nil {
		expectedDTO, ok := expectedOutput.(PurchaseOrderGenericStruct[T])
		if !ok {
			t.Fatal("Failed to convert expectedOutput to PurchaseOrderGenericStruct")
		}

		assert.Equal(t, expectedDTO.Code, responseDTO.Code, "ResponseDTO Code mismatch")
		assert.Equal(t, expectedDTO.Msg, responseDTO.Msg, "ResponseDTO Message mismatch")
		assert.Equal(t, expectedDTO.Data, responseDTO.Data, "ResponseDTO Data mismatch")
	}
}

func TestPurchaseOrderHandler(t *testing.T) {
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
			name:          "[GetById] OK Get purchase order by ID",
			serviceMethod: "GetById",
			httpPath:      "/api/v1/purchaseOrders/1",
			httpMethod:    "GET",
			mockParams:    []interface{}{1},
			mockResponse: models.PurchaseOrder{
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
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
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
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetById] Error Purchase order not found",
			serviceMethod: "GetById",
			httpPath:      "/api/v1/purchaseOrders/99",
			httpMethod:    "GET",
			mockParams:    []interface{}{99},
			mockResponse:  models.PurchaseOrder{},
			mockError:     service.ErrPurchaseOrderDoesNotExist,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusNotFound,
				Msg:  "ERR: The requested purchase order does not exist in the database for the ID: 99",
				Data: nil,
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "[GetById] Error Unexpected error",
			serviceMethod: "GetById",
			httpPath:      "/api/v1/purchaseOrders/1",
			httpMethod:    "GET",
			mockParams:    []interface{}{1},
			mockResponse:  models.PurchaseOrder{},
			mockError:     service.ErrPurchaseOrderUnexpectedError,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested purchase order, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:          "[Create] OK Create new purchase order",
			serviceMethod: "Create",
			httpPath:      "/api/v1/purchaseOrders",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    1,
			},
			mockParams: []interface{}{
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
			mockResponse: models.PurchaseOrder{
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
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
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
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:          "[Create] Error FK ware house ID not valid",
			serviceMethod: "Create",
			httpPath:      "/api/v1/purchaseOrders",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    0,
			},
			mockParams: []interface{}{
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
			mockResponse: models.PurchaseOrder{},
			mockError:    service.ErrPurchaseOrderFKWareHouseIdNotValid,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the warehouse ID of the requested purchase order is not valid",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Create] Error FK buyer ID not valid",
			serviceMethod: "Create",
			httpPath:      "/api/v1/purchaseOrders",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"buyer_id":        0,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    1,
			},
			mockParams: []interface{}{
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
			mockResponse: models.PurchaseOrder{},
			mockError:    service.ErrPurchaseOrderFKBuyerIdNotValid,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the buyer ID of the requested purchase order is not valid",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Create] Error FK order status ID not valid",
			serviceMethod: "Create",
			httpPath:      "/api/v1/purchaseOrders",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 0,
				"warehouse_id":    1,
			},
			mockParams: []interface{}{
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
			mockResponse: models.PurchaseOrder{},
			mockError:    service.ErrPurchaseOrderFKOrderStatusIdNotValid,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the order status ID of the requested purchase order is not valid",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Create] Error FK carrier ID not valid",
			serviceMethod: "Create",
			httpPath:      "/api/v1/purchaseOrders",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"buyer_id":        1,
				"carrier_id":      0,
				"order_status_id": 1,
				"warehouse_id":    1,
			},
			mockParams: []interface{}{
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
			mockResponse: models.PurchaseOrder{},
			mockError:    service.ErrPurchaseOrderFKCarrierIdNotValid,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusConflict,
				Msg:  "ERR: The foreign key for the carrier ID of the requested purchase order is not valid",
				Data: nil,
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:          "[Create] Error Unexpected error",
			serviceMethod: "Create",
			httpPath:      "/api/v1/purchaseOrders",
			httpMethod:    "POST",
			httpBody: map[string]interface{}{
				"buyer_id":        1,
				"carrier_id":      1,
				"order_status_id": 1,
				"warehouse_id":    1,
			},
			mockParams: []interface{}{
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
			mockResponse: models.PurchaseOrder{},
			mockError:    service.ErrPurchaseOrderUnexpectedError,
			expectedOutput: PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested purchase order, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service_mock.MockPurchaseOrderService)

			if tt.mockResponse != nil {
				mockService.On(tt.serviceMethod, tt.mockParams...).
					Return(tt.mockResponse, tt.mockError)
			} else {
				mockService.On(tt.serviceMethod, tt.mockParams...).Return(tt.mockError)
			}

			rt := chi.NewRouter()
			purchaseOrderHandler := handlers.NewPurchaseOrderHandler(mockService)

			switch tt.serviceMethod {
			case "GetById":
				rt.Get("/api/v1/purchaseOrders/{id}", purchaseOrderHandler.GetById())

				result := getPurchaseOrderResult(t, rt, tt.httpPath, tt.httpMethod, nil)
				responseDTO := getPurchaseOrderGenericStruct[PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]](
					t,
					result.Body,
				)
				assertPurchaseOrderResponse(
					t,
					responseDTO,
					result.StatusCode,
					tt.expectedStatusCode,
					tt.expectedOutput,
				)
			case "Create":
				rt.Post("/api/v1/purchaseOrders", purchaseOrderHandler.Create())

				result := getPurchaseOrderResult(t, rt, tt.httpPath, tt.httpMethod, tt.httpBody)
				responseDTO := getPurchaseOrderGenericStruct[PurchaseOrderGenericStruct[*dto.PurchaseOrderDTO]](
					t,
					result.Body,
				)
				assertPurchaseOrderResponse(
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
