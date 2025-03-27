package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type LogService interface {
	Register(log models.LogAttributes) (models.Log, error)
}

var (
	ErrLogUnexpectedError = errors.New("ERR_SRV_LG_UNEXPECTED_ERROR")
)
