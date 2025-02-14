package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type CarrierServiceInterface interface {
	GetAllCarriers() ([]models.Carrier, error)
	CreateCarrier(carrier models.Carrier) (models.Carrier, error)
}
