package error_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/middlewares"
	"github.com/go-playground/validator/v10"
)

var (
	ErrInboundOrderBodyDecoding   = errors.New("error decoding body")
	ErrInboundOrderHandlerDefault = errors.New("internal server error")
)

const (
	messageInboundOrderNumberAlreadyExistsError     = "El número de la orden ya existe."
	messageInboundOrderEmployeeIdNotExistsError     = "El ID del empleado no existe"
	messageInboundOrderProductBatchIdNotExistsError = "El ID del product batch no existe"
	messageInboundOrderWareHouseIdNotExistsError    = "El ID del warehouse no existe"
	messageInboundOrderBodyMalformedError           = "El cuerpo de la petición está mal formado"
	messageInboundOrderLastInsertIdError            = "Error con último ID"
	messageInboundOrderInternalServerError          = "Internal server error"
	messageInboundOrderDateFormatError              = "El formato de la fecha es inválido (AAAA-MM-DD)"
)

type HandlerErrorInboundOrder struct {
	Code    int
	Message string
}

var inboundOrderServiceErrors = map[error]HandlerErrorInboundOrder{
	service.ErrInboundOrderNumberAlreadyExists: {
		Code:    http.StatusConflict,
		Message: messageInboundOrderNumberAlreadyExistsError,
	},
	service.ErrInboundOrderEmployeeIdNotFound: {
		Code:    http.StatusConflict,
		Message: messageInboundOrderEmployeeIdNotExistsError,
	},
	service.ErrInboundOrderProductBatchIdNotFound: {
		Code:    http.StatusConflict,
		Message: messageInboundOrderProductBatchIdNotExistsError,
	},
	service.ErrInboundOrderWareHouseIdNotFound: {
		Code:    http.StatusConflict,
		Message: messageInboundOrderWareHouseIdNotExistsError,
	},
	service.ErrInboundOrderLastInsertId: {
		Code:    http.StatusInternalServerError,
		Message: messageInboundOrderLastInsertIdError,
	},
	service.ErrInboundOrderServiceGeneric: {
		Code:    http.StatusInternalServerError,
		Message: messageInboundOrderInternalServerError,
	},
	service.ErrInboundOrderDateInvalid: {
		Code:    http.StatusBadRequest,
		Message: messageInboundOrderDateFormatError,
	},
	ErrInboundOrderBodyDecoding: {
		Code:    http.StatusBadRequest,
		Message: messageInboundOrderBodyMalformedError,
	},
}

func getInboundOrderErrorMessage(err error) HandlerErrorInboundOrder {
	if e, exists := inboundOrderServiceErrors[err]; exists {
		return e
	}
	return inboundOrderServiceErrors[service.ErrInboundOrderServiceGeneric]
}

func HandleErrorInboundOrder(err error, ctx *context.Context) HandlerErrorInboundOrder {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	switch e := err.(type) {
	case validator.ValidationErrors:
		errors := "Validación fallida: "
		for _, fieldErr := range e {
			errors = fmt.Sprintf("%v %v %v, ", errors, fieldErr.Tag(), fieldErr.Field())
		}
		return HandlerErrorInboundOrder{Code: http.StatusBadRequest, Message: errors}
	default:
		return getInboundOrderErrorMessage(err)
	}
}
