package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

func OrderStatusAttributesToOrderStatusEntity(
	v *models.OrderStatusAttributes,
) *entities.OrderStatusEntity {
	return &entities.OrderStatusEntity{
		Description: v.Description,
	}
}

func OrderStatusEntityToOrderStatus(v *entities.OrderStatusEntity) *models.OrderStatus {
	return &models.OrderStatus{
		ID: v.ID,
		OrderStatusAttributes: models.OrderStatusAttributes{
			Description: v.Description,
		},
	}
}
