package service

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
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
	ErrBuyerNoRegisteredBuyersYet       = errors.New("ERR_SRV_BY_NO_REG_BYS_YET")
	ErrBuyerAlreadyExists               = errors.New("ERR_SRV_BY_ALREADY_EXISTS")
	ErrBuyerDoesNotExist                = errors.New("ERR_SRV_BY_DOES_NOT_EXIST")
	ErrBuyerCannotDeleteBuyerWithOrders = errors.New("ERR_SRV_BY_CANNOT_DEL_BY_WITH_ORDERS")
	ErrBuyerHasNoOrders                 = errors.New("ERR_SRV_BY_HAS_NO_ORDERS")
	ErrBuyerUnexpectedError             = errors.New("ERR_SRV_BY_UNEXPECTED_ERROR")
)
