package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func BuyerToBuyerDTO(v *models.Buyer) *dto.BuyerDTO {
	return &dto.BuyerDTO{
		ID:           v.ID,
		CardNumberID: v.BuyerAttributes.CardNumberID,
		FirstName:    v.BuyerAttributes.FirstName,
		LastName:     v.BuyerAttributes.LastName,
	}
}
