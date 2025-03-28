package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	fakeModels "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/models"
	service_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/service"
	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateLocality(t *testing.T) {
	ctx := context.Background()

	t.Run("create locality", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("CreateLocality", fakeModels.LocalityRequest).
			Return(fakeModels.LocalityRequest, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/localities/", handlerImp.CreateLocality(&ctx))

		localityJSON, errMarshal := json.Marshal(fakeModels.JsonLocalityCreatedDto)
		if errMarshal != nil {
			t.Error(errMarshal)
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/localities/",
			bytes.NewBuffer(localityJSON),
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusCreated

		expectedBody := map[string]any{
			"code":    http.StatusCreated,
			"message": "Locality Created",
			"data":    fakeModels.JsonLocalityCreatedDto,
		}

		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "CreateLocality", fakeModels.LocalityRequest)
	})

	t.Run("create locality fail - Bad Request BODY incomplete", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("CreateLocality", fakeModels.LocalityRequest).
			Return(fakeModels.LocalityBadRequest, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/localities/", handlerImp.CreateLocality(&ctx))

		localityJSON, errMarshal := json.Marshal(fakeModels.JsonBadLocalityCreatedDto)
		if errMarshal != nil {
			t.Error(errMarshal)
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/localities/",
			bytes.NewBuffer(localityJSON),
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusUnprocessableEntity

		expectedBody := map[string]any{
			"code":    http.StatusUnprocessableEntity,
			"message": "The request is invalid because it does not contain the necessary fields: 'LocalityName' is required and must be a string, 'ProvinceName' is required and must be a string, ",
			"data":    nil,
		}

		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNumberOfCalls(t, "CreateLocality", 0)
	})

	t.Run("create locality fail - Bad Request BODY json bad structure", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("CreateLocality", fakeModels.LocalityBadRequest).
			Return(fakeModels.LocalityBadRequest, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/localities/", handlerImp.CreateLocality(&ctx))

		localityJSON, errMarshal := json.Marshal("{,make:'dsad',}")
		//panic(bytes.NewBuffer(localityJSON))
		if errMarshal != nil {
			t.Error(errMarshal)
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/localities/",
			bytes.NewBuffer(localityJSON),
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusBadRequest

		require.Equal(t, expectedCode, res.Code)
		serviceMock.AssertNumberOfCalls(t, "CreateLocality", 0)
	})

	t.Run("create locality fail - Locality already exist", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("CreateLocality", fakeModels.LocalityRequest).
			Return(models.Locality{}, service.ErrLocalityAlreadyExists)
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Post("/api/v1/localities/", handlerImp.CreateLocality(&ctx))

		localityJSON, errMarshal := json.Marshal(fakeModels.JsonLocalityCreatedDto)
		if errMarshal != nil {
			t.Error(errMarshal)
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/localities/",
			bytes.NewBuffer(localityJSON),
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusConflict

		expectedBody := map[string]any{
			"code":    http.StatusConflict,
			"message": "locality already exist",
			"data":    nil,
		}

		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNumberOfCalls(t, "CreateLocality", 1)
	})

}

func TestGetSellersByLocality(t *testing.T) {
	ctx := context.Background()

	t.Run("get sellers by locality", func(t *testing.T) {
		localityId := 1

		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetSellersByLocality", localityId).
			Return(fakeModels.LocalitySellersResponse, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/api/v1/localities/reportSellers", handlerImp.GetSellersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprint("/api/v1/localities/reportSellers?id=", localityId),
			nil,
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusOK

		expectedBody := map[string]any{
			"code":    http.StatusOK,
			"message": "success",
			"data":    fakeModels.JsonLocalitySellersResponse,
		}

		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "GetSellersByLocality", localityId)
	})

	t.Run("get sellers by locality - Bad Request query params", func(t *testing.T) {
		localityId := "a"

		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetSellersByLocality", localityId).
			Return(fakeModels.LocalitySellersResponse, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/api/v1/localities/reportSellers", handlerImp.GetSellersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprint("/api/v1/localities/reportSellers?id=", localityId),
			nil,
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusBadRequest

		expectedBody := map[string]any{
			"code":    http.StatusBadRequest,
			"message": "id must be a number",
			"data":    nil,
		}

		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNumberOfCalls(t, "GetSellersByLocality", 0)
	})

	t.Run("get sellers by locality - Locality Not Found", func(t *testing.T) {
		localityId := 1

		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetSellersByLocality", localityId).
			Return(models.LocalitySellers{}, service.ErrLocalityNotFound)
		handlerImp := NewLocalityHandler(serviceMock)

		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/api/v1/localities/reportSellers", handlerImp.GetSellersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprint("/api/v1/localities/reportSellers?id=", localityId),
			nil,
		)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusNotFound

		expectedBody := map[string]any{
			"code":    http.StatusNotFound,
			"message": "locality id not found",
			"data":    nil,
		}

		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNumberOfCalls(t, "GetSellersByLocality", 1)
	})

}

func TestGetCarriersByLocality(t *testing.T) {
	ctx := context.Background()

	t.Run(" get carriers by all localities", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetCarriersByAllLocalities").Return(fakeModels.LocalityCarriers, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/reportCarries", nil)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusOK

		expectedBody := map[string]any{
			"code":    http.StatusOK,
			"message": "success",
			"data":    fakeModels.ResponseLocalityCarriersDto,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "GetCarriersByAllLocalities")
	})

	t.Run(" get carriers by all localities - Fail", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetCarriersByAllLocalities").Return(nil, service.ErrGetAllLocalities)
		handlerImp := NewLocalityHandler(serviceMock)
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/reportCarries", nil)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusInternalServerError

		expectedBody := map[string]any{
			"code":    http.StatusInternalServerError,
			"message": "internal server error",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "GetCarriersByAllLocalities")
	})

	t.Run(" get carriers by locality id", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetCarriersByLocality", mock.Anything).
			Return(fakeModels.LocalityCarriers[0], error(nil))
		handlerImp := NewLocalityHandler(serviceMock)
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/reportCarries?id=1", nil)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusOK

		expectedBody := map[string]any{
			"code":    http.StatusOK,
			"message": "success",
			"data":    fakeModels.ResponseLocalityCarriersDto[0],
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "GetCarriersByLocality", mock.Anything)
	})

	t.Run(" get carriers by locality id fail - wrond id", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetCarriersByLocality", mock.Anything).
			Return(fakeModels.LocalityCarriers[0], error(nil))
		handlerImp := NewLocalityHandler(serviceMock)
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/reportCarries?id=m", nil)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusBadRequest

		expectedBody := map[string]any{
			"code":    http.StatusBadRequest,
			"message": "id must be a number",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNumberOfCalls(t, "GetCarriersByLocality", 0)
	})
	t.Run(" get carriers by locality id fail - id not found", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetCarriersByLocality", mock.Anything).
			Return(models.LocalityCarriers{}, service.ErrLocalityNotFound)
		handlerImp := NewLocalityHandler(serviceMock)
		app, err := fury.NewWebApplication()
		if err != nil {
			t.Fatal(err)
		}

		router := app.Router
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/reportCarries?id=1", nil)

		router.ServeHTTP(res, req)
		expectedCode := http.StatusNotFound

		expectedBody := map[string]any{
			"code":    http.StatusNotFound,
			"message": "locality id not found",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNumberOfCalls(t, "GetCarriersByLocality", 1)
	})
}
