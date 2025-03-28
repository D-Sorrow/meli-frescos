package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
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
