package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
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

func MapperToCarriersDto(carriers []models.Carrier) []dto.CarrierDto {
	var carriersDto []dto.CarrierDto
	for _, carrier := range carriers {
		carrierDto := MapperToCarrierDto(carrier)
		carriersDto = append(carriersDto, carrierDto)
	}

	return carriersDto
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
