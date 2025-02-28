package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type LocalityService interface {
	CreateLocality(locality models.Locality) (models.Locality, error)
	GetSellersByLocality(localityId int) (models.LocalitySellers, error)
	GetCarriersByAllLocalities() ([]models.LocalityCarriers, error)
	GetCarriersByLocality(id int) (models.LocalityCarriers, error)
}
