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
	service_errors "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	service_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/service"
	"github.com/melisource/fury_go-platform/pkg/fury"
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

type testCaseSeller struct {
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
	ctx := context.Background()

	testCases := []testCaseSeller{
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
			mockService := new(service_mock.SellerServiceMock)

			mockService.On("CreateSeller", tc.requestServiceMock).
				Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Post("/api/v1/sellers", handler.CreateSeller(&ctx))

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
	ctx := context.Background()
	sellersFakeMap := make(map[int]models.Seller)
	sellersFakeMap[1] = SellerFake

	testCases := []testCaseSeller{
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
			mockService := new(service_mock.SellerServiceMock)

			mockService.On("GetSellers").Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Get("/api/v1/sellers", handler.GetSellers(&ctx))

			req := httptest.NewRequest("GET", "/api/v1/sellers", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "GetSellers", tc.expectedCalls)
		})
	}

}

func TestSellerHandlerRead(t *testing.T) {
	ctx := context.Background()

	testCases := []testCaseSeller{
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
			mockService := new(service_mock.SellerServiceMock)

			mockService.On("GetSellerById", tc.requestServiceMock).
				Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Get("/api/v1/sellers/{id}", handler.GetSeller(&ctx))

			req := httptest.NewRequest(
				"GET",
				fmt.Sprint("/api/v1/sellers/", tc.requestUrlParams),
				nil,
			)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "GetSellerById", tc.expectedCalls)
		})
	}

}

func TestSellerHandlerUpdate(t *testing.T) {
	ctx := context.Background()

	testCases := []testCaseSeller{
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
			mockService := new(service_mock.SellerServiceMock)

			mockService.On("UpdateSeller", tc.requestUrlParams, tc.requestServiceMock).
				Return(tc.responseServiceMock, tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Patch("/api/v1/sellers/{id}", handler.UpdateSeller(&ctx))

			var sellerJSON []byte
			if tc.requestPayloadBody != nil {
				var errMarshal error
				sellerJSON, errMarshal = json.Marshal(tc.requestPayloadBody)
				if errMarshal != nil {
					t.Error(errMarshal)
				}
			}

			req := httptest.NewRequest(
				"PATCH",
				fmt.Sprint("/api/v1/sellers/", tc.requestUrlParams),
				bytes.NewBuffer(sellerJSON),
			)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "UpdateSeller", tc.expectedCalls)
		})
	}
}

func TestSellerHandlerDelete(t *testing.T) {
	ctx := context.Background()

	testCases := []testCaseSeller{
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
			mockService := new(service_mock.SellerServiceMock)

			mockService.On("DeleteSeller", tc.requestUrlParams).Return(tc.errorServiceMock)

			handler := handlers.NewHandlerService(mockService)
			app, err := fury.NewWebApplication()
			if err != nil {
				t.Fatal(err)
			}

			router := app.Router
			router.Delete("/api/v1/sellers/{id}", handler.DeleteSeller(&ctx))

			req := httptest.NewRequest(
				"DELETE",
				fmt.Sprint("/api/v1/sellers/", tc.requestUrlParams),
				nil,
			)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			mockService.AssertNumberOfCalls(t, "DeleteSeller", tc.expectedCalls)
		})
	}
}
