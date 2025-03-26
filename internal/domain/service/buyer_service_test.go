package service_test

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	serviceImpl "github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	"github.com/D-Sorrow/meli-frescos/mocks/helpers"
	repository_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
)

func switchBuyerTest(
	t *testing.T,
	test helpers.ServiceTestStruct,
	buyerService *serviceImpl.BuyerService,
) {
	t.Helper()

	switch test.RepositoryMethod {
	case "GetAll":
		result, err := buyerService.GetAll()
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	case "GetById":
		result, err := buyerService.GetById(test.MockParams[0].(int))
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	case "Create":
		param, ok := test.ServiceParams[0].(models.BuyerAttributes)
		if !ok {
			t.Fatalf(
				"Invalid parameter type for Create method. Expected models.BuyerAttributes, got %T",
				test.ServiceParams[0],
			)
		}

		result, err := buyerService.Create(param)
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	case "Patch":
		param, ok := test.ServiceParams[1].(models.BuyerAttributes)
		if !ok {
			t.Fatalf(
				"Invalid parameter type for Patch method. Expected models.BuyerAttributes, got %T",
				test.ServiceParams[0],
			)
		}

		result, err := buyerService.Patch(test.MockParams[0].(int), param)
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	case "Delete":
		err := buyerService.Delete(test.MockParams[0].(int))
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			nil,
			test.ExpectedOutput,
		)
	case "GetReportPurchaseOrders":
		result, err := buyerService.GetReportPurchaseOrders(test.MockParams[0].(*int))
		helpers.CheckServiceResponse(
			t,
			err,
			test.ExpectedErr,
			result,
			test.ExpectedOutput,
		)
	}
}

func assertBuyerService(
	t *testing.T,
	test helpers.ServiceTestStruct,
	mockRepository *repository_mock.MockBuyerRepository,
) {
	t.Helper()

	buyerService := serviceImpl.NewBuyerService(mockRepository)
	switchBuyerTest(t, test, buyerService)
}

func TestBuyerGetAllService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[GetAll] OK Get all buyers",
			RepositoryMethod: "GetAll",
			MockResponse: []entities.BuyerEntity{
				{
					ID:           1,
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
				{
					ID:           2,
					CardNumberID: helpers.PtrStr("F1098765432"),
					FirstName:    helpers.PtrStr("Jane"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			ExpectedOutput: []models.Buyer{
				{
					ID: 1,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: helpers.PtrStr("M1234567890"),
						FirstName:    helpers.PtrStr("John"),
						LastName:     helpers.PtrStr("Doe"),
					},
				},
				{
					ID: 2,
					BuyerAttributes: models.BuyerAttributes{
						CardNumberID: helpers.PtrStr("F1098765432"),
						FirstName:    helpers.PtrStr("Jane"),
						LastName:     helpers.PtrStr("Doe"),
					},
				},
			},
		},
		{
			Name:             "[GetAll] Error No buyers registered yet",
			RepositoryMethod: "GetAll",
			MockResponse:     make([]entities.BuyerEntity, 0),
			MockError:        repository.ErrBuyerNoRegisteredBuyersYet,
			ExpectedOutput:   make([]models.Buyer, 0),
			ExpectedErr:      service.ErrBuyerNoRegisteredBuyersYet,
		},
		{
			Name:             "[GetAll] Error Unexpected error",
			RepositoryMethod: "GetAll",
			MockResponse:     make([]entities.BuyerEntity, 0),
			MockError:        repository.ErrBuyerUnexpectedError,
			ExpectedOutput:   make([]models.Buyer, 0),
			ExpectedErr:      service.ErrBuyerUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertBuyerService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestBuyerGetByIdService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[GetById] OK Get buyer by ID",
			RepositoryMethod: "GetById",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse: entities.BuyerEntity{
				ID:           1,
				CardNumberID: helpers.PtrStr("M1234567890"),
				FirstName:    helpers.PtrStr("John"),
				LastName:     helpers.PtrStr("Doe"),
			},
			ExpectedOutput: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
		},
		{
			Name:             "[GetById] Error Buyer not found",
			RepositoryMethod: "GetById",
			ServiceParams:    []interface{}{99},
			MockParams:       []interface{}{99},
			MockResponse:     entities.BuyerEntity{},
			MockError:        repository.ErrBuyerNotFoundWithID,
			ExpectedOutput:   models.Buyer{},
			ExpectedErr:      service.ErrBuyerDoesNotExist,
		},
		{
			Name:             "[GetById] Error Unexpected error",
			RepositoryMethod: "GetById",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse:     entities.BuyerEntity{},
			MockError:        repository.ErrBuyerUnexpectedError,
			ExpectedOutput:   models.Buyer{},
			ExpectedErr:      service.ErrBuyerUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertBuyerService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestBuyerCreateService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[Create] OK Create new buyer",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234543210"),
					FirstName:    helpers.PtrStr("Baby"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockParams: []interface{}{
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("M1234543210"),
					FirstName:    helpers.PtrStr("Baby"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockResponse: entities.BuyerEntity{
				ID:           3,
				CardNumberID: helpers.PtrStr("M1234543210"),
				FirstName:    helpers.PtrStr("Baby"),
				LastName:     helpers.PtrStr("Doe"),
			},
			ExpectedOutput: models.Buyer{
				ID: 3,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234543210"),
					FirstName:    helpers.PtrStr("Baby"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
		},
		{
			Name:             "[Create] Error Buyer already exists",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockParams: []interface{}{
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockResponse:   entities.BuyerEntity{},
			MockError:      repository.ErrBuyerDuplicateCardNumberID,
			ExpectedOutput: models.Buyer{},
			ExpectedErr:    service.ErrBuyerAlreadyExists,
		},
		{
			Name:             "[Create] Error Unexpected error",
			RepositoryMethod: "Create",
			ServiceParams: []interface{}{
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockParams: []interface{}{
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("M1234567890"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
			MockResponse:   entities.BuyerEntity{},
			MockError:      service.ErrBuyerUnexpectedError,
			ExpectedOutput: models.Buyer{},
			ExpectedErr:    service.ErrBuyerUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertBuyerService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestBuyerPatchService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[Patch] OK Update buyer",
			RepositoryMethod: "Patch",
			ServiceParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockParams: []interface{}{
				1,
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockResponse: entities.BuyerEntity{
				ID:           1,
				CardNumberID: helpers.PtrStr("M1234567876"),
				FirstName:    helpers.PtrStr("John"),
				LastName:     helpers.PtrStr("Doe"),
			},
			ExpectedOutput: models.Buyer{
				ID: 1,
				BuyerAttributes: models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
					FirstName:    helpers.PtrStr("John"),
					LastName:     helpers.PtrStr("Doe"),
				},
			},
		},
		{
			Name:             "[Patch] Error Buyer not found",
			RepositoryMethod: "Patch",
			ServiceParams: []interface{}{
				99,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockParams: []interface{}{
				99,
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockResponse:   entities.BuyerEntity{},
			MockError:      repository.ErrBuyerNotFoundWithID,
			ExpectedOutput: models.Buyer{},
			ExpectedErr:    service.ErrBuyerDoesNotExist,
		},
		{
			Name:             "[Patch] Error Buyer already exists",
			RepositoryMethod: "Patch",
			ServiceParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("F1098765432"),
				},
			},
			MockParams: []interface{}{
				1,
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("F1098765432"),
				},
			},
			MockResponse:   entities.BuyerEntity{},
			MockError:      repository.ErrBuyerDuplicateCardNumberID,
			ExpectedOutput: models.Buyer{},
			ExpectedErr:    service.ErrBuyerAlreadyExists,
		},
		{
			Name:             "[Patch] Error Unexpected error",
			RepositoryMethod: "Patch",
			ServiceParams: []interface{}{
				1,
				models.BuyerAttributes{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockParams: []interface{}{
				1,
				entities.BuyerEntity{
					CardNumberID: helpers.PtrStr("M1234567876"),
				},
			},
			MockResponse:   entities.BuyerEntity{},
			MockError:      repository.ErrBuyerUnexpectedError,
			ExpectedOutput: models.Buyer{},
			ExpectedErr:    service.ErrBuyerUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertBuyerService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestBuyerDeleteService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[Delete] OK Delete buyer",
			RepositoryMethod: "Delete",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse:     nil,
			MockError:        nil,
			ExpectedOutput:   nil,
		},
		{
			Name:             "[Delete] Error Buyer not found",
			RepositoryMethod: "Delete",
			ServiceParams:    []interface{}{99},
			MockParams:       []interface{}{99},
			MockResponse:     nil,
			MockError:        repository.ErrBuyerNotFoundWithID,
			ExpectedOutput:   nil,
			ExpectedErr:      service.ErrBuyerDoesNotExist,
		},
		{
			Name:             "[Delete] Error Cannot delete buyer with orders",
			RepositoryMethod: "Delete",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse:     nil,
			MockError:        repository.ErrBuyerCannotDeleteBuyerWithOrders,
			ExpectedOutput:   nil,
			ExpectedErr:      service.ErrBuyerCannotDeleteBuyerWithOrders,
		},
		{
			Name:             "[Delete] Error Unexpected error",
			RepositoryMethod: "Delete",
			ServiceParams:    []interface{}{1},
			MockParams:       []interface{}{1},
			MockResponse:     nil,
			MockError:        repository.ErrBuyerUnexpectedError,
			ExpectedOutput:   nil,
			ExpectedErr:      service.ErrBuyerUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertBuyerService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestBuyerGetReportPurchaseOrdersService(t *testing.T) {
	tests := []helpers.ServiceTestStruct{
		{
			Name:             "[GetReportPurchaseOrders] Ok Get purchase order report",
			RepositoryMethod: "GetReportPurchaseOrders",
			ServiceParams:    []interface{}{helpers.PtrIntNil()},
			MockParams:       []interface{}{helpers.PtrIntNil()},
			MockResponse: []entities.ReportPurchaseOrdersEntity{
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
			ExpectedOutput: []models.ReportPurchaseOrders{
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
			Name:             "[GetReportPurchaseOrders] Ok Get purchase order report by ID",
			RepositoryMethod: "GetReportPurchaseOrders",
			ServiceParams:    []interface{}{helpers.PtrInt(1)},
			MockParams:       []interface{}{helpers.PtrInt(1)},
			MockResponse: []entities.ReportPurchaseOrdersEntity{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
			},
			ExpectedOutput: []models.ReportPurchaseOrders{
				{
					ID:                  1,
					CardNumberID:        "M1234567890",
					FirstName:           "John",
					LastName:            "Doe",
					PurchaseOrdersCount: 4,
				},
			},
		},
		{
			Name:             "[GetReportPurchaseOrders] Error Unexpected error",
			RepositoryMethod: "GetReportPurchaseOrders",
			ServiceParams:    []interface{}{helpers.PtrInt(1)},
			MockParams:       []interface{}{helpers.PtrInt(1)},
			MockResponse:     make([]entities.ReportPurchaseOrdersEntity, 0),
			MockError:        repository.ErrBuyerUnexpectedError,
			ExpectedOutput:   make([]models.ReportPurchaseOrders, 0),
			ExpectedErr:      service.ErrBuyerUnexpectedError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepository := new(repository_mock.MockBuyerRepository)
			helpers.InitRepositoryMock(t, test, &mockRepository)
			assertBuyerService(t, test, mockRepository)

			mockRepository.AssertExpectations(t)
		})
	}
}
