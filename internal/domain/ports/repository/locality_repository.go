package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type LocalityRepository interface {
	CreateLocality(locality models.Locality) (models.Locality, error)
	GetSellersByLocality(localityId int) (models.LocalitySellers, error)
	GetCarriersByAllLocalities() ([]models.LocalityCarriers, error)
	GetCarriersByLocality(id int) (models.LocalityCarriers, error)
}
