package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

type PurchaseOrderRepository interface {
	GetById(id int) (entities.PurchaseOrderEntity, error)
	Create(purchaseOrder entities.PurchaseOrderEntity) (entities.PurchaseOrderEntity, error)
}

var (
	ErrForeignKeysNotValid         = errors.New("The foreign keys of the requested purchase order are not valid")
	ErrPurchaseOrderNotFoundWithID = errors.New("The requested purchase order does not exist with ID provided")
)
