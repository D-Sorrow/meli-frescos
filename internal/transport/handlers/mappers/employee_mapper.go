package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
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

func EmployeeDTOToModel(employeeDTO dto.EmployeeRequestDTO) *models.Employee {
	return &models.Employee{
		CardNumberId: employeeDTO.CardNumberId,
		FirstName:    employeeDTO.FirstName,
		LastName:     employeeDTO.LastName,
		WarehouseId:  employeeDTO.WarehouseId,
	}
}
