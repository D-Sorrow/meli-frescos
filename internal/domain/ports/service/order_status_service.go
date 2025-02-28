package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type OrderStatusService interface {
	GetAll() ([]models.OrderStatus, error)
}

var (
	ErrOrderStatusNoRegisteredOrderStatusesYet = errors.New("ERR_SRV_OS_NO_REG_OS_YET")
	ErrOrderStatusUnexpectedError              = errors.New("ERR_SRV_OS_UNEXPECTED_ERROR")
)
