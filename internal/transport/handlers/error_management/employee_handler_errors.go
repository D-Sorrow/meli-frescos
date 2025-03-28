package error_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/middlewares"
)

var (
	ErrEmployeeBodyDecoding   = errors.New("error decoding body")
	ErrEmployeeHandlerDefault = errors.New("internal server error")
)

const (
	messageEmployeeNotFoundError      = "Empleado no encontrado"
	messageEmployeeAlreadyExistsError = "Empleado con ese card ID ya existe"
	messageEmployeeIdNotValidError    = "El formato del ID no es válido"
	messageEmployeeInternalError      = "Internal server error"
	messageEmployeeBodyMalformedError = "El cuerpo de la petición está mal formado"
)

type HandlerErrorEmployee struct {
	Code    int
	Message string
}

var employeeHandlerErrors = map[error]HandlerErrorEmployee{
	service.ErrEmployeeNotFound: {
		Code:    http.StatusNotFound,
		Message: messageEmployeeNotFoundError,
	},
	service.ErrEmployeeAlreadyExists: {
		Code:    http.StatusConflict,
		Message: messageEmployeeAlreadyExistsError,
	},
	service.ErrEmployeeDecodingError: {
		Code:    http.StatusBadRequest,
		Message: messageEmployeeIdNotValidError,
	},
	service.ErrEmployeeServiceDefault: {
		Code:    http.StatusInternalServerError,
		Message: messageEmployeeInternalError,
	},
	ErrEmployeeBodyDecoding: {
		Code:    http.StatusBadRequest,
		Message: messageEmployeeBodyMalformedError,
	},
	ErrEmployeeHandlerDefault: {
		Code:    http.StatusInternalServerError,
		Message: messageEmployeeInternalError,
	},
}

func getErrorEmployee(err error) HandlerErrorEmployee {
	if e, exists := employeeHandlerErrors[err]; exists {
		return e
	}
	return employeeHandlerErrors[ErrEmployeeHandlerDefault]
}

func HandleErrorEmployee(err error, ctx *context.Context) HandlerErrorEmployee {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	switch e := err.(type) {
	case *strconv.NumError:
		return HandlerErrorEmployee{
			Code:    http.StatusBadRequest,
			Message: messageEmployeeIdNotValidError,
		}
	case validator.ValidationErrors:
		errors := "Validación fallida: "
		for _, fieldErr := range e {
			errors = fmt.Sprintf("%v %v %v, ", errors, fieldErr.Tag(), fieldErr.Field())
		}
		return HandlerErrorEmployee{
			Code:    http.StatusUnprocessableEntity,
			Message: errors,
		}
	default:
		handlerEmployeeError := getErrorEmployee(err)
		return HandlerErrorEmployee{
			Code:    handlerEmployeeError.Code,
			Message: handlerEmployeeError.Message,
		}
	}
}
