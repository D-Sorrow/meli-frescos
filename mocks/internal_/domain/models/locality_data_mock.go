package models

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	localityModels "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	//"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
)

var LocalityRequest = models.Locality{
	Id:           1,
	ZipCode:      "10702-013",
	Name:         "Litvínovice",
	ProvinceName: "Boyaca",
	CountryName:  "Tunja",
}

var LocalityBadRequest = models.Locality{
	Id:          1,
	ZipCode:     "10702-013",
	CountryName: "Tunja",
}

var JsonLocalityCreatedDto = dto.LocalityDto{
	Id:           1,
	ZipCode:      "10702-013",
	LocalityName: "Litvínovice",
	ProvinceName: "Boyaca",
	CountryName:  "Tunja",
}

var JsonBadLocalityCreatedDto = dto.LocalityDto{
	Id:          1,
	ZipCode:     "10702-013",
	CountryName: "Tunja",
}

var localityID = 1
var zipCode = "10702-013"
var name = "Litvínovice"
var sellersCount = 5

var LocalitySellersResponse = models.LocalitySellers{

	LocalityId:   &localityID,
	ZipCode:      &zipCode,
	Name:         &name,
	SellersCount: &sellersCount,
}

var JsonLocalitySellersResponse = dto.LocalitySellersDto{

	LocalityId:   localityID,
	ZipCode:      zipCode,
	Name:         name,
	SellersCount: sellersCount,
}

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
