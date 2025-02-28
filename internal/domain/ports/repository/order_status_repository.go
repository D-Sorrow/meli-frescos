package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

type OrderStatusRepository interface {
	GetAll() ([]entities.OrderStatusEntity, error)
}

var (
	ErrOrderStatusNoRegisteredOrderStatusesYet = errors.New("ERR_REPO_OS_NO_REG_OS_YET")
	ErrOrderStatusUnexpectedError              = errors.New("ERR_REPO_OS_UNEXPECTED_ERROR")
)
