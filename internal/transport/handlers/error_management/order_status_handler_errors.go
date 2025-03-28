package error_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

var (
	ErrOrderStatusNoRegisteredOrderStatusesYet = errors.New("ERR: No registered order statuses yet")
	ErrOrderStatusUnexpectedError              = errors.New(
		"ERR: An unexpected error occurred while processing the requested order status, please try again later",
	)
)

type OrderStatusHandlerError struct {
	Code int
	Msg  string
	Data interface{}
}

func (b OrderStatusHandlerError) Error() string {
	return fmt.Sprintf("ERROR: %s", b.Msg)
}

func HandleOrderStatusHandlerError(
	err error,
	messages map[string]string,
	args map[string]interface{},
	ctx *context.Context,
) OrderStatusHandlerError {
	buyerHandlerErrors := map[error]func() OrderStatusHandlerError{
		service.ErrOrderStatusNoRegisteredOrderStatusesYet: func() OrderStatusHandlerError {
			return OrderStatusHandlerError{
				Code: http.StatusOK,
				Msg:  ErrOrderStatusNoRegisteredOrderStatusesYet.Error(),
				Data: messages}
		},
	}

	for buyerHandlerErr, errorFunc := range buyerHandlerErrors {
		if errors.Is(err, buyerHandlerErr) {
			return errorFunc()
		}
	}

	return OrderStatusHandlerError{
		Code: http.StatusInternalServerError,
		Msg:  ErrOrderStatusUnexpectedError.Error(),
	}
}
