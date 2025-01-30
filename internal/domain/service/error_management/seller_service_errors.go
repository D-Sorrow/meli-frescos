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

var ErrNotFound *SellerErrors = &SellerErrors{
	Msg: "not found error",
}

var ErrAlreadyExists *SellerErrors = &SellerErrors{
	Msg: "seller already exists",
}
