package error_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/middlewares"
)

const (
	InvalidID          string = "Invalid ID format"
	InvalidJSON        string = "Invalid JSON format"
	InvalidBuyerCreate string = "Invalid Buyer Create DTO"
	InvalidBuyerPatch  string = "Invalid Buyer Patch DTO"
)

type HandlerError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e HandlerError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

func HandlerResponseError(err error, w *http.ResponseWriter, ctx *context.Context) {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	var buyerErr HandlerError

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
