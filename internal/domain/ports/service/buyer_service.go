package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type BuyerService interface {
	GetAll() ([]models.Buyer, error)
	GetById(id int) (models.Buyer, error)
	Create(buyer models.BuyerAttributes) (models.Buyer, error)
	Patch(id int, buyerToPatch models.BuyerAttributes) (models.Buyer, error)
	Delete(id int) error
	GetReportPurchaseOrders(buyerID *int) ([]models.ReportPurchaseOrders, error)
}

var (
	NoRegisteredBuyersYet       = errors.New("No registered buyers yet")
	BuyerDoesNotExist           = errors.New("The requested buyer does not exist in the database for the ID: %d")
	BuyerAlreadyExists          = errors.New("A buyer already exists in the database with the card number ID: %s")
	CannotDeleteBuyerWithOrders = errors.New("We cannot delete the requested buyer because it has orders")
	BuyerHasNoOrders            = errors.New("The requested buyer has no orders yet")
)
