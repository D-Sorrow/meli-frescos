package entities

type PurchaseOrderEntity struct {
	ID            int64
	OrderNumber   string
	OrderDate     string
	TrackingCode  string
	BuyerID       int64
	CarrierID     int64
	OrderStatusID int64
	WarehouseID   int64
}

func (p *PurchaseOrderEntity) GetByIdQuery(id int) (query string, args []interface{}) {
	query = `SELECT id,
		order_number,
		order_date,
		tracking_code,
		buyer_id,
		carrier_id,
		order_status_id,
		wareHouse_id
	FROM purchase_orders
	WHERE id =?
	LIMIT 1;`
	args = []interface{}{id}

	return
}

func (p *PurchaseOrderEntity) GetCreateQuery() (query string, args []interface{}) {
	query = `INSERT INTO purchase_orders (
		order_number,
		order_date,
		tracking_code,
		buyer_id,
		carrier_id,
		order_status_id,
		wareHouse_id
	) VALUES (?,?,?,?,?,?,?);`
	args = []interface{}{
		p.OrderNumber,
		p.OrderDate,
		p.TrackingCode,
		p.BuyerID,
		p.CarrierID,
		p.OrderStatusID,
		p.WarehouseID,
	}

	return
}
