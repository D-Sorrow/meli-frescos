package error_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

var (
	ErrPurchaseOrderInvalidID             = errors.New("ERR: Invalid purchase order ID format")
	ErrPurchaseOrderInvalidJSON           = errors.New("ERR: Invalid purchase order JSON format")
	ErrPurchaseOrderFKWareHouseIdNotValid = errors.New(
		"ERR: The foreign key for the warehouse ID of the requested purchase order is not valid",
	)
	ErrPurchaseOrderFKBuyerIdNotValid = errors.New(
		"ERR: The foreign key for the buyer ID of the requested purchase order is not valid",
	)
	ErrPurchaseOrderFKOrderStatusIdNotValid = errors.New(
		"ERR: The foreign key for the order status ID of the requested purchase order is not valid",
	)
	ErrPurchaseOrderFKCarrierIdNotValid = errors.New(
		"ERR: The foreign key for the carrier ID of the requested purchase order is not valid",
	)
	ErrPurchaseOrderDoesNotExist = errors.New(
		"ERR: The requested purchase order does not exist in the database for the ID: %d",
	)
	ErrPurchaseOrderUnexpectedError = errors.New(
		"ERR: An unexpected error occurred while processing the requested purchase order, please try again later",
	)
)

type PurchaseOrderHandlerError struct {
	Code int
	Msg  string
	Data interface{}
}

func (b PurchaseOrderHandlerError) Error() string {
	return fmt.Sprintf("ERROR: %s", b.Msg)
}

func HandlePurchaseOrderHandlerError(
	err error,
	messages map[string]string,
	args map[string]interface{},
	ctx *context.Context,
) PurchaseOrderHandlerError {
	buyerHandlerErrors := map[error]func() PurchaseOrderHandlerError{
		ErrPurchaseOrderInvalidID: func() PurchaseOrderHandlerError {
			return PurchaseOrderHandlerError{
				Code: http.StatusBadRequest,
				Msg:  ErrPurchaseOrderInvalidID.Error(),
				Data: messages}
		},
		ErrPurchaseOrderInvalidJSON: func() PurchaseOrderHandlerError {
			return PurchaseOrderHandlerError{
				Code: http.StatusBadRequest,
				Msg:  ErrPurchaseOrderInvalidJSON.Error(),
				Data: messages}
		},
		service.ErrPurchaseOrderFKWareHouseIdNotValid: func() PurchaseOrderHandlerError {
			return PurchaseOrderHandlerError{
				Code: http.StatusConflict,
				Msg:  ErrPurchaseOrderFKWareHouseIdNotValid.Error(),
				Data: messages}
		},
		service.ErrPurchaseOrderFKBuyerIdNotValid: func() PurchaseOrderHandlerError {
			return PurchaseOrderHandlerError{
				Code: http.StatusConflict,
				Msg:  ErrPurchaseOrderFKBuyerIdNotValid.Error(),
				Data: messages}
		},
		service.ErrPurchaseOrderFKOrderStatusIdNotValid: func() PurchaseOrderHandlerError {
			return PurchaseOrderHandlerError{
				Code: http.StatusConflict,
				Msg:  ErrPurchaseOrderFKOrderStatusIdNotValid.Error(),
				Data: messages}
		},
		service.ErrPurchaseOrderFKCarrierIdNotValid: func() PurchaseOrderHandlerError {
			return PurchaseOrderHandlerError{
				Code: http.StatusConflict,
				Msg:  ErrPurchaseOrderFKCarrierIdNotValid.Error(),
				Data: messages}
		},
		service.ErrPurchaseOrderDoesNotExist: func() PurchaseOrderHandlerError {
			msg := ErrPurchaseOrderDoesNotExist.Error()
			if id, ok := args["ID"].(int); ok {
				msg = fmt.Sprintf(ErrPurchaseOrderDoesNotExist.Error(), id)
			}

			return PurchaseOrderHandlerError{
				Code: http.StatusNotFound,
				Msg:  msg,
				Data: messages}
		},
	}

	for buyerHandlerErr, errorFunc := range buyerHandlerErrors {
		if errors.Is(err, buyerHandlerErr) {
			return errorFunc()
		}
	}

	return PurchaseOrderHandlerError{
		Code: http.StatusInternalServerError,
		Msg:  ErrPurchaseOrderUnexpectedError.Error(),
	}
}
