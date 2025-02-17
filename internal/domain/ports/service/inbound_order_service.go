package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type InboundOrderService interface {
	CreateInboundOrder(inboundOrder *models.InboundOrder) error
}
