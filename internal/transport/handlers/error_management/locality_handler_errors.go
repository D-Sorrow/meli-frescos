package error_management

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/middlewares"
	"github.com/go-playground/validator/v10"
)

const (
	messageLocalityAlreadyExistsError  = "locality already exist"
	messageLocalityIdNotExistsError    = "locality id not found"
	messageProvinceNotExistsError      = "province not exist"
	messageLocalityInternalServerError = "internal server error"
)

type LocalityHandlerErrors struct {
	Code int
	Msg  string
}

var localityServiceErrors = map[error]LocalityHandlerErrors{
	service.ErrLocalityAlreadyExists: {
		Code: http.StatusConflict,
		Msg:  messageLocalityAlreadyExistsError,
	},
	service.ErrLocalityNotFound: {
		Code: http.StatusNotFound,
		Msg:  messageLocalityIdNotExistsError,
	},
	service.ErrProvinceNotFound: {
		Code: http.StatusNotFound,
		Msg:  messageProvinceNotExistsError,
	},
	service.ErrLocalityRepositoryGeneric: {
		Code: http.StatusInternalServerError,
		Msg:  messageLocalityInternalServerError,
	},
}

func getLocalityErrorMessage(err error) LocalityHandlerErrors {
	if e, exists := localityServiceErrors[err]; exists {
		return e
	}
	return localityServiceErrors[service.ErrLocalityRepositoryGeneric]
}

func HandleErrorLocality(err error, ctx *context.Context) LocalityHandlerErrors {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	switch e := err.(type) {
	case *json.UnmarshalTypeError:
		return LocalityHandlerErrors{Code: http.StatusBadRequest, Msg: fmt.Sprintf("the field '%s' must be a '%s'", err.(*json.UnmarshalTypeError).Field, err.(*json.UnmarshalTypeError).Type)}
	case validator.ValidationErrors:
		errors := "The request is invalid because it does not contain the necessary fields: "
		for _, fieldErr := range e {
			fieldName := fieldErr.Field()
			if field, ok := reflect.TypeOf(dto.SellerDto{}).FieldByName(fieldName); ok {
				errors += fmt.Sprintf("'%s' is %s and must be a %s, ", field.Tag.Get("json"), fieldErr.Tag(), fieldErr.Kind())
			} else {
				errors += fmt.Sprintf("'%s' is %s and must be a %s, ", fieldName, fieldErr.Tag(), fieldErr.Kind())
			}
		}
		return LocalityHandlerErrors{Code: http.StatusUnprocessableEntity, Msg: errors}
	case *json.SyntaxError:
		return LocalityHandlerErrors{Code: http.StatusBadRequest, Msg: "Bad Request - invalid JSON structure"}
	default:
		return getLocalityErrorMessage(err)
	}
}
