package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	repository_errors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type LocalityService struct {
	repo repository.LocalityRepository
}

func NewLocalityService(repo repository.LocalityRepository) *LocalityService {
	return &LocalityService{repo: repo}
}

func (repo *LocalityService) CreateLocality(locality models.Locality) (models.Locality, error) {
	locality, err := repo.repo.CreateLocality(locality)
	if errors.Is(err, repository_errors.ErrLocalityAlreadyExists) {
		return models.Locality{}, service_errors.ErrLocalityAlreadyExists
	}
	return locality, nil
}

func (repo *LocalityService) GetSellersByLocality(localityId int) (models.LocalitySellers, error) {
	localitySellers, err := repo.repo.GetSellersByLocality(localityId)
	if errors.Is(err, repository_errors.ErrLocalityNotFound) {
		return models.LocalitySellers{}, service_errors.ErrLocalityNotFound
	}
	return localitySellers, nil
}

func (ls *LocalityService) GetCarriersByAllLocalities() ([]models.LocalityCarriers, error) {
	carriersByLocalities, err := ls.repo.GetCarriersByAllLocalities()
	if errors.Is(err, repository_errors.ErrGetAllLocalities) {
		return nil, service_errors.ErrGetAllLocalities
	}

	return carriersByLocalities, nil
}

func (ls *LocalityService) GetCarriersByLocality(id int) (models.LocalityCarriers, error) {
	localityCarriers, err := ls.repo.GetCarriersByLocality(id)
	if errors.Is(err, repository_errors.ErrLocalityNotFound) {
		return models.LocalityCarriers{}, service_errors.ErrLocalityNotFound
	}
	return localityCarriers, nil
}
