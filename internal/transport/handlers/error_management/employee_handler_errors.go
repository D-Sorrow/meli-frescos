package error_management

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
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
	service.ErrEmployeeNotFound:       {Code: http.StatusNotFound, Message: messageEmployeeNotFoundError},
	service.ErrEmployeeAlreadyExists:  {Code: http.StatusBadRequest, Message: messageEmployeeAlreadyExistsError},
	service.ErrEmployeeDecodingError:  {Code: http.StatusBadRequest, Message: messageEmployeeIdNotValidError},
	service.ErrEmployeeServiceDefault: {Code: http.StatusInternalServerError, Message: messageEmployeeInternalError},
	ErrEmployeeBodyDecoding:           {Code: http.StatusBadRequest, Message: messageEmployeeBodyMalformedError},
	ErrEmployeeHandlerDefault:         {Code: http.StatusInternalServerError, Message: messageEmployeeInternalError},
}

func getErrorEmployee(err error) HandlerErrorEmployee {
	if e, exists := employeeHandlerErrors[err]; exists {
		return e
	}
	return employeeHandlerErrors[ErrEmployeeHandlerDefault]
}

func HandleErrorEmployee(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *strconv.NumError:
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"error": messageEmployeeIdNotValidError,
		})
	case validator.ValidationErrors:
		errors := make(map[string]string)
		for _, fieldErr := range e {
			errors[fieldErr.Field()] = fmt.Sprintf("Validación fallida: %s", fieldErr.Tag())
		}
		response.JSON(w, http.StatusBadRequest, map[string]any{"error": errors})
	default:
		handlerEmployeeError := getErrorEmployee(err)
		response.JSON(w, handlerEmployeeError.Code, map[string]any{
			"error": handlerEmployeeError.Message,
		})
	}
}
