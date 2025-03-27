package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	serviceErr "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	fakeModels "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/models"
	service_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetAllCarriers(t *testing.T) {
	ctx := context.Background()

	t.Run("find all carriers", func(t *testing.T) {
		serviceMock := service_mock.NewCarryServiceMock()
		serviceMock.On("GetAllCarriers").Return(fakeModels.Carriers, error(nil))
		handlerImp := NewCarryHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/carrier", handlerImp.GetAllCarriers(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/carrier", nil)
		router.ServeHTTP(res, req)

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := map[string]any{
			"code":    200,
			"message": "Carriers got successfully",
			"data":    fakeModels.ResponseCarrierDto,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "GetAllCarriers")
	})

	t.Run("find all carriers fail", func(t *testing.T) {
		serviceMock := service_mock.NewCarryServiceMock()
		serviceMock.On("GetAllCarriers").Return(nil, serviceErr.ErrCarrierServiceDefault)
		handlerImp := NewCarryHandler(serviceMock)
		router := chi.NewRouter()
		router.Get("/api/v1/carrier", handlerImp.GetAllCarriers(&ctx))

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/carrier", nil)
		router.ServeHTTP(res, req)

		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := map[string]any{
			"code":    500,
			"message": "Server Error",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "GetAllCarriers")
	})
}

func TestCreateCarrier(t *testing.T) {
	ctx := context.Background()

	t.Run("create warehouse ok", func(t *testing.T) {
		serviceMock := service_mock.NewCarryServiceMock()
		fakeCarrier := fakeModels.Carriers[0]
		serviceMock.On("CreateCarrier", mock.Anything).Return(fakeCarrier, error(nil))
		handlerImp := NewCarryHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/carrier", handlerImp.CreateCarrier(&ctx))

		res := httptest.NewRecorder()
		reqBody := `{
						"cid": "#086",
						"company_name": "Livetube",
						"address": "Room 192",
						"telephone": "374-776-3015",
						"locality_id": 1
					}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/carrier",
			bytes.NewBuffer([]byte(reqBody)),
		)
		router.ServeHTTP(res, req)

		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := map[string]any{
			"code":    http.StatusCreated,
			"message": "carrier created successsfully",
			"data":    fakeModels.ResponseCarrierDto[0],
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "CreateCarrier", mock.Anything)
	})

	t.Run("create warehouse fail - incomplete body", func(t *testing.T) {
		serviceMock := service_mock.NewCarryServiceMock()
		serviceMock.On("CreateCarrier", mock.Anything).Return(nil, error(nil))
		handlerImp := NewCarryHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/carrier", handlerImp.CreateCarrier(&ctx))

		res := httptest.NewRecorder()
		reqBody := `{
						"company_name": "Livetube",
						"address": "Room 192",
						"telephone": "374-776-3015",
						"locality_id": 1
					}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/carrier",
			bytes.NewBuffer([]byte(reqBody)),
		)
		router.ServeHTTP(res, req)

		expectedCode := http.StatusUnprocessableEntity
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := map[string]any{
			"code":    http.StatusUnprocessableEntity,
			"message": "carrier cid is required",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNotCalled(t, "CreateCarrier", mock.Anything)
	})

	t.Run("create warehouse fail - incorrect body", func(t *testing.T) {
		serviceMock := service_mock.NewCarryServiceMock()
		serviceMock.On("CreateCarrier", mock.Anything).Return(nil, error(nil))
		handlerImp := NewCarryHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/carrier", handlerImp.CreateCarrier(&ctx))

		res := httptest.NewRecorder()
		reqBody := `{
						"cid": 086,
						"company_name": "Livetube",
						"address": "Room 192",
						"telephone": "374-776-3015",
						"locality_id": 1
					}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/carrier",
			bytes.NewBuffer([]byte(reqBody)),
		)
		router.ServeHTTP(res, req)

		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := map[string]any{
			"code":    http.StatusBadRequest,
			"message": "Bad request",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertNotCalled(t, "CreateCarrier", mock.Anything)
	})

	t.Run("create warehouse fail - conflict", func(t *testing.T) {
		serviceMock := service_mock.NewCarryServiceMock()
		serviceMock.On("CreateCarrier", mock.Anything).
			Return(models.Carrier{}, serviceErr.ErrCarrierCidDuplicate)
		handlerImp := NewCarryHandler(serviceMock)
		router := chi.NewRouter()
		router.Post("/api/v1/carrier", handlerImp.CreateCarrier(&ctx))

		res := httptest.NewRecorder()
		reqBody := `{
						"cid": "#086",
						"company_name": "Livetube",
						"address": "Room 192",
						"telephone": "374-776-3015",
						"locality_id": 1
					}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/carrier",
			bytes.NewBuffer([]byte(reqBody)),
		)
		router.ServeHTTP(res, req)

		expectedCode := http.StatusConflict
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := map[string]any{
			"code":    http.StatusConflict,
			"message": "carrier code already exists",
			"data":    nil,
		}
		expectedStringBody, _ := json.Marshal(expectedBody)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedStringBody), res.Body.String())
		serviceMock.AssertCalled(t, "CreateCarrier", mock.Anything)
	})
}
