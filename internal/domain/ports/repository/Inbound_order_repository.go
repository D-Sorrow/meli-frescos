package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type InboundOrderRepository interface {
	CreateInboundOrder(inboundOrder *models.InboundOrder) error
}
