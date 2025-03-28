package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type CarrierRepositoryInterface interface {
	GetAllCarriers() ([]models.Carrier, error)
	CreateCarrier(carrier models.Carrier) (models.Carrier, error)
}

var (
	ErrCarrierDataBase                = errors.New("database error")
	ErrCarrierCidDuplicate            = errors.New("carrier cid duplicate")
	ErrCarrierNotFound                = errors.New("carrier id not found")
	ErrCarrierLocalityId              = errors.New("locality id not found")
	ErrCarrierFKConstraintFail        = errors.New("foreign key constraint fails")
	ErrCarrierGetUpdatedOrCreatedItem = errors.New("error getting updated or created carrier")
)
