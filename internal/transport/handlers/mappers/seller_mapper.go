package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToSeller(seller dto.SellerDto) models.Seller {
	return models.Seller{
		Id:          seller.Id,
		Cid:         seller.Cid,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
		LocalityId:  seller.LocalityId,
	}
}

func MapperToSellerDTO(seller models.Seller) dto.SellerDto {
	return dto.SellerDto{
		Id:          seller.Id,
		Cid:         seller.Cid,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
		LocalityId:  seller.LocalityId,
	}
}

func MapperToSellerPatch(seller dto.SellerUpdateDto) models.SellerPatch {
	return models.SellerPatch{
		Cid:         seller.Cid,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
		LocalityId:  seller.LocalityId,
	}
}
