package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var SellerDtoFake dto.SellerDto = dto.SellerDto{
	Id:          1,
	Cid:         1,
	CompanyName: "company_fake name",
	Address:     "Calle fake 12312",
	Telephone:   "312346456",
	LocalityId:  1,
}

var SellerDtoFakeBad dto.SellerDto = dto.SellerDto{
	Id:         1,
	Cid:        1,
	Telephone:  "312346456",
	LocalityId: 1,
}

var telephone = "123123"
var SellerUpdateDtoFake dto.SellerUpdateDto = dto.SellerUpdateDto{
	Telephone: &telephone,
}

var SellerPatch models.SellerPatch = models.SellerPatch{
	Telephone: &telephone,
}

var SellerFake models.Seller = models.Seller{
	Id:          1,
	Cid:         1,
	CompanyName: "company_fake name",
	Address:     "Calle fake 12312",
	Telephone:   "312346456",
	LocalityId:  1,
}

type testCase struct {
	name               string
	requestUrlParams   interface{}
	requestPayloadBody interface{}
	expectedStatusCode int
	expectedCalls      int
	expectedResponse   dto.ResponseDTO
	testCaseServiceMock
}

type testCaseServiceMock struct {
	responseServiceMock interface{}
	requestServiceMock  interface{}
	errorServiceMock    error
}

func TestSellerHandlerCreate(t *testing.T) {

	testCases := []testCase{
		{name: "Create Seller - Successfully",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerFake,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestPayloadBody: SellerDtoFake,
			expectedStatusCode: http.StatusCreated,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Create Seller - Bad Request",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  models.Seller{},
				responseServiceMock: models.Seller{},
				errorServiceMock:    nil,
			},
			requestPayloadBody: nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Create Seller - Fail miss necesary fields",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerFake,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestPayloadBody: SellerDtoFakeBad,
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Create Seller - Conflict Seller already exist",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerFake,
				responseServiceMock: models.Seller{},
				errorServiceMock:    service_errors.ErrSellerAlreadyExists,
			},
			requestPayloadBody: SellerDtoFake,
			expectedStatusCode: http.StatusConflict,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Create Seller - Bad Request - invalid JSON structure",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerFake,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestPayloadBody: "{make:'dsad',}",
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service.SellerServiceMock)

			mockService.On("CreateSeller", tc.requestServiceMock).Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			router := chi.NewRouter()
			router.Post("/api/v1/sellers", handler.CreateSeller())

			var sellerJSON []byte
			if tc.requestPayloadBody != nil {
				var errMarshal error
				sellerJSON, errMarshal = json.Marshal(tc.requestPayloadBody)
				if errMarshal != nil {
					t.Error(errMarshal)
				}
			}

			req := httptest.NewRequest("POST", "/api/v1/sellers", bytes.NewBuffer(sellerJSON))
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "CreateSeller", tc.expectedCalls)
		})
	}

}

func TestSellerHandlerReadAll(t *testing.T) {
	sellersFakeMap := make(map[int]models.Seller)
	sellersFakeMap[1] = SellerFake

	testCases := []testCase{
		{name: "Find All - Successfully",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  nil,
				responseServiceMock: sellersFakeMap,
				errorServiceMock:    nil,
			},
			requestPayloadBody: nil,
			expectedStatusCode: http.StatusOK,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Find All - Not Found",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  nil,
				responseServiceMock: map[int]models.Seller{},
				errorServiceMock:    service_errors.ErrSellerNotFound,
			},
			requestPayloadBody: nil,
			expectedStatusCode: http.StatusNotFound,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service.SellerServiceMock)

			mockService.On("GetSellers").Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			router := chi.NewRouter()
			router.Get("/api/v1/sellers", handler.GetSellers())

			req := httptest.NewRequest("GET", "/api/v1/sellers", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "GetSellers", tc.expectedCalls)
		})
	}

}

func TestSellerHandlerRead(t *testing.T) {
	testCases := []testCase{
		{name: "Find By ID - Seller No Exists",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  1000,
				responseServiceMock: models.Seller{},
				errorServiceMock:    service_errors.ErrSellerNotFound,
			},
			requestPayloadBody: nil,
			requestUrlParams:   1000,
			expectedStatusCode: http.StatusNotFound,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Find By ID - Successfully",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  1,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestUrlParams:   1,
			expectedStatusCode: http.StatusOK,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Find By ID - Bad Request Param URL",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  1000,
				responseServiceMock: models.Seller{},
				errorServiceMock:    service_errors.ErrSellerNotFound,
			},
			requestUrlParams:   "asd",
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service.SellerServiceMock)

			mockService.On("GetSellerById", tc.requestServiceMock).Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			router := chi.NewRouter()
			router.Get("/api/v1/sellers/{id}", handler.GetSeller())

			req := httptest.NewRequest("GET", fmt.Sprint("/api/v1/sellers/", tc.requestUrlParams), nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "GetSellerById", tc.expectedCalls)
		})
	}

}

func TestSellerHandlerUpdate(t *testing.T) {
	testCases := []testCase{
		{name: "Update Seller by Id - Successfully",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerPatch,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestUrlParams:   1,
			requestPayloadBody: SellerUpdateDtoFake,
			expectedStatusCode: http.StatusOK,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Update Seller by Id - Seller not exist",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerPatch,
				responseServiceMock: models.Seller{},
				errorServiceMock:    service_errors.ErrSellerNotFound,
			},
			requestUrlParams:   1,
			requestPayloadBody: SellerUpdateDtoFake,
			expectedStatusCode: http.StatusNotFound,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Update Seller by Id - Bad Request - invalid JSON structure",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerPatch,
				responseServiceMock: models.Seller{},
				errorServiceMock:    nil,
			},
			requestUrlParams:   1,
			requestPayloadBody: "{make:'dsad',}",
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Update Seller by Id - Bad Request Fail data type fields",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerFake,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestUrlParams:   1,
			requestPayloadBody: "{telephone:123}",
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Update Seller by Id - Bad Request Param URL",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  SellerFake,
				responseServiceMock: SellerFake,
				errorServiceMock:    nil,
			},
			requestUrlParams:   "asd",
			requestPayloadBody: "{telephone:123}",
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service.SellerServiceMock)

			mockService.On("UpdateSeller", tc.requestUrlParams, tc.requestServiceMock).Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			router := chi.NewRouter()
			router.Patch("/api/v1/sellers/{id}", handler.UpdateSeller())

			var sellerJSON []byte
			if tc.requestPayloadBody != nil {
				var errMarshal error
				sellerJSON, errMarshal = json.Marshal(tc.requestPayloadBody)
				if errMarshal != nil {
					t.Error(errMarshal)
				}
			}

			req := httptest.NewRequest("PATCH", fmt.Sprint("/api/v1/sellers/", tc.requestUrlParams), bytes.NewBuffer(sellerJSON))
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "UpdateSeller", tc.expectedCalls)
		})
	}
}

func TestSellerHandlerDelete(t *testing.T) {
	testCases := []testCase{
		{name: "Delete Seller by Id - Successfully",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  1,
				responseServiceMock: nil,
				errorServiceMock:    nil,
			},
			requestUrlParams:   1,
			expectedStatusCode: http.StatusNoContent,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Delete Seller by Id - Seller not exist",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  1000,
				responseServiceMock: nil,
				errorServiceMock:    service_errors.ErrSellerNotFound,
			},
			requestUrlParams:   1000,
			expectedStatusCode: http.StatusNotFound,
			expectedCalls:      1,
			expectedResponse:   dto.ResponseDTO{},
		},
		{name: "Delete Seller by ID - Bad Request Param URL",
			testCaseServiceMock: testCaseServiceMock{
				requestServiceMock:  1000,
				responseServiceMock: models.Seller{},
				errorServiceMock:    service_errors.ErrSellerNotFound,
			},
			requestUrlParams:   "asd",
			expectedStatusCode: http.StatusBadRequest,
			expectedCalls:      0,
			expectedResponse:   dto.ResponseDTO{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(service.SellerServiceMock)

			mockService.On("DeleteSeller", tc.requestUrlParams).Return(tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			router := chi.NewRouter()
			router.Delete("/api/v1/sellers/{id}", handler.DeleteSeller())

			req := httptest.NewRequest("DELETE", fmt.Sprint("/api/v1/sellers/", tc.requestUrlParams), nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "DeleteSeller", tc.expectedCalls)
		})
	}
}
