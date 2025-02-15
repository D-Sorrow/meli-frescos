package models

type PurchaseOrderFKs struct {
	BuyerID       int
	CarrierID     int
	OrderStatusID int
	WarehouseID   int
}

type PurchaseOrderAttributes struct {
	OrderNumber  string
	OrderDate    string
	TrackingCode string
}

type PurchaseOrderAttributesFKs struct {
	PurchaseOrderAttributes
	PurchaseOrderFKs
}

type PurchaseOrder struct {
	ID int
	PurchaseOrderAttributes
	PurchaseOrderFKs
}
