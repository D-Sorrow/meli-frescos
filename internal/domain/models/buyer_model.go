package models

type BuyerAttributes struct {
	CardNumberID string
	FirstName    string
	LastName     string
}

type BuyerPatchAttributes struct {
	CardNumberID *string
	FirstName    *string
	LastName     *string
}

type Buyer struct {
	ID int
	BuyerAttributes
}
