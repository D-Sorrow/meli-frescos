package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
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
