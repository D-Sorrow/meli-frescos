package models

import (
	carrierModels "github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

var Carriers = []carrierModels.Carrier{
	{
		Id:          1,
		Cid:         "#086",
		CompanyName: "Livetube",
		Address:     "Room 192",
		Telephone:   "374-776-3015",
		LocalityId:  1,
	},
	{
		Id:          2,
		Cid:         "#11e",
		CompanyName: "Feedbug",
		Address:     "Apt 308",
		Telephone:   "495-573-9329",
		LocalityId:  2,
	},
}

var ResponseDto = []dto.CarrierDto{
	{
		Id:          1,
		Cid:         "#086",
		CompanyName: "Livetube",
		Address:     "Room 192",
		Telephone:   "374-776-3015",
		LocalityId:  1,
	},
	{
		Id:          2,
		Cid:         "#11e",
		CompanyName: "Feedbug",
		Address:     "Apt 308",
		Telephone:   "495-573-9329",
		LocalityId:  2,
	},
}

// [{
// 	"id": 1,
// 	"cid": "#086",
// 	"company_name": "Livetube",
// 	"address": "Room 192",
// 	"telephone": "374-776-3015",
// 	"locality_id": 1
//   }, {
// 	"id": 2,
// 	"cid": "#11e",
// 	"company_name": "Feedbug",
// 	"address": "Apt 308",
// 	"telephone": "495-573-9329",
// 	"locality_id": 2
//   }]
