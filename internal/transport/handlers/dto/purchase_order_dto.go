package dto

type PurchaseOrderDTO struct {
	ID            int    `json:"id"`
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerID       int    `json:"buyer_id"`
	CarrierID     int    `json:"carrier_id"`
	OrderStatusID int    `json:"order_status_id"`
	WarehouseID   int    `json:"warehouse_id"`
}

type PurchaseOrderCreateDTO struct {
	BuyerID       int `json:"buyer_id"`
	CarrierID     int `json:"carrier_id"`
	OrderStatusID int `json:"order_status_id"`
	WarehouseID   int `json:"warehouse_id"`
}
