package handlers_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	handler2 "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	service_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductRecordHandler_SaveProductRecord(t *testing.T) {
	ctx := context.Background()

	t.Run("Save product record success", func(t *testing.T) {
		mockProductRecordService := new(service_mock.ProductRecordServiceMock)
		mockProductRecordService.On("SaveProductRecord",
			models.ProductRecord{ProductId: 1,
				SalePrice:      29.99,
				PurchasePrice:  19.99,
				LastUpdateTime: "2029-10-01"}).Return(models.ProductRecord{ProductId: 1,
			SalePrice:      29.99,
			PurchasePrice:  19.99,
			LastUpdateTime: "2029-10-01"}, nil).Once()
		jsonBytes := []byte(`{
				"data": {
					"product_id": 1,
					"sale_price": 29.99,
					"purchase_price": 19.99,
					"last_update_date": "2029-10-01"
				}
    		}`)
		bodyExpected := `{"code":201,"message":"Product record saved","data":{"ProductId":1,"SalePrice":29.99,"PurchasePrice":19.99,"LastUpdateTime":"2029-10-01"}}`

		handler := handler2.NewProductRecordHandler(mockProductRecordService)

		router := chi.NewRouter()
		router.Post("/api/v1/productRecords", handler.SaveProductRecord(&ctx))

		rq := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/productRecords",
			bytes.NewBuffer(jsonBytes),
		)
		rc := httptest.NewRecorder()

		router.ServeHTTP(rc, rq)

		assert.Equal(t, http.StatusCreated, rc.Code)
		assert.Equal(t, bodyExpected, rc.Body.String())
		assert.Equal(t, "application/json", rc.Header().Get("Content-Type"))
	})

	t.Run("Error mapping json to model", func(t *testing.T) {
		mockProductRecordService := new(service_mock.ProductRecordServiceMock)
		jsonBytes := []byte(`{
				"data": {
					"product_id": 1,
					"sale_price": null,
					"purchase_price": 19.99,
					"last_update_date": "2029-10-01"
				}
    		}`)
		bodyExpected := `{"code":422,"message":"sale price must be greater than zero","data":null}`

		handler := handler2.NewProductRecordHandler(mockProductRecordService)

		router := chi.NewRouter()
		router.Post("/api/v1/productRecords", handler.SaveProductRecord(&ctx))

		rq := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/productRecords",
			bytes.NewBuffer(jsonBytes),
		)
		rc := httptest.NewRecorder()

		router.ServeHTTP(rc, rq)

		assert.Equal(t, http.StatusUnprocessableEntity, rc.Code)
		assert.Equal(t, bodyExpected, rc.Body.String())
		assert.Equal(t, "application/json", rc.Header().Get("Content-Type"))
	})

	t.Run("Error validation json", func(t *testing.T) {
		mockProductRecordService := new(service_mock.ProductRecordServiceMock)
		mockProductRecordService.On("SaveProductRecord",
			models.ProductRecord{}).Return(models.ProductRecord{}, nil).Once()
		bodyExpected := `{"code":500,"message":"Error, data invalid","data":null}`

		handler := handler2.NewProductRecordHandler(mockProductRecordService)

		router := chi.NewRouter()
		router.Post("/api/v1/productRecords", handler.SaveProductRecord(&ctx))

		rq := httptest.NewRequest(http.MethodPost, "/api/v1/productRecords", nil)
		rc := httptest.NewRecorder()

		router.ServeHTTP(rc, rq)

		assert.Equal(t, http.StatusInternalServerError, rc.Code)
		assert.Equal(t, bodyExpected, rc.Body.String())
		assert.Equal(t, "application/json", rc.Header().Get("Content-Type"))
	})
	t.Run("Error service", func(t *testing.T) {
		mockProductRecordService := new(service_mock.ProductRecordServiceMock)
		mockProductRecordService.On("SaveProductRecord",
			models.ProductRecord{
				ProductId:      1,
				SalePrice:      20.00,
				PurchasePrice:  19.99,
				LastUpdateTime: "2023-10-01",
			}).Return(models.ProductRecord{}, service.ErrServiceProductRecordBusinessRules)
		jsonBytes := []byte(`{
				"data": {
					"product_id": 1,
					"sale_price": 20.00,
					"purchase_price": 19.99,
					"last_update_date": "2023-10-01"
				}
    		}`)
		bodyExpected := `{"code":422,"message":"product record business error","data":null}`

		handler := handler2.NewProductRecordHandler(mockProductRecordService)

		router := chi.NewRouter()
		router.Post("/api/v1/productRecords", handler.SaveProductRecord(&ctx))

		rq := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/productRecords",
			bytes.NewBuffer(jsonBytes),
		)
		rc := httptest.NewRecorder()

		router.ServeHTTP(rc, rq)

		assert.Equal(t, http.StatusUnprocessableEntity, rc.Code)
		assert.Equal(t, bodyExpected, rc.Body.String())
		assert.Equal(t, "application/json", rc.Header().Get("Content-Type"))
	})
}

func TestProductRecordHandler_GetProductRecord(t *testing.T) {
	ctx := context.Background()

	t.Run("Get product record success", func(t *testing.T) {
		mockProductRecordService := new(service_mock.ProductRecordServiceMock)
		mockProductRecordService.On("GetProductRecord",
			mock.AnythingOfType("int")).Return(map[int]models.ProductRecordResponse{}, nil).Once()
		bodyExpected := `{"code":200,"message":"","data":{}}`

		handler := handler2.NewProductRecordHandler(mockProductRecordService)

		router := chi.NewRouter()
		router.Get("/api/v1/productRecords", handler.GetProductRecord(&ctx))

		rq := httptest.NewRequest(http.MethodGet, "/api/v1/productRecords", nil)
		rc := httptest.NewRecorder()

		router.ServeHTTP(rc, rq)

		assert.Equal(t, bodyExpected, rc.Body.String())
		assert.Equal(t, 200, rc.Code)
		assert.Equal(t, "application/json", rc.Header().Get("Content-Type"))
	})
	t.Run("Get product record success", func(t *testing.T) {
		mockProductRecordService := new(service_mock.ProductRecordServiceMock)
		mockProductRecordService.On(
			"GetProductRecord",
			mock.AnythingOfType(
				"int",
			),
		).Return(map[int]models.ProductRecordResponse{}, service.ErrServiceProductRecordNotFound).Once()
		bodyExpected := `{"code":409,"message":"product not found","data":null}`

		handler := handler2.NewProductRecordHandler(mockProductRecordService)

		router := chi.NewRouter()
		router.Get("/api/v1/productRecords", handler.GetProductRecord(&ctx))

		rq := httptest.NewRequest(http.MethodGet, "/api/v1/productRecords", nil)
		rc := httptest.NewRecorder()

		router.ServeHTTP(rc, rq)

		assert.Equal(t, bodyExpected, rc.Body.String())
		assert.Equal(t, 409, rc.Code)
		assert.Equal(t, "application/json", rc.Header().Get("Content-Type"))
	})
}
