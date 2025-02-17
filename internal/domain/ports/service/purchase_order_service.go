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
	ForeignKeysNotValid       = errors.New("The foreign keys of the requested purchase order are not valid")
	PurchaseOrderDoesNotExist = errors.New("The requested purchase order does not exist in the database for the ID: %d")
)
