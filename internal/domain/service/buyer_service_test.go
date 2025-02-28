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

func ptrStr(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}

func ptrIntNil() *int {
	return nil
}

func assertBuyerResponse(
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

func TestBuyerService(t *testing.T) {
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
			name:             "[GetAll] OK Get all buyers",
			repositoryMethod: "GetAll",
			mockResponse: []entities.BuyerEntity{
				{
					ID:           1,
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
				{
					ID:           2,
					CardNumberID: ptrStr("F1098765432"),
					FirstName:    ptrStr("Jane"),
					LastName:     ptrStr("Doe"),
				},
			},
			expectedOutput: []models.Buyer{
				{
					ID: 1,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: ptrStr("M1234567890"),
						FirstName:    ptrStr("John"),
						LastName:     ptrStr("Doe"),
					},
				},
				{
					ID: 2,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: ptrStr("F1098765432"),
						FirstName:    ptrStr("Jane"),
						LastName:     ptrStr("Doe"),
					},
				},
			},
		},
		{
			name:             "[GetAll] Error No buyers registered yet",
			repositoryMethod: "GetAll",
			mockResponse:     make([]entities.BuyerEntity, 0),
			mockError:        repository.ErrBuyerNoRegisteredBuyersYet,
			expectedOutput:   make([]models.Buyer, 0),
			expectedErr:      service.ErrBuyerNoRegisteredBuyersYet,
		},
		{
			name:             "[GetAll] Error Unexpected error",
			repositoryMethod: "GetAll",
			mockResponse:     make([]entities.BuyerEntity, 0),
			mockError:        repository.ErrBuyerUnexpectedError,
			expectedOutput:   make([]models.Buyer, 0),
			expectedErr:      service.ErrBuyerUnexpectedError,
		},
		{
			name:             "[GetById] OK Get buyer by ID",
			repositoryMethod: "GetById",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse: entities.BuyerEntity{
				ID:           1,
				CardNumberID: ptrStr("M1234567890"),
				FirstName:    ptrStr("John"),
				LastName:     ptrStr("Doe"),
			},
			expectedOutput: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
		},
		{
			name:             "[GetById] Error Buyer not found",
			repositoryMethod: "GetById",
			serviceParams:    []interface{}{99},
			mockParams:       []interface{}{99},
			mockResponse:     entities.BuyerEntity{},
			mockError:        repository.ErrBuyerNotFoundWithID,
			expectedOutput:   models.Buyer{},
			expectedErr:      service.ErrBuyerDoesNotExist,
		},
		{
			name:             "[GetById] Error Unexpected error",
			repositoryMethod: "GetById",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse:     entities.BuyerEntity{},
			mockError:        repository.ErrBuyerUnexpectedError,
			expectedOutput:   models.Buyer{},
			expectedErr:      service.ErrBuyerUnexpectedError,
		},
		{
			name:             "[Create] OK Create new buyer",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234543210"),
					FirstName:    ptrStr("Baby"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockParams: []interface{}{
				entities.BuyerEntity{
					CardNumberID: ptrStr("M1234543210"),
					FirstName:    ptrStr("Baby"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockResponse: entities.BuyerEntity{
				ID:           3,
				CardNumberID: ptrStr("M1234543210"),
				FirstName:    ptrStr("Baby"),
				LastName:     ptrStr("Doe"),
			},
			expectedOutput: models.Buyer{
				ID: 3,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: ptrStr("M1234543210"),
					FirstName:    ptrStr("Baby"),
					LastName:     ptrStr("Doe"),
				},
			},
		},
		{
			name:             "[Create] Error Buyer already exists",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockParams: []interface{}{
				entities.BuyerEntity{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockResponse:   entities.BuyerEntity{},
			mockError:      repository.ErrBuyerDuplicateCardNumberID,
			expectedOutput: models.Buyer{},
			expectedErr:    service.ErrBuyerAlreadyExists,
		},
		{
			name:             "[Create] Error Unexpected error",
			repositoryMethod: "Create",
			serviceParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockParams: []interface{}{
				entities.BuyerEntity{
					CardNumberID: ptrStr("M1234567890"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
			mockResponse:   entities.BuyerEntity{},
			mockError:      service.ErrBuyerUnexpectedError,
			expectedOutput: models.Buyer{},
			expectedErr:    service.ErrBuyerUnexpectedError,
		},
		{
			name:             "[Patch] OK Update buyer",
			repositoryMethod: "Patch",
			serviceParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockParams: []interface{}{
				1,
				entities.BuyerEntity{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockResponse: entities.BuyerEntity{
				ID:           1,
				CardNumberID: ptrStr("M1234567876"),
				FirstName:    ptrStr("John"),
				LastName:     ptrStr("Doe"),
			},
			expectedOutput: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
					FirstName:    ptrStr("John"),
					LastName:     ptrStr("Doe"),
				},
			},
		},
		{
			name:             "[Patch] Error Buyer not found",
			repositoryMethod: "Patch",
			serviceParams: []interface{}{
				99,
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockParams: []interface{}{
				99,
				entities.BuyerEntity{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockResponse:   entities.BuyerEntity{},
			mockError:      repository.ErrBuyerNotFoundWithID,
			expectedOutput: models.Buyer{},
			expectedErr:    service.ErrBuyerDoesNotExist,
		},
		{
			name:             "[Patch] Error Buyer already exists",
			repositoryMethod: "Patch",
			serviceParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: ptrStr("F1098765432"),
				},
			},
			mockParams: []interface{}{
				1,
				entities.BuyerEntity{
					CardNumberID: ptrStr("F1098765432"),
				},
			},
			mockResponse:   entities.BuyerEntity{},
			mockError:      repository.ErrBuyerDuplicateCardNumberID,
			expectedOutput: models.Buyer{},
			expectedErr:    service.ErrBuyerAlreadyExists,
		},
		{
			name:             "[Patch] Error Unexpected error",
			repositoryMethod: "Patch",
			serviceParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockParams: []interface{}{
				1,
				entities.BuyerEntity{
					CardNumberID: ptrStr("M1234567876"),
				},
			},
			mockResponse:   entities.BuyerEntity{},
			mockError:      repository.ErrBuyerUnexpectedError,
			expectedOutput: models.Buyer{},
			expectedErr:    service.ErrBuyerUnexpectedError,
		},
		{
			name:             "[Delete] OK Delete buyer",
			repositoryMethod: "Delete",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse:     nil,
			mockError:        nil,
			expectedOutput:   nil,
		},
		{
			name:             "[Delete] Error Buyer not found",
			repositoryMethod: "Delete",
			serviceParams:    []interface{}{99},
			mockParams:       []interface{}{99},
			mockResponse:     nil,
			mockError:        repository.ErrBuyerNotFoundWithID,
			expectedOutput:   nil,
			expectedErr:      service.ErrBuyerDoesNotExist,
		},
		{
			name:             "[Delete] Error Cannot delete buyer with orders",
			repositoryMethod: "Delete",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse:     nil,
			mockError:        repository.ErrBuyerCannotDeleteBuyerWithOrders,
			expectedOutput:   nil,
			expectedErr:      service.ErrBuyerCannotDeleteBuyerWithOrders,
		},
		{
			name:             "[Delete] Error Unexpected error",
			repositoryMethod: "Delete",
			serviceParams:    []interface{}{1},
			mockParams:       []interface{}{1},
			mockResponse:     nil,
			mockError:        repository.ErrBuyerUnexpectedError,
			expectedOutput:   nil,
			expectedErr:      service.ErrBuyerUnexpectedError,
		},
		{
			name:             "[GetReportPurchaseOrders] Ok Get purchase order report",
			repositoryMethod: "GetReportPurchaseOrders",
			serviceParams:    []interface{}{ptrIntNil()},
			mockParams:       []interface{}{ptrIntNil()},
			mockResponse: []entities.ReportPurchaseOrdersEntity{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
				{
					ID:                  2,
					CardNumberID:        "F1098765432",
					FirstName:           "Jane",
					LastName:            "Doe",
					PurchaseOrdersCount: 27,
				},
			},
			expectedOutput: []models.ReportPurchaseOrders{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
				{
					ID:                  2,
					CardNumberID:        "F1098765432",
					FirstName:           "Jane",
					LastName:            "Doe",
					PurchaseOrdersCount: 27,
				},
			},
		},
		{
			name:             "[GetReportPurchaseOrders] Ok Get purchase order report by ID",
			repositoryMethod: "GetReportPurchaseOrders",
			serviceParams:    []interface{}{ptrInt(1)},
			mockParams:       []interface{}{ptrInt(1)},
			mockResponse: []entities.ReportPurchaseOrdersEntity{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
			},
			expectedOutput: []models.ReportPurchaseOrders{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)

			if tt.mockResponse != nil {
				mockRepository.On(tt.repositoryMethod, tt.mockParams...).
					Return(tt.mockResponse, tt.mockError)
			} else {
				mockRepository.On(tt.repositoryMethod, tt.mockParams...).Return(tt.mockError)
			}

			buyerService := serviceImpl.NewBuyerService(mockRepository)

			switch tt.repositoryMethod {
			case "GetAll":
				result, err := buyerService.GetAll()
				assertBuyerResponse(
					t,
					err,
					tt.expectedErr,
					result,
					tt.expectedOutput,
				)
			case "GetById":
				result, err := buyerService.GetById(tt.mockParams[0].(int))
				assertBuyerResponse(
					t,
					err,
					tt.expectedErr,
					result,
					tt.expectedOutput,
				)
			case "Create":
				param, ok := tt.serviceParams[0].(models.BuyerAttributes)
				if !ok {
					t.Fatalf(
						"Invalid parameter type for Create method. Expected models.BuyerAttributes, got %T",
						tt.serviceParams[0],
					)
				}

				result, err := buyerService.Create(param)
				assertBuyerResponse(
					t,
					err,
					tt.expectedErr,
					result,
					tt.expectedOutput,
				)
			case "Patch":
				param, ok := tt.serviceParams[1].(models.BuyerAttributes)
				if !ok {
					t.Fatalf(
						"Invalid parameter type for Create method. Expected models.BuyerAttributes, got %T",
						tt.serviceParams[0],
					)
				}

				result, err := buyerService.Patch(tt.mockParams[0].(int), param)
				assertBuyerResponse(
					t,
					err,
					tt.expectedErr,
					result,
					tt.expectedOutput,
				)
			case "Delete":
				err := buyerService.Delete(tt.mockParams[0].(int))
				assertBuyerResponse(
					t,
					err,
					tt.expectedErr,
					nil,
					tt.expectedOutput,
				)
			case "GetReportPurchaseOrders":
				result, err := buyerService.GetReportPurchaseOrders(tt.mockParams[0].(*int))
				assertBuyerResponse(
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
