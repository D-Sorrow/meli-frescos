package handler_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
)

type SellerHandlerErrors struct {
	Code int
	Msg  string
}

func (se *SellerHandlerErrors) Error() string {
	return fmt.Sprintf("%d: %s", se.Code, se.Msg)
}

var ErrNotFound *SellerHandlerErrors = &SellerHandlerErrors{
	Code: http.StatusNotFound,
	Msg:  "not found error",
}

var ErrAlreadyExists *SellerHandlerErrors = &SellerHandlerErrors{
	Code: http.StatusNotFound,
	Msg:  "seller already exists",
}

var ErrNoExists *SellerHandlerErrors = &SellerHandlerErrors{
	Code: http.StatusNotFound,
	Msg:  "seller doesnt exists",
}

func ResponseError(err error, w http.ResponseWriter) {
	var sellerErr *SellerHandlerErrors
	if errors.As(err, &sellerErr) {
		response.JSON(w, sellerErr.Code, dto.ResponseDTO{
			Code: sellerErr.Code,
			Msg:  sellerErr.Msg,
			Data: nil,
		})
	}

	var validatorErr validator.ValidationErrors
	if errors.As(err, &validatorErr) {
		var validationErrors []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s and must be a %s", validationErr.Field(), validationErr.ActualTag(), validationErr.Kind()))
		}
		response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
			Code: http.StatusBadRequest,
			Msg:  "Bad Request",
			Data: validationErrors,
		})
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
			Code: http.StatusBadRequest,
			Msg:  "Bad Request - invalid JSON structure",
			Data: nil,
		})
	}

}
