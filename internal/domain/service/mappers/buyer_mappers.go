package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

func BuyerAttributesToBuyerEntity(v *models.BuyerAttributes) *entities.BuyerEntity {
	return &entities.BuyerEntity{
		CardNumberID: v.CardNumberID,
		FirstName:    v.FirstName,
		LastName:     v.LastName,
	}
}

func BuyerEntityToBuyer(v *entities.BuyerEntity) *models.Buyer {
	return &models.Buyer{
		ID: v.ID,
		BuyerAttributes: models.BuyerAttributes{
			CardNumberID: v.CardNumberID,
			FirstName:    v.FirstName,
			LastName:     v.LastName,
		},
	}
}
