package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	serviceMock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/service"
	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/require"
)

type inboundResponseDto[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data"`
}

func TestCreateInboundOrder(t *testing.T) {
	ctx := context.Background()

	t.Run("CreateInboundOrder success with valid request", func(t *testing.T) {
		mockService := new(serviceMock.MockInboundOrderService)
		handler := NewInboundOrderHandler(mockService)

		inboundOrderRequest := dto.InboundOrderRequestDTO{
			OrderDate:      "2024-10-01",
			OrderNumber:    "12345",
			EmployeeId:     1,
			ProductBatchId: 1,
			WarehouseId:    1,
		}

		inboundOrderModel := mappers.InboundOrderRequestDTOToModel(inboundOrderRequest)
		inboundOrderResponse := mappers.InboundOrderModelToResponseDTO(*inboundOrderModel)

		mockService.On("CreateInboundOrder", inboundOrderModel).Return(nil)

		reqBody, _ := json.Marshal(inboundOrderRequest)
		req := httptest.NewRequest("POST", "/api/v1/inboundOrders", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/inboundOrders", handler.CreateInboundOrder(&ctx))
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusCreated, rr.Code)
		var response inboundResponseDto[*dto.InboundOrderResponseDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, "Success", response.Msg)
		require.Equal(t, inboundOrderResponse, response.Data)
		mockService.AssertExpectations(t)
	})

	t.Run("CreateInboundOrder fails with invalid request body", func(t *testing.T) {
		mockService := new(serviceMock.MockInboundOrderService)
		handler := NewInboundOrderHandler(mockService)

		expectedResponse := dto.ResponseDTO{
			Code: http.StatusBadRequest,
			Msg:  "El cuerpo de la petición está mal formado",
			Data: nil,
		}

		req := httptest.NewRequest(
			"POST",
			"/api/v1/inboundOrders",
			bytes.NewBuffer([]byte("invalid body")),
		)
		rr := httptest.NewRecorder()
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/inboundOrders", handler.CreateInboundOrder(&ctx))
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.ResponseDTO
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("CreateInboundOrder fails with validation error", func(t *testing.T) {
		mockService := new(serviceMock.MockInboundOrderService)
		handler := NewInboundOrderHandler(mockService)

		inboundOrderRequest := dto.InboundOrderRequestDTO{
			OrderDate:      "",
			OrderNumber:    "12345",
			EmployeeId:     1,
			ProductBatchId: 1,
			WarehouseId:    1,
		}

		expectedResponse := dto.ResponseDTO{
			Code: http.StatusBadRequest,
			Msg:  "Validación fallida:  required OrderDate, ",
			Data: nil,
		}

		reqBody, _ := json.Marshal(inboundOrderRequest)
		req := httptest.NewRequest("POST", "/api/v1/inboundOrders", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/inboundOrders", handler.CreateInboundOrder(&ctx))
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.ResponseDTO
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("CreateInboundOrder fails with service error", func(t *testing.T) {
		mockService := new(serviceMock.MockInboundOrderService)
		handler := NewInboundOrderHandler(mockService)

		inboundOrderRequest := dto.InboundOrderRequestDTO{
			OrderDate:      "2024-10-01",
			OrderNumber:    "12345",
			EmployeeId:     1,
			ProductBatchId: 1,
			WarehouseId:    1,
		}

		inboundOrderModel := mappers.InboundOrderRequestDTOToModel(inboundOrderRequest)

		expectedResponse := dto.ResponseDTO{
			Code: http.StatusInternalServerError,
			Msg:  "Internal server error",
			Data: nil,
		}

		mockService.On("CreateInboundOrder", inboundOrderModel).
			Return(service.ErrInboundOrderServiceGeneric)

		reqBody, _ := json.Marshal(inboundOrderRequest)
		req := httptest.NewRequest("POST", "/api/v1/inboundOrders", bytes.NewBuffer(reqBody))
		rr := httptest.NewRecorder()
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/inboundOrders", handler.CreateInboundOrder(&ctx))
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.ResponseDTO
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
}
