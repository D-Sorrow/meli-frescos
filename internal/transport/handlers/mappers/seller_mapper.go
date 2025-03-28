package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
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
