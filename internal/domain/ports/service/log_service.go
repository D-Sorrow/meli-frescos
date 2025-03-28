package service

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type LogService interface {
	Register(log models.LogAttributes) (models.Log, error)
}

var (
	ErrLogUnexpectedError = errors.New("ERR_SRV_LG_UNEXPECTED_ERROR")
)
