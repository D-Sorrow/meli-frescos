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

func EmployeePatchRequestDTOToModel(employeePatchRequestDTO dto.EmployeePatchRequestDTO) *models.EmployeePatchRequest {
	return &models.EmployeePatchRequest{
		CardNumberId: employeePatchRequestDTO.CardNumberId,
		FirstName:    employeePatchRequestDTO.FirstName,
		LastName:     employeePatchRequestDTO.LastName,
		WarehouseId:  employeePatchRequestDTO.WarehouseId,
	}
}

func EmployeeReportInboundOrdersModelToDTO(model models.EmployeeReportInboundOrders) *dto.EmployeeReportInboundOrdersDTO {
	return &dto.EmployeeReportInboundOrdersDTO{
		Id:                 model.Id,
		CardNumberId:       model.CardNumberId,
		FirstName:          model.FirstName,
		LastName:           model.LastName,
		WarehouseId:        model.WarehouseId,
		InboundOrdersCount: model.InboundOrderCount,
	}
}
