package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"net/http"
)

type ErrHandlerProductRecord struct {
	Message string
	Code    int
}

const (
	messageProductRecordNotFound = "product not found"
	messageProductRecordUnknown  = "product record unknown error"
	messageProductRecordBusiness = "product record business error"
)

var productRecordHandlerErrors = map[error]ErrHandlerProductRecord{
	service.ErrServiceProductRecordNotFound:      {Message: messageProductRecordNotFound, Code: http.StatusConflict},
	service.ErrServiceProductRecordBusinessRules: {Message: messageProductRecordBusiness, Code: http.StatusUnprocessableEntity},
	service.ErrServiceProductRecordUnknown:       {Message: messageProductRecordUnknown, Code: http.StatusInternalServerError},
}

func (e *ErrHandlerProductRecord) Error() string {
	return e.Message
}
func (e *ErrHandlerProductRecord) GetCode() int {
	return e.Code
}
func getErrorProductRecord(err error) ErrHandlerProductRecord {
	if e, exists := productRecordHandlerErrors[err]; exists {
		return e
	}
	if err.Error() == "EOF" {
		return ErrHandlerProductRecord{
			Message: "Error, data invalid",
			Code:    500,
		}
	}
	return productRecordHandlerErrors[service.ErrServiceProductRecordUnknown]
}
func HandlerErrProductRecord(err error) ErrHandlerProductRecord {
	switch err.(type) {
	default:
		return getErrorProductRecord(err)
	}
}
