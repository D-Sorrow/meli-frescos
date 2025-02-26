package service

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repositoryErr "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	serviceInterface "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetWarehouses(t *testing.T) {
	t.Run("find all warehouses", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		warehousesFake := map[int]models.Warehouse{}
		warehousesFake[1] = models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		warehousesFake[2] = models.Warehouse{
			Id:                 2,
			WarehouseCode:      "b6b225e6-c83f-46a4-ac63-b6df8794ba59",
			Address:            "Room 192",
			Telephone:          "(917) 6928569",
			MinimunCapacity:    15,
			MinimunTemperature: 0,
			LocalityId:         2,
		}
		mockRepository.On("GetWarehouses").Return(warehousesFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.GetWarehouses()

		require.Equal(t, warehousesFake, response)
		require.Equal(t, 2, len(response))
		require.NoError(t, err)
		mockRepository.AssertCalled(t, "GetWarehouses")
	})

	t.Run("find all warehouses fail", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		mockRepository.On("GetWarehouses").Return(nil, error(repositoryErr.ErrWarehouseDataBase))
		serviceImp := NewWarehouseService(mockRepository)

		_, err := serviceImp.GetWarehouses()

		require.ErrorIs(t, err, serviceInterface.ErrWarehouseServiceDefault)
		mockRepository.AssertCalled(t, "GetWarehouses")
	})
}

func TestGetWarehouseById(t *testing.T) {
	t.Run("find warehouse successfull", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		warehouseFake := models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		mockRepository.On("GetWarehouseById", 1).Return(warehouseFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.GetWarehouseById(1)

		require.NoError(t, err)
		require.Equal(t, warehouseFake, response)
		require.Equal(t, warehouseFake.WarehouseCode, response.WarehouseCode)
		mockRepository.AssertCalled(t, "GetWarehouseById", 1)
	})

	t.Run("warehouse not found", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		mockRepository.On("GetWarehouseById", 2).Return(models.Warehouse{}, repositoryErr.ErrWarehouseNotFound)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.GetWarehouseById(2)

		require.ErrorIs(t, err, serviceInterface.ErrWarehouseNotFound)
		require.Equal(t, 0, response.Id)
		mockRepository.AssertCalled(t, "GetWarehouseById", 2)
	})
}

func TestCreateWarehouse(t *testing.T) {
	t.Run("create warehouse ok", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		warehouseFake := models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		mockRepository.On("CreateWarehouse", warehouseFake).Return(warehouseFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.CreateWarehouse(warehouseFake)

		require.NoError(t, err)
		require.Equal(t, warehouseFake, response)
		mockRepository.AssertCalled(t, "CreateWarehouse", warehouseFake)
		mockRepository.AssertExpectations(t)
	})

	t.Run("create warehouse conflict", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		mockRepository.On("CreateWarehouse", models.Warehouse{}).Return(models.Warehouse{}, repositoryErr.ErrWarehouseCodeDuplicate)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.CreateWarehouse(models.Warehouse{})

		require.ErrorIs(t, err, serviceInterface.ErrWarehouseCodeDuplicate)
		require.IsType(t, models.Warehouse{}, response)
		mockRepository.AssertCalled(t, "CreateWarehouse", models.Warehouse{})
	})
}

func TestPatchWarehouse(t *testing.T) {
	t.Run("update warehouse successful", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		warehouseFakeMap := map[string]interface{}{
			"Id":                 1,
			"WarehouseCode":      "96b5c517-cce3-4ed3-a06b-24234ee8d009",
			"Address":            "PO Box 16130",
			"Telephone":          "(617) 4591159",
			"MinimunCapacity":    92,
			"MinimunTemperature": 10,
			"LocalityId":         1,
		}
		warehouseFake := models.Warehouse{
			Id:                 1,
			WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
			Address:            "Apt 1639",
			Telephone:          "(639) 5350508",
			MinimunCapacity:    99,
			MinimunTemperature: -14,
			LocalityId:         1,
		}
		mockRepository.On("PatchWarehouse", 1, mock.Anything).Return(warehouseFake, error(nil))
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.PatchWarehouse(1, warehouseFakeMap)

		require.NoError(t, err)
		require.Equal(t, warehouseFake, response)
		mockRepository.AssertExpectations(t)
	})

	t.Run("update warehouse not found", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		warehouseFakeMap := map[string]interface{}{
			"Id":                 1,
			"WarehouseCode":      "96b5c517-cce3-4ed3-a06b-24234ee8d009",
			"Address":            "PO Box 16130",
			"Telephone":          "(617) 4591159",
			"MinimunCapacity":    92,
			"MinimunTemperature": 10,
			"LocalityId":         1,
		}
		mockRepository.On("PatchWarehouse", 1, mock.Anything).Return(models.Warehouse{}, repositoryErr.ErrWarehouseNotFound)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.PatchWarehouse(1, warehouseFakeMap)

		require.ErrorIs(t, err, serviceInterface.ErrWarehouseNotFound)
		require.IsType(t, models.Warehouse{}, response)
		mockRepository.AssertExpectations(t)
	})

	t.Run("update warehouse conflict", func(t *testing.T) {
		mockRepository := repository.NewWarehouseRepositoryMock()
		warehouseFakeMap := map[string]interface{}{
			"Id":                 1,
			"WarehouseCode":      "96b5c517-cce3-4ed3-a06b-24234ee8d009",
			"Address":            "PO Box 16130",
			"Telephone":          "(617) 4591159",
			"MinimunCapacity":    92,
			"MinimunTemperature": 10,
			"LocalityId":         1,
		}
		mockRepository.On("PatchWarehouse", 1, mock.Anything).Return(models.Warehouse{}, repositoryErr.ErrWarehouseCodeDuplicate)
		serviceImp := NewWarehouseService(mockRepository)

		response, err := serviceImp.PatchWarehouse(1, warehouseFakeMap)

		require.ErrorIs(t, err, serviceInterface.ErrWarehouseCodeDuplicate)
		require.IsType(t, models.Warehouse{}, response)
		mockRepository.AssertExpectations(t)
	})
}
