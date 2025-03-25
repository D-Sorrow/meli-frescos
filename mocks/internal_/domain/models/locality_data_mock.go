package models

import (
	localityModels "github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	//"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

var LocalityCarriers = []localityModels.LocalityCarriers{
	{
		LocalityId:    1,
		ZipCode:       "10702-013",
		Name:          "Litvínovice",
		CarriersCount: 38,
	},
	{
		LocalityId:    2,
		ZipCode:       "47682-018",
		Name:          "Biritiba Mirim",
		CarriersCount: 4,
	},
}

var ResponseLocalityCarriersDto = []dto.LocalityCarriersDto{
	{
		LocalityId:    1,
		ZipCode:       "10702-013",
		Name:          "Litvínovice",
		CarriersCount: 38,
	},
	{
		LocalityId:    2,
		ZipCode:       "47682-018",
		Name:          "Biritiba Mirim",
		CarriersCount: 4,
	},
}
