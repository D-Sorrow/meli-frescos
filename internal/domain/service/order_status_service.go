package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/mappers"
)

type OrderStatusService struct {
	repo repository.OrderStatusRepository
}

func NewOrderStatusService(repository repository.OrderStatusRepository) *OrderStatusService {
	return &OrderStatusService{repo: repository}
}

func (b *OrderStatusService) GetAll() (buyers []models.OrderStatus, err error) {
	buyers = make([]models.OrderStatus, 0)

	buyerEntities, err := b.repo.GetAll()
	if err != nil {
		err = error_management.HandleServiceError(
			error_management.HandleOrderStatusServiceError(err),
			err,
		)
		return
	}

	for _, buyerEntity := range buyerEntities {
		buyers = append(buyers, *mappers.OrderStatusEntityToOrderStatus(&buyerEntity))
	}

	return
}
