package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
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
		if errors.Is(err, repository.ErrNoRegisteredOrderStatusesYet) {
			err = service.NoRegisteredOrderStatusesYet
		}

		return
	}

	for _, buyerEntity := range buyerEntities {
		buyers = append(buyers, *mappers.OrderStatusEntityToOrderStatus(&buyerEntity))
	}

	return
}
