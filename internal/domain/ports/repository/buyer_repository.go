package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

type BuyerRepository interface {
	GetAll() ([]entities.BuyerEntity, error)
	GetById(id int) (entities.BuyerEntity, error)
	Create(buyer entities.BuyerEntity) (entities.BuyerEntity, error)
	Patch(id int, buyerToPatch entities.BuyerEntity) (entities.BuyerEntity, error)
	Delete(id int) error
	GetReportPurchaseOrders(buyerID *int) ([]entities.ReportPurchaseOrdersEntity, error)
}

var (
	ErrNoRegisteredBuyersYet           = errors.New("No registered buyers yet")
	ErrDuplicateCardNumberID           = errors.New("Duplicate card number ID provided")
	ErrBuyerNotFoundWithID             = errors.New("The requested buyer does not exist with ID provided")
	ErrCannotDeleteBuyerWithOrders     = errors.New("We cannot delete the requested buyer because it has orders")
	ErrBuyerNotFoundOrBuyerHasNoOrders = errors.New("The requested buyer does not exist with ID or the requested buyer has no orders")
)
