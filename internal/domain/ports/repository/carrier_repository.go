package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type CarrierRepositoryInterface interface {
	GetAllCarriers() ([]models.Carrier, error)
	CreateCarrier(carrier models.Carrier) (models.Carrier, error)
}
