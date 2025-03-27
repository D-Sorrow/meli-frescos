package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type InboundOrderService struct {
	repository repository.InboundOrderRepository
}

func NewInboundOrderService(repository repository.InboundOrderRepository) *InboundOrderService {
	return &InboundOrderService{repository: repository}
}

func (service *InboundOrderService) CreateInboundOrder(inboundOrder *models.InboundOrder) error {
	err := service.repository.CreateInboundOrder(inboundOrder)
	if err != nil {
		return error_management.HandleInboundOrderServiceError(err)
	}
	return nil
}
