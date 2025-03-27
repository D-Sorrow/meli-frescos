package models

type PurchaseOrderFKs struct {
	BuyerID       int64
	CarrierID     int64
	OrderStatusID int64
	WarehouseID   int64
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
	ID int64
	PurchaseOrderAttributes
	PurchaseOrderFKs
}
