package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func BuyerDTOToBuyer(b *dto.BuyerDTO) *models.Buyer {
	return &models.Buyer{
		ID: b.ID,
		BuyerAttributes: models.BuyerAttributes{
			CardNumberID: b.CardNumberID,
			FirstName:    b.FirstName,
			LastName:     b.LastName,
		},
	}
}

func BuyerCreateDTOToBuyerAttributes(b *dto.BuyerCreateDTO) *models.BuyerAttributes {
	return &models.BuyerAttributes{
		CardNumberID: *b.CardNumberID,
		FirstName:    *b.FirstName,
		LastName:     *b.LastName,
	}
}

func BuyerPatchDTOToBuyerPatchAttributes(b *dto.BuyerPatchDTO) *models.BuyerPatchAttributes {
	return &models.BuyerPatchAttributes{
		CardNumberID: b.CardNumberID,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}
