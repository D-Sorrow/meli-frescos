package error_management

import (
	"errors"
)

var (
	ErrNoRegisteredBuyersYet = errors.New("No registered buyers yet")
	ErrDuplicateCardNumberID = errors.New("Duplicate card number ID provided")
	ErrBuyerNotFoundWithID   = errors.New("The requested buyer does not exist with ID provided")
)
