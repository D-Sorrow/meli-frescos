package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

func OrderStatusAttributesToOrderStatusEntity(v *models.OrderStatusAttributes) *entities.OrderStatusEntity {
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
