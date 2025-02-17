package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

type OrderStatusRepository interface {
	GetAll() ([]entities.OrderStatusEntity, error)
}

var (
	ErrNoRegisteredOrderStatusesYet = errors.New("No registered order statuses yet")
)
