package error_management

import (
	"fmt"
)

type SellerErrors struct {
	Msg string
}

func (se *SellerErrors) Error() string {
	return fmt.Sprintf("%s", se.Msg)
}

var ErrSellerNotFound *SellerErrors = &SellerErrors{
	Msg: "not found error",
}

var ErrSellerAlreadyExists *SellerErrors = &SellerErrors{
	Msg: "seller already exists",
}
