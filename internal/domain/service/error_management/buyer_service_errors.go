package error_management

import "errors"

var (
	NoRegisteredBuyersYet = errors.New("No registered buyers yet")
	BuyerDoesNotExist     = errors.New("The requested buyer does not exist in the database for the ID: %d")
	BuyerAlreadyExists    = errors.New("A buyer already exists in the database with the card number ID: %s")
)
