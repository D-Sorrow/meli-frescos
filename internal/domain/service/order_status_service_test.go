package service_test

import (
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	serviceImpl "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/helpers"
	repository_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/infrastructure/repository"
)

func switchOrderStatusTest(
	t *testing.T,
	test helpers.ServiceTestStruct,
	orderStatusService *serviceImpl.OrderStatusService,
) {
	t.Helper()

	switch test.RepositoryMethod {
	case "GetAll":
		result, err := orderStatusService.GetAll()
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	}
}

func assertOrderStatusService(
	t *testing.T,
	test helpers.ServiceTestStruct,
	mockRepository *repository_mock.MockOrderStatusRepository,
) {
	t.Helper()

	orderStatusService := serviceImpl.NewOrderStatusService(mockRepository)
	switchOrderStatusTest(t, test, orderStatusService)
}

func TestOrderStatusGetAllService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[GetAll] OK Get all order statuses",
			RepositoryMethod: "GetAll",
			MockResponse: []entities.OrderStatusEntity{
				{
					ID:          1,
					Description: "Pendiente por pago",
				},
				{
					ID:          2,
					Description: "Pagado",
				},
			},
			ExpectedOutput: []models.OrderStatus{
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
		},
		{
			Name:             "[GetAll] Error No order statuses registered yet",
			RepositoryMethod: "GetAll",
			MockResponse:     make([]entities.OrderStatusEntity, 0),
			MockError:        repository.ErrOrderStatusNoRegisteredOrderStatusesYet,
			ExpectedOutput:   make([]models.OrderStatus, 0),
			ExpectedErr:      service.ErrOrderStatusNoRegisteredOrderStatusesYet,
		},
		{
			Name:             "[GetAll] Error Unexpected error",
			RepositoryMethod: "GetAll",
			MockResponse:     make([]entities.OrderStatusEntity, 0),
			MockError:        repository.ErrOrderStatusUnexpectedError,
			ExpectedOutput:   make([]models.OrderStatus, 0),
			ExpectedErr:      service.ErrOrderStatusUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockOrderStatusRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertOrderStatusService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}
