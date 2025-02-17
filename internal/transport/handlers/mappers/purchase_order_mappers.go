package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func PurchaseOrderCreateDTOToPurchaseOrderAttributesFKs(p *dto.PurchaseOrderCreateDTO) *models.PurchaseOrderAttributesFKs {
	return &models.PurchaseOrderAttributesFKs{
		PurchaseOrderFKs: models.PurchaseOrderFKs{
			BuyerID:       p.BuyerID,
			CarrierID:     p.CarrierID,
			OrderStatusID: p.OrderStatusID,
			WarehouseID:   p.WarehouseID,
		},
	}
}

func PurchaseOrderToPurchaseOrderDTO(p *models.PurchaseOrder) *dto.PurchaseOrderDTO {
	return &dto.PurchaseOrderDTO{
		ID:            p.ID,
		OrderNumber:   p.PurchaseOrderAttributes.OrderNumber,
		OrderDate:     p.PurchaseOrderAttributes.OrderDate,
		TrackingCode:  p.PurchaseOrderAttributes.TrackingCode,
		BuyerID:       p.PurchaseOrderFKs.BuyerID,
		CarrierID:     p.PurchaseOrderFKs.CarrierID,
		OrderStatusID: p.PurchaseOrderFKs.OrderStatusID,
		WarehouseID:   p.PurchaseOrderFKs.WarehouseID,
	}
}
