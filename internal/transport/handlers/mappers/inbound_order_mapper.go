package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
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
