package service

import (
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	repositoryErr "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	serviceErr "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	fakeModels "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/domain/models"
	repository_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetWarehouses(t *testing.T) {
	t.Run("find all warehouses", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		mockRepository.On("GetWarehouses").Return(fakeModels.WarehousesFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.GetWarehouses()

		require.Equal(t, fakeModels.WarehousesFake, response)
		require.Equal(t, 2, len(response))
		require.NoError(t, err)
		mockRepository.AssertCalled(t, "GetWarehouses")
	})

	t.Run("find all warehouses fail", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		mockRepository.On("GetWarehouses").Return(nil, error(repositoryErr.ErrWarehouseDataBase))
		serviceImp := NewWarehouseService(mockRepository)

		_, err := serviceImp.GetWarehouses()

		require.ErrorIs(t, err, serviceErr.ErrWarehouseServiceDefault)
		mockRepository.AssertCalled(t, "GetWarehouses")
	})
}

func TestGetWarehouseById(t *testing.T) {
	t.Run("find warehouse successfull", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		warehouseFake := fakeModels.WarehousesFake[1]
		mockRepository.On("GetWarehouseById", 1).Return(warehouseFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.GetWarehouseById(1)

		require.NoError(t, err)
		require.Equal(t, warehouseFake, response)
		require.Equal(t, warehouseFake.WarehouseCode, response.WarehouseCode)
		mockRepository.AssertCalled(t, "GetWarehouseById", 1)
	})

	t.Run("warehouse not found", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		mockRepository.On("GetWarehouseById", 2).
			Return(models.Warehouse{}, repositoryErr.ErrWarehouseNotFound)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.GetWarehouseById(2)

		require.ErrorIs(t, err, serviceErr.ErrWarehouseNotFound)
		require.Equal(t, 0, response.Id)
		mockRepository.AssertCalled(t, "GetWarehouseById", 2)
	})
}

func TestCreateWarehouse(t *testing.T) {
	t.Run("create warehouse ok", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		warehouseFake := fakeModels.WarehousesFake[1]
		mockRepository.On("CreateWarehouse", warehouseFake).Return(warehouseFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.CreateWarehouse(warehouseFake)

		require.NoError(t, err)
		require.Equal(t, warehouseFake, response)
		mockRepository.AssertCalled(t, "CreateWarehouse", warehouseFake)
		mockRepository.AssertExpectations(t)
	})

	t.Run("create warehouse conflict", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		mockRepository.On("CreateWarehouse", models.Warehouse{}).
			Return(models.Warehouse{}, repositoryErr.ErrWarehouseCodeDuplicate)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.CreateWarehouse(models.Warehouse{})

		require.ErrorIs(t, err, serviceErr.ErrWarehouseCodeDuplicate)
		require.IsType(t, models.Warehouse{}, response)
		mockRepository.AssertCalled(t, "CreateWarehouse", models.Warehouse{})
	})
}

func TestPatchWarehouse(t *testing.T) {
	t.Run("update warehouse successful", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		warehouseFakeMap := fakeModels.WarehouseFakeMap
		warehouseFake := fakeModels.WarehousesFake[1]
		mockRepository.On("PatchWarehouse", 1, mock.Anything).Return(warehouseFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.PatchWarehouse(1, warehouseFakeMap)

		require.NoError(t, err)
		require.Equal(t, warehouseFake, response)
		mockRepository.AssertExpectations(t)
	})

	t.Run("update warehouse not found", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		warehouseFakeMap := fakeModels.WarehouseFakeMap
		mockRepository.On("PatchWarehouse", 1, mock.Anything).
			Return(models.Warehouse{}, repositoryErr.ErrWarehouseNotFound)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.PatchWarehouse(1, warehouseFakeMap)

		require.ErrorIs(t, err, serviceErr.ErrWarehouseNotFound)
		require.IsType(t, models.Warehouse{}, response)
		mockRepository.AssertExpectations(t)
	})

	t.Run("update warehouse conflict", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		warehouseFakeMap := fakeModels.WarehouseFakeMap
		mockRepository.On("PatchWarehouse", 1, mock.Anything).
			Return(models.Warehouse{}, repositoryErr.ErrWarehouseCodeDuplicate)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.PatchWarehouse(1, warehouseFakeMap)

		require.ErrorIs(t, err, serviceErr.ErrWarehouseCodeDuplicate)
		require.IsType(t, models.Warehouse{}, response)
		mockRepository.AssertExpectations(t)
	})
}

func TestDeleteWarehouse(t *testing.T) {
	t.Run("delete warehouse successful", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		mockRepository.On("DeleteWarehouse", 1).Return(error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		err := serviceImp.DeleteWarehouse(1)

		require.NoError(t, err)
		mockRepository.AssertCalled(t, "DeleteWarehouse", 1)
		mockRepository.AssertExpectations(t)
	})

	t.Run("delete warehouse not found", func(t *testing.T) {
		mockRepository := repository_mock.NewWarehouseRepositoryMock()
		mockRepository.On("DeleteWarehouse", 1).Return(repositoryErr.ErrWarehouseNotFound)
		serviceImp := NewWarehouseService(mockRepository)

		err := serviceImp.DeleteWarehouse(1)

		require.ErrorIs(t, err, serviceErr.ErrWarehouseNotFound)
		mockRepository.AssertCalled(t, "DeleteWarehouse", 1)
		mockRepository.AssertExpectations(t)
	})
}
