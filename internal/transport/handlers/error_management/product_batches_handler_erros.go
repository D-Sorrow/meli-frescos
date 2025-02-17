package error_management

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
)

type ProductBatchesHandlerErr struct {
	Code int
	Msg  string
}

func (se *ProductBatchesHandlerErr) Error() string {
	return fmt.Sprintf("%d: %s", se.Code, se.Msg)
}

var LocalityNotFound *ProductBatchesHandlerErr = &ProductBatchesHandlerErr{
	Code: http.StatusNotFound,
	Msg:  "not found error",
}

var LocalityAlreadyExists *ProductBatchesHandlerErr = &ProductBatchesHandlerErr{
	Code: http.StatusConflict,
	Msg:  "ProductBatches  already exists",
}

func ResponseErrorProductBatches(err error, w http.ResponseWriter) {
	var ProductBatchErr *ProductBatchesHandlerErr
	if errors.As(err, &ProductBatchErr) {
		response.JSON(w, ProductBatchErr.Code, dto.ResponseDTO{
			Code: ProductBatchErr.Code,
			Msg:  ProductBatchErr.Msg,
			Data: nil,
		})
	}

	var validatorErr validator.ValidationErrors
	if errors.As(err, &validatorErr) {
		var validationErrors []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s and must be a %s", validationErr.Field(), validationErr.ActualTag(), validationErr.Kind()))
		}
		response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
			Code: http.StatusUnprocessableEntity,
			Msg:  "The request is invalid because it does not contain the necessary fields",
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
