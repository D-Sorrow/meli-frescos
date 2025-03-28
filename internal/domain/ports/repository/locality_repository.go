package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type LocalityRepository interface {
	CreateLocality(locality models.Locality) (models.Locality, error)
	GetSellersByLocality(localityId int) (models.LocalitySellers, error)
	GetCarriersByAllLocalities() ([]models.LocalityCarriers, error)
	GetCarriersByLocality(id int) (models.LocalityCarriers, error)
}

var (
	ErrLocalityNotFound      = errors.New("not found error")
	ErrLocalityAlreadyExists = errors.New("locality already exists")
	//ErrLocalityCannotBeCreated = errors.New("locality cant be created")
	ErrGetAllLocalities          = errors.New("could not get all localities")
	ErrProvinceNotFound          = errors.New("not found error")
	ErrLocalityRepositoryGeneric = errors.New("internal server error")
)
