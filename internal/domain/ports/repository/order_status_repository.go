package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

type OrderStatusRepository interface {
	GetAll() ([]entities.OrderStatusEntity, error)
}

var (
	ErrOrderStatusNoRegisteredOrderStatusesYet = errors.New("ERR_REPO_OS_NO_REG_OS_YET")
	ErrOrderStatusUnexpectedError              = errors.New("ERR_REPO_OS_UNEXPECTED_ERROR")
)
