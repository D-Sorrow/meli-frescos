package error_management

import (
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
)

type HandlerErrorEmployee struct {
	Code    string
	Message string
}

var errorMessages = map[string]string{
	"ENF-SV":       "Empleado no encontrado",
	"EAE-SV":       "Empleado con ese card ID ya existe",
	"ID-DEC-ERR":   "El formato del ID no es válido",
	"BODY-DEC-ERR": "El cuerpo de la petición está mal formado",
	"default":      "Internal server error",
}

func getErrorMessage(code string) string {
	if msg, exists := errorMessages[code]; exists {
		return msg
	}
	return errorMessages["default"]
}

func getHttpStatusCode(errorCode string) int {
	fmt.Printf("errorCode: %s", errorCode)
	switch errorCode {
	case "ENF-SV":
		return http.StatusNotFound
	case "EAE-SV", "ID-DEC-ERR", "BODY-DEC-ERR":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func HandleErrorEmployee(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errors[fieldErr.Field()] = fmt.Sprintf("Validación fallida:  %s", fieldErr.Tag())
		}
		response.JSON(w, http.StatusBadRequest, map[string]any{"error": errors})
		return
	}

	httpStatusCode := getHttpStatusCode(err.Error())
	response.JSON(w, httpStatusCode, map[string]any{
		"error": getErrorMessage(err.Error()),
	})
}
