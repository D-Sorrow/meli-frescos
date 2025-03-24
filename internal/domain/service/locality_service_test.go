package service

import (
	"testing"

	LocalityModels "github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/models"
	repository_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

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
		repositoryMock.On("GetCarriersByAllLocalities").Return(nil, error(repository.ErrGetAllLocalities))
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
		repositoryMock.On("GetCarriersByLocality", mock.Anything).Return(models.LocalityCarriers[0], error(nil))
		serviceImp := NewLocalityService(repositoryMock)

		carrier, err := serviceImp.GetCarriersByLocality(1)

		require.NoError(t, err)
		require.Equal(t, models.LocalityCarriers[0], carrier)
		require.Equal(t, models.LocalityCarriers[0].LocalityId, carrier.LocalityId)
		repositoryMock.AssertCalled(t, "GetCarriersByLocality", mock.Anything)
	})

	t.Run("get carrier by locality id fail", func(t *testing.T) {
		repositoryMock := repository_mock.NewLocalityRepositoryMock()
		repositoryMock.On("GetCarriersByLocality", mock.Anything).Return(LocalityModels.LocalityCarriers{}, repository.ErrLocalityNotFound)
		serviceImp := NewLocalityService(repositoryMock)

		carrier, err := serviceImp.GetCarriersByLocality(1)

		require.ErrorIs(t, service.ErrLocalityNotFound, err)
		require.IsType(t, LocalityModels.LocalityCarriers{}, carrier)
		repositoryMock.AssertCalled(t, "GetCarriersByLocality", mock.Anything)
	})
}
