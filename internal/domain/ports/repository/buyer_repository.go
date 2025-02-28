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
	ErrBuyerNoRegisteredBuyersYet       = errors.New("ERR_REPO_BY_NO_REG_BYS_YET")
	ErrBuyerDuplicateCardNumberID       = errors.New("ERR_REPO_BY_DUP_CARD_NUM_ID")
	ErrBuyerNotFoundWithID              = errors.New("ERR_REPO_BY_NOT_FOUND_ID")
	ErrBuyerCannotDeleteBuyerWithOrders = errors.New("ERR_REPO_BY_CANNOT_DEL_BY_WITH_ORDERS")
	ErrBuyerHasNoOrders                 = errors.New("ERR_REPO_BY_HAS_NO_ORDERS")
	ErrBuyerUnexpectedError             = errors.New("ERR_REPO_BY_UNEXPECTED_ERROR")
)
