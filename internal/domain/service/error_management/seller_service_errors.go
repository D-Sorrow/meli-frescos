package service_errors

import (
	"fmt"
	"net/http"
)

type SellerErrors struct {
	Code int
	Msg  string
}

func (se *SellerErrors) Error() string {
	return fmt.Sprintf("%d: %s", se.Code, se.Msg)
}

var ErrNotFound *SellerErrors = &SellerErrors{
	Code: http.StatusNotFound,
	Msg:  "not found error",
}

var ErrAlreadyExists *SellerErrors = &SellerErrors{
	Code: http.StatusNotFound,
	Msg:  "seller already exists",
}
