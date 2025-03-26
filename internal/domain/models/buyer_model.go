package models

type ReportPurchaseOrders struct {
	ID                  int64
	CardNumberID        string
	FirstName           string
	LastName            string
	PurchaseOrdersCount int
}

type BuyerAttributes struct {
	CardNumberID *string
	FirstName    *string
	LastName     *string
}

type Buyer struct {
	ID int64
	BuyerAttributes
}
