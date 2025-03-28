package service

import (
	"testing"

	LocalityModels "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	models_app "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/models"
	repository_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateLocality(t *testing.T) {
	t.Run("create a new locality", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("CreateLocality", models.LocalityRequest).
			Return(models.LocalityRequest, error(nil))
		serviceImp := NewLocalityService(repositoryMock)

		res, err := serviceImp.CreateLocality(models.LocalityRequest)

		require.NoError(t, err)
		require.Equal(t, models.LocalityRequest, res)
		repositoryMock.AssertCalled(t, "CreateLocality", models.LocalityRequest)
		repositoryMock.AssertNumberOfCalls(t, "CreateLocality", 1)

	})

	t.Run("create a new locality fail - locality already exist", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("CreateLocality", models.LocalityRequest).
			Return(models_app.Locality{}, repository.ErrLocalityAlreadyExists)
		serviceImp := NewLocalityService(repositoryMock)

		res, err := serviceImp.CreateLocality(models.LocalityRequest)

		require.Equal(t, res, models_app.Locality{})
		require.ErrorIs(t, err, service.ErrLocalityAlreadyExists)
		repositoryMock.AssertCalled(t, "CreateLocality", models.LocalityRequest)
		repositoryMock.AssertNumberOfCalls(t, "CreateLocality", 1)

	})

	t.Run("create a new locality fail - province not found ", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("CreateLocality", models.LocalityRequest).
			Return(models_app.Locality{}, repository.ErrProvinceNotFound)
		serviceImp := NewLocalityService(repositoryMock)

		res, err := serviceImp.CreateLocality(models.LocalityRequest)

		require.Equal(t, res, models_app.Locality{})
		require.ErrorIs(t, err, service.ErrProvinceNotFound)
		repositoryMock.AssertCalled(t, "CreateLocality", models.LocalityRequest)
		repositoryMock.AssertNumberOfCalls(t, "CreateLocality", 1)

	})
}

func TestGetSellersByLocality(t *testing.T) {
	t.Run("get all carriers by localities", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetSellersByLocality", 1).
			Return(models.LocalitySellersResponse, error(nil))
		serviceImp := NewLocalityService(repositoryMock)

		res, err := serviceImp.GetSellersByLocality(1)

		require.NoError(t, err)
		require.Equal(t, models.LocalitySellersResponse, res)
		repositoryMock.AssertCalled(t, "GetSellersByLocality", 1)
		repositoryMock.AssertNumberOfCalls(t, "GetSellersByLocality", 1)

	})

	t.Run("get all carriers by localities fail - locality not found", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetSellersByLocality", 1).
			Return(models_app.LocalitySellers{}, repository.ErrLocalityNotFound)
		serviceImp := NewLocalityService(repositoryMock)

		res, err := serviceImp.GetSellersByLocality(1)

		require.Equal(t, res, models_app.LocalitySellers{})
		require.ErrorIs(t, err, service.ErrLocalityNotFound)
		repositoryMock.AssertCalled(t, "GetSellersByLocality", 1)
		repositoryMock.AssertNumberOfCalls(t, "GetSellersByLocality", 1)

	})
}

func TestGetCarriersByAllLocalities(t *testing.T) {
	t.Run("get all carriers by localities", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetCarriersByAllLocalities").Return(models.LocalityCarriers, error(nil))
		serviceImp := NewLocalityService(repositoryMock)

		carriersByLocalities, err := serviceImp.GetCarriersByAllLocalities()

		require.NoError(t, err)
		require.Len(t, carriersByLocalities, 2)
		require.Equal(t, models.LocalityCarriers, carriersByLocalities)
		repositoryMock.AssertCalled(t, "GetCarriersByAllLocalities")

	})

	t.Run("get all carriers by localities fail", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetCarriersByAllLocalities").
			Return(nil, error(repository.ErrGetAllLocalities))
		serviceImp := NewLocalityService(repositoryMock)

		carriersByLocalities, err := serviceImp.GetCarriersByAllLocalities()

		require.Nil(t, carriersByLocalities)
		require.ErrorIs(t, service.ErrGetAllLocalities, err)
		repositoryMock.AssertCalled(t, "GetCarriersByAllLocalities")

	})
}

func TestGetCarriersByLocality(t *testing.T) {
	t.Run("get carrier by locality id", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetCarriersByLocality", mock.Anything).
			Return(models.LocalityCarriers[0], error(nil))
		serviceImp := NewLocalityService(repositoryMock)

		carrier, err := serviceImp.GetCarriersByLocality(1)

		require.NoError(t, err)
		require.Equal(t, models.LocalityCarriers[0], carrier)
		require.Equal(t, models.LocalityCarriers[0].LocalityId, carrier.LocalityId)
		repositoryMock.AssertCalled(t, "GetCarriersByLocality", mock.Anything)
	})

	t.Run("get carrier by locality id fail", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetCarriersByLocality", mock.Anything).
			Return(LocalityModels.LocalityCarriers{}, repository.ErrLocalityNotFound)
		serviceImp := NewLocalityService(repositoryMock)

		carrier, err := serviceImp.GetCarriersByLocality(1)

		require.ErrorIs(t, service.ErrLocalityNotFound, err)
		require.IsType(t, LocalityModels.LocalityCarriers{}, carrier)
		repositoryMock.AssertCalled(t, "GetCarriersByLocality", mock.Anything)
	})
}
