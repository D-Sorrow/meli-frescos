package dto

type InboundOrderRequestDTO struct {
	OrderDate      string `json:"order_date" validate:"required"`
	OrderNumber    string `json:"order_number" validate:"required"`
	EmployeeId     int    `json:"employee_id" validate:"required"`
	ProductBatchId int    `json:"product_batch_id" validate:"required"`
	WarehouseId    int    `json:"warehouse_id" validate:"required"`
}

type InboundOrderResponseDTO struct {
	Id             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}
