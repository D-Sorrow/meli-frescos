package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type CarrierServiceInterface interface {
	GetAllCarriers() ([]models.Carrier, error)
	CreateCarrier(carrier models.Carrier) (models.Carrier, error)
}

var (
	ErrCarrierServiceDefault          = errors.New("internal server error")
	ErrCarrierCidDuplicate            = errors.New("carrier cid already exists")
	ErrCarrierNotFound                = errors.New("carrier id not found")
	ErrCarrierLocalityIdNotFound      = errors.New("locality id not found")
	ErrCarrierFKConstraintFail        = errors.New("foreign key constraint fails")
	ErrCarrierGetUpdatedOrCreatedItem = errors.New("the carrier was created or updated but could not be displayed")
)
