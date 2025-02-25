package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	service2 "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	modelsMock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/models"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductHandler_SaveProduct_Success(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("SaveProduct", modelsMock.ReturnMockProductModel()).Return(nil)
	productJSON, errMarshal := json.Marshal(modelsMock.ReturnMockProduct())
	if errMarshal != nil {
		t.Error(errMarshal)
	}

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Post("/api/v1/products", handler.SaveProduct())

	req := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(productJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertNumberOfCalls(t, "SaveProduct", 1)
}
func TestProductHandler_SaveProduct_Fail(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("SaveProduct", models.Product{}).Return(service2.ErrServiceProductAlreadyExists)

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Post("/api/v1/products", handler.SaveProduct())

	req := httptest.NewRequest("POST", "/api/v1/products", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	mockService.AssertNumberOfCalls(t, "SaveProduct", 0)
}
func TestProductHandler_SaveProduct_Conflict(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("SaveProduct", modelsMock.ReturnMockProductModel()).Return(service2.ErrServiceProductAlreadyExists)
	productJSON, errMarshal := json.Marshal(modelsMock.ReturnMockProduct())
	if errMarshal != nil {
		t.Error(errMarshal)
	}

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Post("/api/v1/products", handler.SaveProduct())

	req := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(productJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusConflict, rec.Code)
	mockService.AssertNumberOfCalls(t, "SaveProduct", 1)
}

func TestProductHandler_GetProducts(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("GetProducts").Return(modelsMock.ReturnProductModelMap(), nil)

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Get("/api/v1/products", handler.GetProducts())

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertNumberOfCalls(t, "GetProducts", 1)
}

func TestProductHandler_GetProductsById_NonExistent(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("GetProductByID", 1).Return(models.Product{}, service2.ErrServiceProductNotFound)

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Get("/api/v1/products/{id}", handler.GetProductByID())

	req := httptest.NewRequest("GET", "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestProductHandler_GetProductsById_Existent(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("GetProductByID", 1).Return(modelsMock.ReturnMockProductModel(), nil)

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Get("/api/v1/products/{id}", handler.GetProductByID())

	req := httptest.NewRequest("GET", "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestProductHandler_GetProductsById_IdInvalid(t *testing.T) {
	mockService := new(service.ProductServiceMock)

	handler := handlers.NewProductHandler(mockService)

	router := chi.NewRouter()
	router.Get("/api/v1/products/{id}", handler.GetProductByID())

	req := httptest.NewRequest("GET", "/api/v1/products/1A", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestProductHandler_UpdateProduct(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("UpdateProduct", mock.Anything, mock.Anything).Return(modelsMock.ReturnMockProductModel(), nil)

	handler := handlers.NewProductHandler(mockService)
	router := chi.NewRouter()

	attributeJSON, _ := json.Marshal(modelsMock.ReturnAttributesModel())
	router.Put("/api/v1/products/{id}", handler.UpdateProduct())
	req := httptest.NewRequest("PUT", "/api/v1/products/1", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestProductHandler_UpdateProduct_NonExistent(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("UpdateProduct", mock.Anything, mock.Anything).Return(modelsMock.ReturnMockProductModel(), service2.ErrServiceProductNotFound)

	handler := handlers.NewProductHandler(mockService)
	router := chi.NewRouter()

	attributeJSON, _ := json.Marshal(modelsMock.ReturnAttributesModel())
	router.Put("/api/v1/products/{id}", handler.UpdateProduct())
	req := httptest.NewRequest("PUT", "/api/v1/products/1", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
func TestProductHandler_UpdateProduct_IdInvalid(t *testing.T) {
	mockService := new(service.ProductServiceMock)

	handler := handlers.NewProductHandler(mockService)
	router := chi.NewRouter()

	attributeJSON, _ := json.Marshal(modelsMock.ReturnAttributesModel())
	router.Put("/api/v1/products/{id}", handler.UpdateProduct())
	req := httptest.NewRequest("PUT", "/api/v1/products/1A", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestProductHandler_DeleteProduct_NonExistent(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("DeleteProduct", 1).Return(service2.ErrServiceProductNotFound)

	handler := handlers.NewProductHandler(mockService)
	router := chi.NewRouter()

	router.Delete("/api/v1/products/{id}", handler.DeleteProduct())
	req := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
func TestProductHandler_DeleteProduct_Existent(t *testing.T) {
	mockService := new(service.ProductServiceMock)
	mockService.On("DeleteProduct", 1).Return(nil)

	handler := handlers.NewProductHandler(mockService)
	router := chi.NewRouter()

	router.Delete("/api/v1/products/{id}", handler.DeleteProduct())
	req := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestProductHandler_DeleteProduct_IdInvalid(t *testing.T) {
	mockService := new(service.ProductServiceMock)

	handler := handlers.NewProductHandler(mockService)
	router := chi.NewRouter()

	router.Delete("/api/v1/products/{id}", handler.DeleteProduct())
	req := httptest.NewRequest("DELETE", "/api/v1/products/1A", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
