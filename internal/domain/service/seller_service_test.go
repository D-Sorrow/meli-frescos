package service_test

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repository_errors "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func TestSellerService_CreateSeller(t *testing.T) {
	t.Run("seller created succesfully", func(t *testing.T) {
		sellerFake := models.Seller{
			Id:          1,
			Cid:         1,
			CompanyName: "fake_company",
			Address:     "Cll fake # 3",
			Telephone:   "312433213",
			LocalityId:  1,
		}

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("CreateSeller", sellerFake).Return(
			sellerFake, nil,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.CreateSeller(sellerFake)

		require.NoError(t, err)
		require.Equal(t, sellerFake, response)
		mockRepository.AssertCalled(t, "CreateSeller", sellerFake)
		mockRepository.AssertExpectations(t)
	})
	t.Run("seller create conflict", func(t *testing.T) {
		sellerFake := models.Seller{
			Id:          1,
			Cid:         1,
			CompanyName: "fake_company",
			Address:     "Cll fake # 3",
			Telephone:   "312433213",
			LocalityId:  1,
		}

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("CreateSeller", sellerFake).Return(
			models.Seller{}, repository_errors.ErrSellerAlreadyExists,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.CreateSeller(sellerFake)

		require.ErrorIs(t, err, service_errors.ErrSellerAlreadyExists)
		require.IsType(t, models.Seller{}, response)
		mockRepository.AssertCalled(t, "CreateSeller", sellerFake)
		mockRepository.AssertExpectations(t)
	})
}

func TestSellerService_ReadSellers(t *testing.T) {
	t.Run("GetAll Sellers succesfully", func(t *testing.T) {
		sellersFake := make(map[int]models.Seller)
		sellersFake[1] = models.Seller{
			Id:          1,
			Cid:         1,
			CompanyName: "fake_company",
			Address:     "Cll fake # 3",
			Telephone:   "312433213",
			LocalityId:  1,
		}
		sellersFake[2] = models.Seller{
			Id:          2,
			Cid:         2,
			CompanyName: "fake_company 2",
			Address:     "Cll fake # 123",
			Telephone:   "3124456456",
			LocalityId:  2,
		}

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("GetSellers").Return(
			sellersFake, nil,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.GetSellers()

		require.NoError(t, err)
		require.Equal(t, sellersFake, response)
		require.Equal(t, 2, len(response))
		mockRepository.AssertCalled(t, "GetSellers")
		mockRepository.AssertExpectations(t)
	})

	t.Run("GetAll Sellers - not found sellers", func(t *testing.T) {

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("GetSellers").Return(
			map[int]models.Seller{}, repository_errors.ErrSellerNotFound,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.GetSellers()

		require.ErrorIs(t, err, service_errors.ErrSellerNotFound)
		require.IsType(t, map[int]models.Seller{}, response)
		mockRepository.AssertCalled(t, "GetSellers")
		mockRepository.AssertExpectations(t)
	})

	t.Run("FindById Sellers succesfully", func(t *testing.T) {
		sellerToFind := 1

		sellerFake := models.Seller{
			Id:          1,
			Cid:         1,
			CompanyName: "fake_company",
			Address:     "Cll fake # 3",
			Telephone:   "312433213",
			LocalityId:  1,
		}

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("GetSellerById", sellerToFind).Return(
			sellerFake, nil,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.GetSellerById(sellerToFind)

		require.NoError(t, err)
		require.Equal(t, sellerFake, response)
		require.Equal(t, sellerToFind, response.Id)
		mockRepository.AssertCalled(t, "GetSellerById", sellerToFind)
		mockRepository.AssertExpectations(t)
	})

	t.Run("FindById Sellers not exist", func(t *testing.T) {
		sellerToFind := 10

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("GetSellerById", sellerToFind).Return(
			models.Seller{}, repository_errors.ErrSellerNotFound,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.GetSellerById(sellerToFind)

		require.ErrorIs(t, err, service_errors.ErrSellerNotFound)
		require.IsType(t, models.Seller{}, response)
		mockRepository.AssertCalled(t, "GetSellerById", sellerToFind)
		mockRepository.AssertExpectations(t)
	})
}

func TestSellerService_Update(t *testing.T) {
	t.Run("Update seller successfully", func(t *testing.T) {
		id := 1
		cId := 10
		companyName := "fake_company_updated"
		address := "Cll fake # 3"
		telephone := "312433213"
		localityId := 1

		sellerToUpdateFake := models.SellerPatch{
			Cid:         &cId,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityId:  &localityId,
		}

		sellerExpected := models.Seller{
			Id:          id,
			Cid:         cId,
			CompanyName: companyName,
			Address:     address,
			Telephone:   telephone,
			LocalityId:  localityId,
		}

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("UpdateSeller", id, sellerToUpdateFake).Return(
			sellerExpected, nil,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.UpdateSeller(id, sellerToUpdateFake)

		require.NoError(t, err)
		require.Equal(t, sellerExpected, response)
		mockRepository.AssertCalled(t, "UpdateSeller", id, sellerToUpdateFake)
		mockRepository.AssertExpectations(t)
	})

	t.Run("Update seller fail - seller not exist", func(t *testing.T) {
		id := 1
		cId := 10
		companyName := "fake_company_updated"
		address := "Cll fake # 3"
		telephone := "312433213"
		localityId := 1

		sellerToUpdateFake := models.SellerPatch{
			Cid:         &cId,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityId:  &localityId,
		}

		sellerExpected := models.Seller{}

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("UpdateSeller", id, sellerToUpdateFake).Return(
			sellerExpected, repository_errors.ErrSellerNotFound,
		)

		serviceTest := service.NewSellerService(mockRepository)

		response, err := serviceTest.UpdateSeller(id, sellerToUpdateFake)

		require.ErrorIs(t, err, service_errors.ErrSellerNotFound)
		require.Equal(t, sellerExpected, response)
		mockRepository.AssertCalled(t, "UpdateSeller", id, sellerToUpdateFake)
		mockRepository.AssertExpectations(t)
	})
}

func TestSellerService_Delete(t *testing.T) {
	t.Run("Delete seller successfully", func(t *testing.T) {
		id := 1

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("DeleteSeller", id).Return(
			nil,
		)

		serviceTest := service.NewSellerService(mockRepository)

		err := serviceTest.DeleteSeller(id)

		require.NoError(t, err)
		mockRepository.AssertCalled(t, "DeleteSeller", id)
		mockRepository.AssertExpectations(t)
	})

	t.Run("Delete seller fail - seller not exist", func(t *testing.T) {
		id := 1

		mockRepository := new(repository_mock.SellerRepositoryMock)
		mockRepository.On("DeleteSeller", id).Return(
			repository_errors.ErrSellerNotFound,
		)

		serviceTest := service.NewSellerService(mockRepository)

		err := serviceTest.DeleteSeller(id)

		require.ErrorIs(t, err, service_errors.ErrSellerNotFound)
		mockRepository.AssertCalled(t, "DeleteSeller", id)
		mockRepository.AssertExpectations(t)
	})
}
