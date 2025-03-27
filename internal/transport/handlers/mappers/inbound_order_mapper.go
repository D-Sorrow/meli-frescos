package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func InboundOrderModelToResponseDTO(model models.InboundOrder) *dto.InboundOrderResponseDTO {
	return &dto.InboundOrderResponseDTO{
		Id:             model.Id,
		OrderDate:      model.OrderDate,
		OrderNumber:    model.OrderNumber,
		EmployeeId:     model.EmployeeId,
		ProductBatchId: model.ProductBatchId,
		WarehouseId:    model.WarehouseId,
	}
}

func InboundOrderRequestDTOToModel(dto dto.InboundOrderRequestDTO) *models.InboundOrder {
	return &models.InboundOrder{
		OrderDate:      dto.OrderDate,
		OrderNumber:    dto.OrderNumber,
		EmployeeId:     dto.EmployeeId,
		ProductBatchId: dto.ProductBatchId,
		WarehouseId:    dto.WarehouseId,
	}
}
