package error_management

import (
	"context"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/middlewares"
)

type ErrHandlerProduct struct {
	Message string
	Code    int
}

const (
	messageProductNotFound      = "Product not found"
	messageProductUnknown       = "Product unknown"
	messageProductAlreadyExists = "Product already exists"
)

var productHandlerErrors = map[error]ErrHandlerProduct{
	service.ErrServiceProductNotFound: {
		Message: messageProductNotFound,
		Code:    http.StatusNotFound,
	},
	service.ErrServiceProductUnknown: {
		Message: messageProductUnknown,
		Code:    http.StatusInternalServerError,
	},
	service.ErrServiceProductAlreadyExists: {
		Message: messageProductAlreadyExists,
		Code:    http.StatusConflict,
	},
}

func (e *ErrHandlerProduct) Error() string {
	return e.Message
}
func (e *ErrHandlerProduct) GetCode() int {
	return e.Code
}
func getErrorProduct(err error) ErrHandlerProduct {
	if e, exists := productHandlerErrors[err]; exists {
		return e
	}
	return productHandlerErrors[service.ErrServiceProductNotFound]
}
func HandlerErrorProduct(err error, ctx *context.Context) ErrHandlerProduct {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	switch err.(type) {
	default:
		return getErrorProduct(err)
	}
}
