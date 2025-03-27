package handlers_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/melisource/fury_go-core/pkg/web"
	"github.com/melisource/fury_go-platform/pkg/fury"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/mocks/helpers"
	service_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
)

func switchOrderStatusTest(
	t *testing.T,
	test helpers.HandlerTestStruct,
	rt *web.Router,
	orderStatusHandler *handlers.OrderStatusHandler,
) {
	t.Helper()
	ctx := context.Background()

	switch test.ServiceMethod {
	case "GetAll":
		rt.Get("/api/v1/orderStatus", orderStatusHandler.GetAll(&ctx))
		result := helpers.GetResult(t, rt, test.HttpPath, test.HttpMethod, nil)
		helpers.CheckHandlerResponse[[]dto.OrderStatusDTO](t, test, result)
	}
}

func assertOrderStatusHandler(
	t *testing.T,
	test helpers.HandlerTestStruct,
	mockService *service_mock.MockOrderStatusService,
) {
	t.Helper()

	app, err := fury.NewWebApplication()
	if err != nil {
		t.Fatal(err)
	}

	rt := app.Router
	orderStatusHandler := handlers.NewOrderStatusHandler(mockService)
	switchOrderStatusTest(t, test, rt, orderStatusHandler)
}

func TestOrderStatusGetAllHandler(t *testing.T) {
	tests := []helpers.HandlerTestStruct{
		{
			Name:          "[GetAll] OK Get all order statuses",
			ServiceMethod: "GetAll",
			HttpPath:      "/api/v1/orderStatus",
			HttpMethod:    "GET",
			MockResponse: []models.OrderStatus{
				{
					ID: 1,
					OrderStatusAttributes: models.OrderStatusAttributes{
						Description: "Pendiente por pago",
					},
				},
				{
					ID: 2,
					OrderStatusAttributes: models.OrderStatusAttributes{
						Description: "Pagado",
					},
				},
			},
			ExpectedOutput: helpers.TestGenericStruct[[]dto.OrderStatusDTO]{
				Code: http.StatusOK,
				Msg:  "Get all order statuses successful",
				Data: []dto.OrderStatusDTO{
					{
						ID:          1,
						Description: "Pendiente por pago",
					},
					{
						ID:          2,
						Description: "Pagado",
					},
				},
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetAll] Error No order statuses registered yet",
			ServiceMethod: "GetAll",
			HttpPath:      "/api/v1/orderStatus",
			HttpMethod:    "GET",
			MockResponse:  make([]models.OrderStatus, 0),
			MockError:     service.ErrOrderStatusNoRegisteredOrderStatusesYet,
			ExpectedOutput: helpers.TestGenericStruct[[]dto.OrderStatusDTO]{
				Code: http.StatusOK,
				Msg:  "ERR: No registered order statuses yet",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedCalls:      1,
		},
		{
			Name:          "[GetAll] Error Unexpected error",
			ServiceMethod: "GetAll",
			HttpPath:      "/api/v1/orderStatus",
			HttpMethod:    "GET",
			MockResponse:  make([]models.OrderStatus, 0),
			MockError:     service.ErrOrderStatusUnexpectedError,
			ExpectedOutput: helpers.TestGenericStruct[[]dto.OrderStatusDTO]{
				Code: http.StatusInternalServerError,
				Msg:  "ERR: An unexpected error occurred while processing the requested order status, please try again later",
				Data: nil,
			},
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedCalls:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockService := new(service_mock.MockOrderStatusService)
			helpers.InitServiceMock(t, test, &mockService)
			assertOrderStatusHandler(t, test, mockService)

			mockService.AssertExpectations(t)
			mockService.AssertNumberOfCalls(
				t,
				test.ServiceMethod,
				test.ExpectedCalls,
			)
		})
	}
}
