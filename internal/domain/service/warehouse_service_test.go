package service

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repositoryErr "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	serviceInterface "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
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
