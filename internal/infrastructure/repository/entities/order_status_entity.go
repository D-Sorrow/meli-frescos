package entities

type OrderStatusEntity struct {
	ID          int
	Description string
}

func (p *OrderStatusEntity) GetAllQuery() (query string, args []interface{}) {
	query = `SELECT  id,
		description
	FROM order_status;`
	args = []interface{}{}

	return
}
