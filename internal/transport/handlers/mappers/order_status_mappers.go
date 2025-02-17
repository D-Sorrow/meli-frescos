package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func OrderStatusDTOToOrderStatus(b *dto.OrderStatusDTO) *models.OrderStatus {
	return &models.OrderStatus{
		ID: b.ID,
		OrderStatusAttributes: models.OrderStatusAttributes{
			Description: b.Description,
		},
	}
}

func OrderStatusToOrderStatusDTO(v *models.OrderStatus) *dto.OrderStatusDTO {
	return &dto.OrderStatusDTO{
		ID:          v.ID,
		Description: v.OrderStatusAttributes.Description,
	}
}
