package dto

type EmployeeDTO struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

type EmployeeRequestDTO struct {
	CardNumberId string `json:"card_number_id" validate:"required"`
	FirstName    string `json:"first_name" validate:"required,min=1,max=100"`
	LastName     string `json:"last_name" validate:"required,min=1,max=100"`
	WarehouseId  int    `json:"warehouse_id" validate:"required"`
}

type EmployeePatchRequestDTO struct {
	CardNumberId *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseId  *int    `json:"warehouse_id"`
}
