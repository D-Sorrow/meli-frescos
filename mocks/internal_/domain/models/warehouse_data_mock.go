package models

import (
	warehouseModel "github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

var WarehousesFake = map[int]warehouseModel.Warehouse{
	1: {
		Id:                 1,
		WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
		Address:            "Apt 1639",
		Telephone:          "(639) 5350508",
		MinimunCapacity:    99,
		MinimunTemperature: -14,
		LocalityId:         1,
	},
	2: {
		Id:                 2,
		WarehouseCode:      "b6b225e6-c83f-46a4-ac63-b6df8794ba59",
		Address:            "Room 192",
		Telephone:          "(917) 6928569",
		MinimunCapacity:    15,
		MinimunTemperature: 0,
		LocalityId:         2,
	},
}

var WarehouseFakeMap = map[string]interface{}{
	"Id":                 1,
	"WarehouseCode":      "96b5c517-cce3-4ed3-a06b-24234ee8d009",
	"Address":            "PO Box 16130",
	"Telephone":          "(617) 4591159",
	"MinimunCapacity":    92,
	"MinimunTemperature": 10,
	"LocalityId":         1,
}

var ResponseWarehouseDto = []dto.WarehouseDto{
	{
		Id:                 1,
		WarehouseCode:      "6e9168d9-ae9f-46be-a541-959f0cc2a650",
		Address:            "Apt 1639",
		Telephone:          "(639) 5350508",
		MinimumCapacity:    99,
		MinimumTemperature: -14,
		LocalityId:         1,
	},
	{
		Id:                 2,
		WarehouseCode:      "b6b225e6-c83f-46a4-ac63-b6df8794ba59",
		Address:            "Room 192",
		Telephone:          "(917) 6928569",
		MinimumCapacity:    15,
		MinimumTemperature: 0,
		LocalityId:         2,
	},
}
