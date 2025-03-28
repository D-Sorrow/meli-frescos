package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/loader/entity"
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
