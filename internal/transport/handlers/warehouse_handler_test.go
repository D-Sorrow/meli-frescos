package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	serviceErr "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	handlerErr "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetWarehouses(t *testing.T) {
	t.Run("find all warehouses", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		warehousesFake := map[int]models.Warehouse{}
		warehousesFake[1] = models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		serviceMock.On("GetWarehouses").Return(warehousesFake, error(nil))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/warehouses", handler.GetWarehouses())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 200,
							"message": "Warehouses got successfully",
							"data": [
								{
									"id": 1,
									"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
									"address": "Apt 1639",
									"telephone": "(639) 5350508",
									"minimum_capacity": 99,
									"minimum_temperature": -14,
									"locality_id": 1
								}
							]
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("find all warehouses fail", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()

		serviceMock.On("GetWarehouses").Return(nil, error(serviceErr.ErrWarehouseServiceDefault))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/warehouses", handler.GetWarehouses())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 500,
							"message": "internal server error",
							"data": null
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})
}

func TestGetWarehouseById(t *testing.T) {
	t.Run("find warehouse by id", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		warehouseFake := models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		serviceMock.On("GetWarehouseById", mock.Anything).Return(warehouseFake, error(nil))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/warehouses/{id}", handler.GetWarehouseById())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 200,
							"message": "Warehouse got successfully",
							"data":
								{
									"id": 1,
									"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
									"address": "Apt 1639",
									"telephone": "(639) 5350508",
									"minimum_capacity": 99,
									"minimum_temperature": -14,
									"locality_id": 1
								}
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("warehouse not found", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		serviceMock.On("GetWarehouseById", mock.Anything).Return(models.Warehouse{}, error(serviceErr.ErrWarehouseNotFound))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/warehouses/{id}", handler.GetWarehouseById())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 404,
							"message": "warehouse not found",
							"data": null
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("invalid warehouse id", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		serviceMock.On("GetWarehouseById", mock.Anything).Return(models.Warehouse{}, error(handlerErr.ErrWarehouseIdNotValid))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/warehouses/{id}", handler.GetWarehouseById())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/m", nil)
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 400,
							"message": "invalid id",
							"data": null
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestCreateWarehouse(t *testing.T) {
	t.Run("create warehouse ok", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		warehouseFake := models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		serviceMock.On("CreateWarehouse", mock.Anything).Return(warehouseFake, error(nil))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/warehouses", handler.CreateWarehouse())

		res := httptest.NewRecorder()

		reqBody := `{
					"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
					"address": "Apt 1639",
					"telephone": "(639) 5350508",
					"minimun_capacity": 99,
					"minimun_temperature": -14,
					"locality_id": 1
					}`
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewBuffer([]byte(reqBody)))
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 201,
							"message": "Warehouse created successsfully",
							"data": 
							{
								"id": 1,
								"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
								"address": "Apt 1639",
								"telephone": "(639) 5350508",
								"minimum_capacity": 99,
								"minimum_temperature": -14,
								"locality_id": 1
							}
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("create warehouse fail - incorrect body", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		serviceMock.On("CreateWarehouse", mock.Anything).Return(nil, error(nil))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/warehouses", handler.CreateWarehouse())

		res := httptest.NewRecorder()

		reqBody := `{
					"address": "Apt 1639",
					"telephone": "(639) 5350508",
					"minimun_capacity": 99,
					"minimun_temperature": -14,
					"locality_id": 1
					}`
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewBuffer([]byte(reqBody)))
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusUnprocessableEntity
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 422,
							"message": "warehouse code is required",
							"data": null
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("create warehouse fail - conflic", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		serviceMock.On("CreateWarehouse", mock.Anything).Return(models.Warehouse{}, error(serviceErr.ErrWarehouseCodeDuplicate))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/warehouses", handler.CreateWarehouse())

		res := httptest.NewRecorder()

		reqBody := `{
					"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
					"address": "Apt 1639",
					"telephone": "(639) 5350508",
					"minimun_capacity": 99,
					"minimun_temperature": -14,
					"locality_id": 1
					}`
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", bytes.NewBuffer([]byte(reqBody)))
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusConflict
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 409,
							"message": "warehouse code already exists",
							"data": null
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})
}

func TestPatchWarehouse(t *testing.T) {
	t.Run("update warehouse ok", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		warehouseFake := models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		serviceMock.On("PatchWarehouse", mock.Anything, mock.Anything).Return(warehouseFake, error(nil))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/api/v1/warehouses/{id}", handler.PatchWarehouse())

		res := httptest.NewRecorder()

		reqBody := `{
					"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
					"address": "Apt 1639",
					"telephone": "(639) 5350508",
					"minimum_capacity": 99,
					"minimum_temperature": -14,
					"locality_id": 1
					}`
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", bytes.NewBuffer([]byte(reqBody)))
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 200,
							"message": "Warehouse updated",
							"data": 
							{
								"id": 1,
								"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
								"address": "Apt 1639",
								"telephone": "(639) 5350508",
								"minimum_capacity": 99,
								"minimum_temperature": -14,
								"locality_id": 1
							}
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})

	t.Run("update warehouse not found", func(t *testing.T) {
		serviceMock := service.NewWarehouseServiceMock()
		serviceMock.On("PatchWarehouse", mock.Anything, mock.Anything).Return(models.Warehouse{}, error(serviceErr.ErrWarehouseNotFound))
		handler := NewWarehouseHandler(serviceMock)
		router := chi.NewRouter()
		router.Patch("/api/v1/warehouses/{id}", handler.PatchWarehouse())

		res := httptest.NewRecorder()
		reqBody := `{
			"warehouse_code": "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			"address": "Apt 1639",
			"telephone": "(639) 5350508",
			"minimum_capacity": 99,
			"minimum_temperature": -14,
			"locality_id": 1
			}`
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", bytes.NewBuffer([]byte(reqBody)))
		router.ServeHTTP(res, req)

		expectedStatusCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
							"code": 404,
							"message": "warehouse not found",
							"data": null
						}`
		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		serviceMock.AssertExpectations(t)
	})
}
