package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type CarrierServiceInterface interface {
	CreateCarrier(carrier models.Carrier) (models.Carrier, error)
}
