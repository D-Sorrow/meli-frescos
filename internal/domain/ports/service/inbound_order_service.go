package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type InboundOrderService interface {
	CreateInboundOrder(inboundOrder *models.InboundOrder) error
}

var (
	ErrInboundOrderNumberAlreadyExists    = errors.New("order number already exists")
	ErrInboundOrderEmployeeIdNotFound     = errors.New("employee id not found")
	ErrInboundOrderProductBatchIdNotFound = errors.New("product batch id not found")
	ErrInboundOrderWareHouseIdNotFound    = errors.New("warehouse id not found")
	ErrInboundOrderLastInsertId           = errors.New("error with last insert id")
	ErrInboundOrderServiceGeneric         = errors.New("internal server error")
	ErrInboundOrderDateInvalid            = errors.New("incorrect order date value")
)
