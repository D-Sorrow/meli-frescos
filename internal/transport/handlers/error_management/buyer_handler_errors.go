package error_management

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/bootcamp-go/web/response"
)

const (
	InvalidID          string = "Invalid ID format"
	InvalidJSON        string = "Invalid JSON format"
	InvalidBuyerCreate string = "Invalid Buyer Create DTO"
	InvalidBuyerPatch  string = "Invalid Buyer Patch DTO"
)

type BuyerError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e BuyerError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

func ResponseError(err error, w *http.ResponseWriter) {
	var buyerErr BuyerError

	if errors.As(err, &buyerErr) {
		response.JSON(*w, buyerErr.Code, dto.ResponseDTO{
			Code: buyerErr.Code,
			Msg:  buyerErr.Msg,
		})
		return
	}

	response.JSON(*w, http.StatusInternalServerError, dto.ResponseDTO{
		Code: http.StatusInternalServerError,
		Msg:  "An unexpected error occurred, please try again later",
	})
}
