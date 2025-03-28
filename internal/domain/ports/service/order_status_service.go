package service

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type OrderStatusService interface {
	GetAll() ([]models.OrderStatus, error)
}

var (
	ErrOrderStatusNoRegisteredOrderStatusesYet = errors.New("ERR_SRV_OS_NO_REG_OS_YET")
	ErrOrderStatusUnexpectedError              = errors.New("ERR_SRV_OS_UNEXPECTED_ERROR")
)
