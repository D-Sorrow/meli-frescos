package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

type PurchaseOrderRepository interface {
	GetById(id int) (entities.PurchaseOrderEntity, error)
	Create(purchaseOrder entities.PurchaseOrderEntity) (entities.PurchaseOrderEntity, error)
}

var (
	ErrPurchaseOrderFKWareHouseIdNotValid   = errors.New("ERR_REPO_PO_FK_WAREHOUSE_ID_NOT_VALID")
	ErrPurchaseOrderFKBuyerIdNotValid       = errors.New("ERR_REPO_PO_FK_BUYER_ID_NOT_VALID")
	ErrPurchaseOrderFKOrderStatusIdNotValid = errors.New("ERR_REPO_PO_FK_ORDER_STATUS_ID_NOT_VALID")
	ErrPurchaseOrderFKCarrierIdNotValid     = errors.New("ERR_REPO_PO_FK_CARRIER_ID_NOT_VALID")
	ErrPurchaseOrderNotFoundWithID          = errors.New("ERR_REPO_PO_NOT_FOUND_ID")
	ErrPurchaseOrderUnexpectedError         = errors.New("ERR_REPO_PO_UNEXPECTED_ERROR")
)
