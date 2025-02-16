package error_management

import "net/http"

type ErrHandlerProductRecord struct {
	Message string
	Code    int
}

func (e *ErrHandlerProductRecord) Error() string {
	return e.Message
}
func (e *ErrHandlerProductRecord) GetCode() int {
	return e.Code
}
func HandlerErrProductRecord(err error) ErrHandlerProductRecord {
	if err.Error() == "001" {
		return ErrHandlerProductRecord{
			Message: "Error get Product",
			Code:    http.StatusInternalServerError,
		}
	} else if err.Error() == "002" {
		return ErrHandlerProductRecord{
			Message: "Error business rules",
			Code:    http.StatusBadRequest,
		}
	}
	return ErrHandlerProductRecord{
		Message: "an unknown error occurred",
		Code:    http.StatusInternalServerError,
	}
}
