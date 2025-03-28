package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	service2 "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	modelsMock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/models"
	service_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/service"
	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name               string
	requestPayload     interface{}
	mockServiceReturn  []interface{}
	expectedStatusCode int
	expectedCalls      int
}

func TestProductHandler_SaveProduct(t *testing.T) {
	ctx := context.Background()

	testCases := []testCase{
		{
			name:           "Success",
			requestPayload: modelsMock.ReturnMockProduct(),
			mockServiceReturn: []interface{}{
				nil,
			},
			expectedStatusCode: http.StatusCreated,
			expectedCalls:      1,
		},
		{
			name:           "Fail: No Payload",
			requestPayload: nil,
			mockServiceReturn: []interface{}{
				nil,
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedCalls:      0,
		},
		{
			name:           "Conflict: Product Already Exists",
			requestPayload: modelsMock.ReturnMockProduct(),
			mockServiceReturn: []interface{}{
				service2.ErrServiceProductAlreadyExists,
			},
			expectedStatusCode: http.StatusConflict,
			expectedCalls:      1,
		},
		{

			name:           "Field invalid",
			requestPayload: modelsMock.ReturnProductModelMap()[2],
			mockServiceReturn: []interface{}{
				nil,
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedCalls:      0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service_mock.ProductServiceMock)
			if tc.mockServiceReturn == nil {
				mockService.On("SaveProduct", modelsMock.ReturnMockProductModel()).
					Return(tc.mockServiceReturn[0])
			} else {
				mockService.On("SaveProduct", mock.Anything).Return(tc.mockServiceReturn[0])
			}

			handler := handlers.NewProductHandler(mockService)
			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Post("/api/v1/products", handler.SaveProduct(&ctx))

			var productJSON []byte
			if tc.requestPayload != nil {
				var errMarshal error
				productJSON, errMarshal = json.Marshal(tc.requestPayload)
				if errMarshal != nil {
					t.Error(errMarshal)
				}
			}

			req := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(productJSON))
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatusCode, rec.Code)
			mockService.AssertNumberOfCalls(t, "SaveProduct", tc.expectedCalls)
		})
	}
}
func TestProductHandler_GetProductByID(t *testing.T) {
	ctx := context.Background()

	testCases := []testCase{
		{
			name:           "GetProductsById_NonExistent",
			requestPayload: 1,
			mockServiceReturn: []interface{}{
				models.Product{},
				service2.ErrServiceProductNotFound,
			},
			expectedStatusCode: http.StatusNotFound,
			expectedCalls:      1,
		},
		{
			name:           "GetProductsById_Existent",
			requestPayload: 1,
			mockServiceReturn: []interface{}{
				modelsMock.ReturnMockProductModel(),
				nil,
			},
			expectedStatusCode: http.StatusOK,
			expectedCalls:      1,
		},
		{
			name:               "GetProductsById_InvalidId",
			requestPayload:     "1A",
			mockServiceReturn:  nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service_mock.ProductServiceMock)
			if tc.mockServiceReturn != nil {
				mockService.On("GetProductByID", 1).
					Return(tc.mockServiceReturn[0], tc.mockServiceReturn[1])
			}

			handler := handlers.NewProductHandler(mockService)

			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Get("/api/v1/products/{id}", handler.GetProductByID(&ctx))

			req := httptest.NewRequest(
				"GET",
				fmt.Sprint("/api/v1/products/", tc.requestPayload),
				nil,
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatusCode, rec.Code)
			mockService.AssertNumberOfCalls(t, "GetProductByID", tc.expectedCalls)
		})
	}
}
func TestProductHandler_GetProducts(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	mockService.On("GetProducts").Return(modelsMock.ReturnProductModelMap(), nil)

	handler := handlers.NewProductHandler(mockService)

	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router
	router.Get("/api/v1/products", handler.GetProducts(&ctx))

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertNumberOfCalls(t, "GetProducts", 1)
}
func TestProductHandler_GetProducts_Err(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	mockService.On("GetProducts").
		Return(modelsMock.ReturnProductModelMap(), service2.ErrServiceProductUnknown)

	handler := handlers.NewProductHandler(mockService)

	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router
	router.Get("/api/v1/products", handler.GetProducts(&ctx))

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertNumberOfCalls(t, "GetProducts", 1)
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	mockService.On("UpdateProduct", mock.Anything, mock.Anything).
		Return(modelsMock.ReturnMockProductModel(), nil)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	attributeJSON, _ := json.Marshal(modelsMock.ReturnAttributesModel())
	router.Patch("/api/v1/products/{id}", handler.UpdateProduct(&ctx))
	req := httptest.NewRequest("PATCH", "/api/v1/products/1", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestProductHandler_UpdateProduct_NonExistent(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	mockService.On("UpdateProduct", mock.Anything, mock.Anything).
		Return(modelsMock.ReturnMockProductModel(), service2.ErrServiceProductNotFound)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	attributeJSON, _ := json.Marshal(modelsMock.ReturnAttributesModel())
	router.Patch("/api/v1/products/{id}", handler.UpdateProduct(&ctx))
	req := httptest.NewRequest("PATCH", "/api/v1/products/1", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
func TestProductHandler_UpdateProduct_IdInvalid(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	attributeJSON, _ := json.Marshal(modelsMock.ReturnAttributesModel())
	router.Patch("/api/v1/products/{id}", handler.UpdateProduct(&ctx))
	req := httptest.NewRequest("PATCH", "/api/v1/products/1A", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestProductHandler_UpdateProduct_DecoderErr(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	router.Patch("/api/v1/products/{id}", handler.UpdateProduct(&ctx))
	req := httptest.NewRequest(
		"PATCH",
		"/api/v1/products/1",
		bytes.NewBuffer([]byte("invalid_json")),
	)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockService.AssertNumberOfCalls(t, "UpdateProduct", 0)
}

func TestProductHandler_UpdateProduct_ErrValidation(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	att := modelsMock.ReturnAttributesModel()
	att.FreezingRate = 10

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	attributeJSON, _ := json.Marshal(att)
	router.Patch("/api/v1/products/{id}", handler.UpdateProduct(&ctx))
	req := httptest.NewRequest("PATCH", "/api/v1/products/1", bytes.NewBuffer(attributeJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockService.AssertNumberOfCalls(t, "UpdateProduct", 0)
}

func TestProductHandler_DeleteProduct_NonExistent(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	mockService.On("DeleteProduct", 1).Return(service2.ErrServiceProductNotFound)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	router.Delete("/api/v1/products/{id}", handler.DeleteProduct(&ctx))
	req := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
func TestProductHandler_DeleteProduct_Existent(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)
	mockService.On("DeleteProduct", 1).Return(nil)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	router.Delete("/api/v1/products/{id}", handler.DeleteProduct(&ctx))
	req := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestProductHandler_DeleteProduct_IdInvalid(t *testing.T) {
	ctx := context.Background()

	mockService := new(service_mock.ProductServiceMock)

	handler := handlers.NewProductHandler(mockService)
	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	router := app.Router

	router.Delete("/api/v1/products/{id}", handler.DeleteProduct(&ctx))
	req := httptest.NewRequest("DELETE", "/api/v1/products/1A", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
