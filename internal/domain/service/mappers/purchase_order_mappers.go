package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

func PurchaseOrderToPurchaseOrderEntity(v *models.PurchaseOrder) *entities.PurchaseOrderEntity {
	return &entities.PurchaseOrderEntity{
		ID:            v.ID,
		OrderNumber:   v.PurchaseOrderAttributes.OrderNumber,
		OrderDate:     v.PurchaseOrderAttributes.OrderDate,
		TrackingCode:  v.PurchaseOrderAttributes.TrackingCode,
		BuyerID:       v.PurchaseOrderFKs.BuyerID,
		CarrierID:     v.PurchaseOrderFKs.CarrierID,
		OrderStatusID: v.PurchaseOrderFKs.OrderStatusID,
		WarehouseID:   v.PurchaseOrderFKs.WarehouseID,
	}
}

func PurchaseOrderAttributesFKsToPurchaseOrderEntity(v *models.PurchaseOrderAttributesFKs) *entities.PurchaseOrderEntity {
	return &entities.PurchaseOrderEntity{
		OrderNumber:   v.OrderNumber,
		OrderDate:     v.OrderDate,
		TrackingCode:  v.TrackingCode,
		BuyerID:       v.BuyerID,
		CarrierID:     v.CarrierID,
		OrderStatusID: v.OrderStatusID,
		WarehouseID:   v.WarehouseID,
	}
}

func PurchaseOrderEntityToPurchaseOrder(v *entities.PurchaseOrderEntity) *models.PurchaseOrder {
	return &models.PurchaseOrder{
		ID: v.ID,
		PurchaseOrderAttributes: models.PurchaseOrderAttributes{
			OrderNumber:  v.OrderNumber,
			OrderDate:    v.OrderDate,
			TrackingCode: v.TrackingCode,
		},
		PurchaseOrderFKs: models.PurchaseOrderFKs{
			BuyerID:       v.BuyerID,
			CarrierID:     v.CarrierID,
			OrderStatusID: v.OrderStatusID,
			WarehouseID:   v.WarehouseID,
		},
	}
}

func ReportPurchaseOrdersEntityToReportPurchaseOrders(v *entities.ReportPurchaseOrdersEntity) *models.ReportPurchaseOrders {
	return &models.ReportPurchaseOrders{
		ID:                  v.ID,
		CardNumberID:        v.CardNumberID,
		FirstName:           v.FirstName,
		LastName:            v.LastName,
		PurchaseOrdersCount: v.PurchaseOrdersCount,
	}
}
