package models

type Employee struct {
	Id           int
	CardNumberId int
	FirstName    string
	LastName     string
	WarehouseId  int
}

type EmployeePatchRequest struct {
	CardNumberId *int
	FirstName    *string
	LastName     *string
	WarehouseId  *int
}
