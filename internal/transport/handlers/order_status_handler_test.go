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

type OrderStatusGenericStruct[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data,omitempty"`
}

func getOrderStatusResult(
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

func getOrderStatusGenericStruct[T interface{}](
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

func assertOrderStatusResponse[T interface{}](
	t *testing.T,
	responseDTO OrderStatusGenericStruct[T],
	responseStatusCode int,
	expectedStatusCode int,
	expectedOutput interface{},
) {
	t.Helper()

	assert.Equal(t, expectedStatusCode, responseStatusCode, "HTTP Status Code mismatch")

	if expectedOutput != nil {
		expectedDTO, ok := expectedOutput.(OrderStatusGenericStruct[T])
		if !ok {
			t.Fatal("Failed to convert expectedOutput to OrderStatusGenericStruct")
		}

		assert.Equal(t, expectedDTO.Code, responseDTO.Code, "ResponseDTO Code mismatch")
		assert.Equal(t, expectedDTO.Msg, responseDTO.Msg, "ResponseDTO Message mismatch")
		assert.Equal(t, expectedDTO.Data, responseDTO.Data, "ResponseDTO Data mismatch")
	}
}

func TestOrderStatusHandler(t *testing.T) {
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
			name:          "[GetAll] OK Get all order statuses",
			serviceMethod: "GetAll",
			httpPath:      "/api/v1/orderStatus",
			httpMethod:    "GET",
			mockResponse: []models.OrderStatus{
				{
					ID: 1,
					OrderStatusAttributes: models.OrderStatusAttributes{
						Description: "Pendiente por pago",
					},
				},
				{
					ID: 2,
					OrderStatusAttributes: models.OrderStatusAttributes{
						Description: "Pagado",
					},
				},
			},
			expectedOutput: OrderStatusGenericStruct[[]dto.OrderStatusDTO]{
				Code: http.StatusOK,
				Msg:  "Get all order statuses successful",
				Data: []dto.OrderStatusDTO{
					{
						ID:          1,
						Description: "Pendiente por pago",
					},
					{
						ID:          2,
						Description: "Pagado",
					},
				},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetAll] Error No order statuses registered yet",
			serviceMethod: "GetAll",
			httpPath:      "/api/v1/orderStatus",
			httpMethod:    "GET",
			mockResponse:  make([]models.OrderStatus, 0),
			mockError:     service.ErrOrderStatusNoRegisteredOrderStatusesYet,
			expectedOutput: OrderStatusGenericStruct[[]dto.OrderStatusDTO]{
				Code: http.StatusOK,
				Msg:  "ERR: No registered order statuses yet",
				Data: nil,
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:          "[GetAll] Error Unexpected error",
			serviceMethod: "GetAll",
			httpPath:      "/api/v1/orderStatus",
			httpMethod:    "GET",
			mockResponse:  make([]models.OrderStatus, 0),
			mockError:     service.ErrOrderStatusUnexpectedError,
			expectedOutput: OrderStatusGenericStruct[[]dto.OrderStatusDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested order status, please try again later",
				Data: nil,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(service_mock.MockOrderStatusService)

			if tt.mockResponse != nil {
				mockService.On(tt.serviceMethod, tt.mockParams...).
					Return(tt.mockResponse, tt.mockError)
			} else {
				mockService.On(tt.serviceMethod, tt.mockParams...).Return(tt.mockError)
			}

			rt := chi.NewRouter()
			orderStatusHandler := handlers.NewOrderStatusHandler(mockService)

			switch tt.serviceMethod {
			case "GetAll":
				rt.Get("/api/v1/orderStatus", orderStatusHandler.GetAll())

				result := getOrderStatusResult(t, rt, tt.httpPath, tt.httpMethod, nil)
				responseDTO := getOrderStatusGenericStruct[OrderStatusGenericStruct[[]dto.OrderStatusDTO]](
					t,
					result.Body,
				)
				assertOrderStatusResponse(
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
