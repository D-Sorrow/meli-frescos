package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToWarehouseDto(warehouse models.Warehouse) dto.WarehouseDto {
	return dto.WarehouseDto{
		Id:                 warehouse.Id,
		WarehouseCode:      warehouse.WarehouseCode,
		Address:            warehouse.Address,
		Telephone:          warehouse.Telephone,
		MinimunCapacity:    warehouse.MinimunCapacity,
		MinimunTemperature: warehouse.MinimunTemperature,
	}
}

func MapperToWarehousesDto(warehouses map[int]models.Warehouse) []dto.WarehouseDto {
	var warehousesDto []dto.WarehouseDto
	for _, warehouse := range warehouses {
		warehouseDto := MapperToWarehouseDto(warehouse)
		warehousesDto = append(warehousesDto, warehouseDto)
	}
	return warehousesDto
}
