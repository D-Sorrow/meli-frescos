package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
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
		if err.Error() == "ONAE_DB" {
			return errors.New("ONAE_SV")
		}
		if err.Error() == "EIDNF_DB" {
			return errors.New("EIDNF_SV")
		}
		return err
	}
	return nil
}
