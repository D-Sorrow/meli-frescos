package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToLocality(locality dto.LocalityDto) models.Locality {
	return models.Locality{
		Id:         locality.Id,
		ZipCode:    locality.ZipCode,
		Name:       locality.LocalityName,
		ProvinceId: locality.ProvinceId,
	}
}

func MapperToLocalityDTO(locality models.Locality) dto.LocalityDto {
	return dto.LocalityDto{
		Id:           locality.Id,
		ZipCode:      locality.ZipCode,
		LocalityName: locality.Name,
		ProvinceId:   locality.ProvinceId,
	}
}

func MapperToLocalitySellers(localitySellers dto.LocalitySellersDto) models.LocalitySellers {
	return models.LocalitySellers{
		LocalityId:   &localitySellers.LocalityId,
		ZipCode:      &localitySellers.ZipCode,
		Name:         &localitySellers.Name,
		SellersCount: &localitySellers.SellersCount,
	}
}

func MapperToLocalitySellersDTO(localitySellers models.LocalitySellers) dto.LocalitySellersDto {
	return dto.LocalitySellersDto{
		LocalityId:   *localitySellers.LocalityId,
		ZipCode:      *localitySellers.ZipCode,
		Name:         *localitySellers.Name,
		SellersCount: *localitySellers.SellersCount,
	}
}
