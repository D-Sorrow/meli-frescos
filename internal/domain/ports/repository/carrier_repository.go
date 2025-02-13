package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type CarrierRepositoryInterface interface {
	CreateCarrier(carrier models.Carrier) (models.Carrier, error)
}
