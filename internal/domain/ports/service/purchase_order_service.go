package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type PurchaseOrderService interface {
	GetById(id int) (models.PurchaseOrder, error)
	Create(purchaseOrder models.PurchaseOrderAttributesFKs) (models.PurchaseOrder, error)
}

var (
	ErrPurchaseOrderFKWareHouseIdNotValid   = errors.New("ERR_SRV_PO_FK_WAREHOUSE_ID_NOT_VALID")
	ErrPurchaseOrderFKBuyerIdNotValid       = errors.New("ERR_SRV_PO_FK_BUYER_ID_NOT_VALID")
	ErrPurchaseOrderFKOrderStatusIdNotValid = errors.New("ERR_SRV_PO_FK_ORDER_STATUS_ID_NOT_VALID")
	ErrPurchaseOrderFKCarrierIdNotValid     = errors.New("ERR_SRV_PO_FK_CARRIER_ID_NOT_VALID")
	ErrPurchaseOrderDoesNotExist            = errors.New("ERR_SRV_PO_DOES_NOT_EXIST")
	ErrPurchaseOrderUnexpectedError         = errors.New("ERR_SRV_PO_UNEXPECTED_ERROR")
)
