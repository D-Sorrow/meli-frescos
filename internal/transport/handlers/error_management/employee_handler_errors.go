package error_management

import (
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

var errorMessages = map[string]string{
	"ENF-SV":     "Empleado no encontrado",
	"EAE-SV":     "Empleado con ese card ID ya existe",
	"ID-DEC-ERR": "El formato del ID no es válido",
	"BODY-ERR":   "El cuerpo de la petición es inválido: ",
	"default":    "Internal server error",
}

func getErrorMessage(code string, optionalBodyErr error) string {
	if msg, exists := errorMessages[code]; exists {
		if optionalBodyErr != nil {
			return msg + optionalBodyErr.Error()
		}
		return msg
	}
	return errorMessages["default"]
}

func getHttpStatusCode(errorCode string) int {
	switch errorCode {
	case "ENF-SV":
		return http.StatusNotFound
	case "EAE-SV", "ID-DEC-ERR", "BODY-ERR":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func HandleErrorEmployee(w http.ResponseWriter, errorCode string, optionalBodyErr error) {
	fmt.Printf("error: %s", errorCode)
	httpStatusCode := getHttpStatusCode(errorCode)
	response.JSON(w, httpStatusCode, map[string]any{
		"error": getErrorMessage(errorCode, optionalBodyErr),
	})
}
