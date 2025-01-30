package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/entity"
)

func JsonToWarehouseModel(warehousesJSON []entity.WarehouseJSON) map[int]models.Warehouse {
	warehouses := make(map[int]models.Warehouse)

	for _, warehouseJSON := range warehousesJSON {
		warehouses[warehouseJSON.Id] = models.Warehouse{
			Id:                 warehouseJSON.Id,
			WarehouseCode:      warehouseJSON.WarehouseCode,
			Address:            warehouseJSON.Address,
			Telephone:          warehouseJSON.Telephone,
			MinimunCapacity:    warehouseJSON.MinimunCapacity,
			MinimunTemperature: warehouseJSON.MinimunTemperature,
		}
	}

	return warehouses
}
