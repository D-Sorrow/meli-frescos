package models

type Employee struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
	WarehouseId  int
}

type EmployeePatchRequest struct {
	CardNumberId *string
	FirstName    *string
	LastName     *string
	WarehouseId  *int
}

type EmployeeReportInboundOrders struct {
	Id                int
	CardNumberId      string
	FirstName         string
	LastName          string
	WarehouseId       int
	InboundOrderCount int
}
