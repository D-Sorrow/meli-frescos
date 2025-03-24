package models

import (
	localityModels "github.com/D-Sorrow/meli-frescos/internal/domain/models"
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

// var ResponseCarrierDto = []dto.CarrierDto{
// 	{
// 		Id:          1,
// 		Cid:         "#086",
// 		CompanyName: "Livetube",
// 		Address:     "Room 192",
// 		Telephone:   "374-776-3015",
// 		LocalityId:  1,
// 	},
// 	{
// 		Id:          2,
// 		Cid:         "#11e",
// 		CompanyName: "Feedbug",
// 		Address:     "Apt 308",
// 		Telephone:   "495-573-9329",
// 		LocalityId:  2,
// 	},
// }

// [{
// 	"LocalityId": 1,
// 	"ZipCode": "10702-013",
// 	"Name": "Litvínovice",
// 	"CarriersCount": "38482"
//   }, {
// 	"LocalityId": 2,
// 	"ZipCode": "47682-018",
// 	"Name": "Biritiba Mirim",
// 	"CarriersCount": "316"
//   }]
