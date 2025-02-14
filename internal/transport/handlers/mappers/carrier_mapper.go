package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToCarrierDto(carrier models.Carrier) dto.CarrierDto {
	return dto.CarrierDto{
		Id:          carrier.Id,
		Cid:         carrier.Cid,
		CompanyName: carrier.CompanyName,
		Address:     carrier.Address,
		Telephone:   carrier.Telephone,
		LocalityId:  carrier.LocalityId,
	}
}

func MapperToCarrierModel(carrier dto.CarrierDto) models.Carrier {
	return models.Carrier{
		Id:          carrier.Id,
		Cid:         carrier.Cid,
		CompanyName: carrier.CompanyName,
		Address:     carrier.Address,
		Telephone:   carrier.Telephone,
		LocalityId:  carrier.LocalityId,
	}
}
