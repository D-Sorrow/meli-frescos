package error_management

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

var (
	ErrBuyerInvalidID             = errors.New("ERR: Invalid buyer ID format")
	ErrBuyerInvalidJSON           = errors.New("ERR: Invalid buyer JSON format")
	ErrBuyerInvalidCreateDTO      = errors.New("ERR: Invalid buyer Create DTO")
	ErrBuyerInvalidPatchDTO       = errors.New("ERR: Invalid buyer Patch DTO")
	ErrBuyerNoRegisteredBuyersYet = errors.New("ERR: No registered buyers yet")
	ErrBuyerAlreadyExists         = errors.New(
		"ERR: A buyer already exists in the database with the card number ID: %s",
	)
	ErrBuyerDoesNotExist = errors.New(
		"ERR: The requested buyer does not exist in the database for the ID: %d",
	)
	ErrBuyerCannotDeleteBuyerWithOrders = errors.New(
		"ERR: Cannot delete buyer with orders. Please delete the related purchase orders first",
	)
	ErrBuyerHasNoOrders     = errors.New("ERR: The requested buyer has no orders yet")
	ErrBuyerUnexpectedError = errors.New(
		"ERR: An unexpected error occurred while processing the requested buyer, please try again later",
	)
)

type BuyerHandlerError struct {
	Code int
	Msg  string
	Data interface{}
}

func (b BuyerHandlerError) Error() string {
	return fmt.Sprintf("ERROR: %s", b.Msg)
}

func HandleBuyerHandlerError(
	err error,
	messages map[string]string,
	args map[string]interface{},
) BuyerHandlerError {
	var data interface{}
	if messages != nil {
		data = messages
	}

	buyerHandlerErrors := map[error]func() BuyerHandlerError{
		ErrBuyerInvalidID: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusBadRequest,
				Msg:  ErrBuyerInvalidID.Error(),
				Data: data,
			}
		},
		ErrBuyerInvalidJSON: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusBadRequest,
				Msg:  ErrBuyerInvalidJSON.Error(),
				Data: data,
			}
		},
		ErrBuyerInvalidCreateDTO: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusBadRequest,
				Msg:  ErrBuyerInvalidCreateDTO.Error(),
				Data: data,
			}
		},
		ErrBuyerInvalidPatchDTO: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusBadRequest,
				Msg:  ErrBuyerInvalidPatchDTO.Error(),
				Data: data,
			}
		},
		service.ErrBuyerNoRegisteredBuyersYet: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusOK,
				Msg:  ErrBuyerNoRegisteredBuyersYet.Error(),
				Data: data,
			}
		},
		service.ErrBuyerAlreadyExists: func() BuyerHandlerError {
			msg := ErrBuyerAlreadyExists.Error()
			if cardNumberID, ok := args["CardNumberID"].(string); ok {
				msg = fmt.Sprintf(ErrBuyerAlreadyExists.Error(), cardNumberID)
			}

			return BuyerHandlerError{
				Code: http.StatusConflict,
				Msg:  msg,
				Data: data,
			}
		},
		service.ErrBuyerDoesNotExist: func() BuyerHandlerError {
			msg := ErrBuyerDoesNotExist.Error()
			if id, ok := args["ID"].(int); ok {
				msg = fmt.Sprintf(ErrBuyerDoesNotExist.Error(), id)
			}

			return BuyerHandlerError{
				Code: http.StatusNotFound,
				Msg:  msg,
				Data: data,
			}
		},
		service.ErrBuyerCannotDeleteBuyerWithOrders: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusConflict,
				Msg:  ErrBuyerCannotDeleteBuyerWithOrders.Error(),
				Data: data,
			}
		},
		service.ErrBuyerHasNoOrders: func() BuyerHandlerError {
			return BuyerHandlerError{
				Code: http.StatusOK,
				Msg:  ErrBuyerHasNoOrders.Error(),
				Data: data,
			}
		},
	}

	for buyerHandlerErr, errorFunc := range buyerHandlerErrors {
		if errors.Is(err, buyerHandlerErr) {
			return errorFunc()
		}
	}

	return BuyerHandlerError{
		Code: http.StatusInternalServerError,
		Msg:  ErrBuyerUnexpectedError.Error(),
	}
}
