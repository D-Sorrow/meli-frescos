package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	serviceImpl "github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	repository_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
)

func assertOrderStatusResponse(
	t *testing.T,
	responseErr error,
	expectedErr error,
	responseOutput interface{},
	expectedOutput interface{},
) {
	t.Helper()

	assert.True(t, errors.Is(responseErr, expectedErr), "Error mismatch")
	assert.Equal(t, responseOutput, expectedOutput, "Output mismatch")
}

func TestOrderStatusService(t *testing.T) {
	tests := []struct {
		name             string
		repositoryMethod string
		serviceParams    []interface{}
		mockParams       []interface{}
		mockResponse     interface{}
		mockError        error
		expectedOutput   interface{}
		expectedErr      error
	}{
		{
			name:             "[GetAll] OK Get all order statuses",
			repositoryMethod: "GetAll",
			mockResponse: []entities.OrderStatusEntity{
				{
					ID:          1,
					Description: "Pendiente por pago",
				},
				{
					ID:          2,
					Description: "Pagado",
				},
			},
			expectedOutput: []models.OrderStatus{
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
			name:             "[GetAll] Error No order statuses registered yet",
			repositoryMethod: "GetAll",
			mockResponse:     make([]entities.OrderStatusEntity, 0),
			mockError:        repository.ErrOrderStatusNoRegisteredOrderStatusesYet,
			expectedOutput:   make([]models.OrderStatus, 0),
			expectedErr:      service.ErrOrderStatusNoRegisteredOrderStatusesYet,
		},
		{
			name:             "[GetAll] Error Unexpected error",
			repositoryMethod: "GetAll",
			mockResponse:     make([]entities.OrderStatusEntity, 0),
			mockError:        repository.ErrOrderStatusUnexpectedError,
			expectedOutput:   make([]models.OrderStatus, 0),
			expectedErr:      service.ErrOrderStatusUnexpectedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockOrderStatusRepository)

			if tt.mockResponse != nil {
				mockRepository.On(tt.repositoryMethod, tt.mockParams...).
					Return(tt.mockResponse, tt.mockError)
			} else {
				mockRepository.On(tt.repositoryMethod, tt.mockParams...).Return(tt.mockError)
			}

			orderStatusService := serviceImpl.NewOrderStatusService(mockRepository)

			switch tt.repositoryMethod {
			case "GetAll":
				result, err := orderStatusService.GetAll()
				assertOrderStatusResponse(
					t,
					err,
					tt.expectedErr,
					result,
					tt.expectedOutput,
				)
			}

			mockRepository.AssertExpectations(t)
		})
	}
}
