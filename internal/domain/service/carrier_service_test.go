package service

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repositoryErr "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	serviceErr "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	fakeModels "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/models"
	repository_mock "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetAllCarriers(t *testing.T) {
	t.Run("find all carriers", func(t *testing.T) {
		repositoryMock := repository_mock.NewCarrierRepositoryMock()
		repositoryMock.On("GetAllCarriers").Return(fakeModels.Carriers, error(nil))
		serviceImp := NewCarryService(repositoryMock)

		carriers, err := serviceImp.GetAllCarriers()

		require.NoError(t, err)
		require.Equal(t, fakeModels.Carriers, carriers)
		require.Len(t, carriers, 2)
		repositoryMock.AssertNumberOfCalls(t, "GetAllCarriers", 1)
	})

	t.Run("find all carriers fail", func(t *testing.T) {
		repositoryMock := repository_mock.NewCarrierRepositoryMock()
		repositoryMock.On("GetAllCarriers").Return(nil, repositoryErr.ErrCarrierDataBase)
		serviceImp := NewCarryService(repositoryMock)

		carriers, err := serviceImp.GetAllCarriers()

		require.Nil(t, carriers)
		require.ErrorIs(t, err, serviceErr.ErrCarrierServiceDefault)
		repositoryMock.AssertNumberOfCalls(t, "GetAllCarriers", 1)
	})
}

func TestCreateCarrier(t *testing.T) {
	t.Run("create carrier", func(t *testing.T) {
		repositoryMock := repository_mock.NewCarrierRepositoryMock()
		carrierFake := fakeModels.Carriers[0]
		repositoryMock.On("CreateCarrier", mock.Anything).Return(carrierFake, error(nil))
		serviceImp := NewCarryService(repositoryMock)

		carrier, err := serviceImp.CreateCarrier(carrierFake)

		require.NoError(t, err)
		require.Equal(t, carrierFake, carrier)
		repositoryMock.AssertNumberOfCalls(t, "CreateCarrier", 1)
	})

	t.Run("create carriers fail", func(t *testing.T) {
		repositoryMock := repository_mock.NewCarrierRepositoryMock()
		repositoryMock.On("CreateCarrier", mock.Anything).Return(models.Carrier{}, repositoryErr.ErrCarrierCidDuplicate)
		serviceImp := NewCarryService(repositoryMock)

		carrier, err := serviceImp.CreateCarrier(fakeModels.Carriers[0])

		require.IsType(t, fakeModels.Carriers[0], carrier)
		require.ErrorIs(t, err, serviceErr.ErrCarrierCidDuplicate)
		repositoryMock.AssertNumberOfCalls(t, "CreateCarrier", 1)
	})
}
