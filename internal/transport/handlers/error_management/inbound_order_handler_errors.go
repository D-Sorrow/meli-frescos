package error_management

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/bootcamp-go/web/response"
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
	service.ErrInboundOrderNumberAlreadyExists:    {Code: http.StatusConflict, Message: messageInboundOrderNumberAlreadyExistsError},
	service.ErrInboundOrderEmployeeIdNotFound:     {Code: http.StatusConflict, Message: messageInboundOrderEmployeeIdNotExistsError},
	service.ErrInboundOrderProductBatchIdNotFound: {Code: http.StatusConflict, Message: messageInboundOrderProductBatchIdNotExistsError},
	service.ErrInboundOrderWareHouseIdNotFound:    {Code: http.StatusConflict, Message: messageInboundOrderWareHouseIdNotExistsError},
	service.ErrInboundOrderLastInsertId:           {Code: http.StatusInternalServerError, Message: messageInboundOrderLastInsertIdError},
	service.ErrInboundOrderServiceGeneric:         {Code: http.StatusInternalServerError, Message: messageInboundOrderInternalServerError},
	service.ErrInboundOrderDateInvalid:            {Code: http.StatusBadRequest, Message: messageInboundOrderDateFormatError},
	ErrInboundOrderBodyDecoding:                   {Code: http.StatusBadRequest, Message: messageInboundOrderBodyMalformedError},
}

func getInboundOrderErrorMessage(err error) HandlerErrorInboundOrder {
	if e, exists := inboundOrderServiceErrors[err]; exists {
		return e
	}
	return inboundOrderServiceErrors[service.ErrInboundOrderServiceGeneric]
}

func HandleErrorInboundOrder(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case validator.ValidationErrors:
		errors := make(map[string]string)
		for _, fieldErr := range e {
			errors[fieldErr.Field()] = fmt.Sprintf("Validación fallida: %s", fieldErr.Tag())
		}
		response.JSON(w, http.StatusBadRequest, map[string]any{"error": errors})
	default:
		handlerEmployeeError := getInboundOrderErrorMessage(err)
		response.JSON(w, handlerEmployeeError.Code, map[string]any{
			"error": handlerEmployeeError.Message,
		})
	}
}
