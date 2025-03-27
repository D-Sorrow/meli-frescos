package models

type OrderStatusAttributes struct {
	Description string
}

type OrderStatus struct {
	ID int64
	OrderStatusAttributes
}
