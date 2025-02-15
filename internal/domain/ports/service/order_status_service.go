package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type OrderStatusService interface {
	GetAll() ([]models.OrderStatus, error)
}

var (
	NoRegisteredOrderStatusesYet = errors.New("No registered order statuses yet")
)
