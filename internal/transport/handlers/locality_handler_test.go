package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	fakeModels "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/models"
	service_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetCarriersByLocality(t *testing.T) {
	t.Run(" get carriers by all localities", func(t *testing.T) {
		serviceMock := service_mock.NewLocalityServiceMock()
		serviceMock.On("GetCarriersByAllLocalities").Return(fakeModels.LocalityCarriers, error(nil))
		handlerImp := NewLocalityHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality())

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
		router := chi.NewRouter()
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality())

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
		serviceMock.On("GetCarriersByLocality", mock.Anything).Return(fakeModels.LocalityCarriers[0], error(nil))
		handlerImp := NewLocalityHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality())

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
		serviceMock.On("GetCarriersByLocality", mock.Anything).Return(fakeModels.LocalityCarriers[0], error(nil))
		handlerImp := NewLocalityHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality())

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
		serviceMock.On("GetCarriersByLocality", mock.Anything).Return(models.LocalityCarriers{}, service.ErrLocalityNotFound)
		handlerImp := NewLocalityHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/reportCarries", handlerImp.GetCarriersByLocality())

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
