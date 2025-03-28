package error_management

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/middlewares"
)

const (
	messageSellerAlreadyExistsError  = "seller already exist"
	messageSellerIdNotExistsError    = "seller id not found"
	messageSellerInternalServerError = "internal server error"
)

type SellerHandlerErrors struct {
	Code int
	Msg  string
}

var sellerServiceErrors = map[error]SellerHandlerErrors{
	service.ErrSellerAlreadyExists: {
		Code: http.StatusConflict,
		Msg:  messageSellerAlreadyExistsError,
	},
	service.ErrSellerNotFound: {
		Code: http.StatusNotFound,
		Msg:  messageSellerIdNotExistsError,
	},
	service.ErrSellerServiceGeneric: {
		Code: http.StatusInternalServerError,
		Msg:  messageSellerInternalServerError,
	},
}

func getSellerErrorMessage(err error) SellerHandlerErrors {
	if e, exists := sellerServiceErrors[err]; exists {
		return e
	}
	return sellerServiceErrors[service.ErrSellerServiceGeneric]
}

func HandleErrorSeller(err error, ctx *context.Context) SellerHandlerErrors {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	switch e := err.(type) {
	case *json.UnmarshalTypeError:
		return SellerHandlerErrors{Code: http.StatusBadRequest, Msg: fmt.Sprintf("the field '%s' must be a '%s'", err.(*json.UnmarshalTypeError).Field, err.(*json.UnmarshalTypeError).Type)}
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
		return SellerHandlerErrors{Code: http.StatusUnprocessableEntity, Msg: errors}
	case *json.SyntaxError:
		return SellerHandlerErrors{Code: http.StatusBadRequest, Msg: "Bad Request - invalid JSON structure"}
	default:
		return getSellerErrorMessage(err)
	}
}
