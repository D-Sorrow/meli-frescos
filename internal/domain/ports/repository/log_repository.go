package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

type LogRepository interface {
	Register(log entities.LogEntity) (entities.LogEntity, error)
}

var (
	ErrLogUnexpectedError = errors.New("ERR_REPO_LG_UNEXPECTED_ERROR")
)
