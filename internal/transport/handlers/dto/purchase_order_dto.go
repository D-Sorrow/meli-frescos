package dto

type PurchaseOrderDTO struct {
	ID            int64  `json:"id"`
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerID       int64  `json:"buyer_id"`
	CarrierID     int64  `json:"carrier_id"`
	OrderStatusID int64  `json:"order_status_id"`
	WarehouseID   int64  `json:"warehouse_id"`
}

type PurchaseOrderCreateDTO struct {
	BuyerID       int64 `json:"buyer_id"`
	CarrierID     int64 `json:"carrier_id"`
	OrderStatusID int64 `json:"order_status_id"`
	WarehouseID   int64 `json:"warehouse_id"`
}
