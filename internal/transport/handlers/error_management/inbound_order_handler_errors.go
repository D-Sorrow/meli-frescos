package error_management

import (
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
)

var inboundOrderErrorMessages = map[string]string{
	"ONAE_SV":      "El número de la orden ya existe.",
	"BODY-DEC-ERR": "El cuerpo de la petición está mal formado",
	"EIDNF_SV":     "El ID del empleado no existe",
	"PBIDNF_DB":    "El ID del product batch no existe",
	"WHIDNF_DB":    "El ID del warehouse no existe",
}

func getInboundOrderErrorMessage(code string) string {
	if msg, exists := inboundOrderErrorMessages[code]; exists {
		return msg
	}
	return errorMessages["default"]
}

func getInboundOrderHttpStatusCode(errorCode string) int {
	fmt.Printf("errorCode: %s", errorCode)
	switch errorCode {
	case "ONAE_SV", "EIDNF_SV", "PBIDNF_DB", "WHIDNF_DB":
		return http.StatusConflict
	case "BODY-DEC-ERR":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func HandleErrorInboundOrder(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errors[fieldErr.Field()] = fmt.Sprintf("Validación fallida:  %s", fieldErr.Tag())
		}
		response.JSON(w, http.StatusUnprocessableEntity, map[string]any{"error": errors})
		return
	}

	httpStatusCode := getInboundOrderHttpStatusCode(err.Error())
	response.JSON(w, httpStatusCode, map[string]any{
		"error": getInboundOrderErrorMessage(err.Error()),
	})
}
