package error_management

import (
	"net/http"
)

type ErrHandler struct {
	Message string
	Code    int
}

func (e *ErrHandler) Error() string {
	return e.Message
}
func (e *ErrHandler) GetCode() int {
	return e.Code
}
func HandlerErr(err error) ErrHandler {
	if err.Error() == "001" {
		return ErrHandler{
			Message: "Error get Product",
			Code:    http.StatusInternalServerError,
		}
	} else if err.Error() == "002" {
		return ErrHandler{
			Message: "Error get Product",
			Code:    http.StatusNotFound,
		}
	} else if err.Error() == "003" {
		return ErrHandler{
			Message: "Error in this operation",
			Code:    http.StatusBadRequest,
		}
	} else if err.Error() == "004" {
		return ErrHandler{
			Message: "Error business rules",
			Code:    http.StatusBadRequest,
		}
	} else if err.Error() == "005" {
		return ErrHandler{
			Message: "Error delete product",
			Code:    http.StatusNotFound,
		}
	} else if err.Error() == "006" {
		return ErrHandler{
			Message: "Error delete isn't possible",
			Code:    http.StatusBadRequest,
		}
	}
	return ErrHandler{
		Message: "an unknown error occurred",
		Code:    http.StatusInternalServerError,
	}
}
