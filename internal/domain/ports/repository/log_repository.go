package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

type LogRepository interface {
	Register(log entities.LogEntity) (entities.LogEntity, error)
}

var (
	ErrLogUnexpectedError = errors.New("ERR_REPO_LG_UNEXPECTED_ERROR")
)
