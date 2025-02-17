package models

type InboundOrder struct {
	Id             int
	OrderDate      string
	OrderNumber    string
	EmployeeId     int
	ProductBatchId int
	WarehouseId    int
}
