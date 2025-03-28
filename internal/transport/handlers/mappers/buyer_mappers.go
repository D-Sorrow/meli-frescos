package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
)

func BuyerDTOToBuyer(b *dto.BuyerDTO) *models.Buyer {
	return &models.Buyer{
		ID: b.ID,
		BuyerAttributes: models.BuyerAttributes{
			CardNumberID: &b.CardNumberID,
			FirstName:    &b.FirstName,
			LastName:     &b.LastName,
		},
	}
}

func BuyerCreateDTOToBuyerAttributes(b *dto.BuyerCreateDTO) *models.BuyerAttributes {
	return &models.BuyerAttributes{
		CardNumberID: b.CardNumberID,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}

func BuyerPatchDTOToBuyerPatchAttributes(b *dto.BuyerPatchDTO) *models.BuyerAttributes {
	return &models.BuyerAttributes{
		CardNumberID: b.CardNumberID,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}

func BuyerToBuyerDTO(v *models.Buyer) *dto.BuyerDTO {
	return &dto.BuyerDTO{
		ID:           v.ID,
		CardNumberID: *v.BuyerAttributes.CardNumberID,
		FirstName:    *v.BuyerAttributes.FirstName,
		LastName:     *v.BuyerAttributes.LastName,
	}
}

func ReportPurchaseOrdersToReportPurchaseOrdersDTO(
	v *models.ReportPurchaseOrders,
) *dto.ReportPurchaseOrdersDTO {
	return &dto.ReportPurchaseOrdersDTO{
		ID:                  v.ID,
		CardNumberID:        v.CardNumberID,
		FirstName:           v.FirstName,
		LastName:            v.LastName,
		PurchaseOrdersCount: v.PurchaseOrdersCount,
	}
}
