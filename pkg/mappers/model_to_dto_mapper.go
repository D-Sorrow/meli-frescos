package mappers

import (
	"github.com/D-Sorrow/meli-frescos/pkg/dto"
	"github.com/D-Sorrow/meli-frescos/pkg/models"
)

func EmployeeModelToDTO(employeeModel models.Employee) *dto.EmployeeDTO {
	return &dto.EmployeeDTO{
		Id:           employeeModel.Id,
		CardNumberId: employeeModel.CardNumberId,
		FirstName:    employeeModel.FirstName,
		LastName:     employeeModel.LastName,
		WarehouseId:  employeeModel.WarehouseId,
	}
}
