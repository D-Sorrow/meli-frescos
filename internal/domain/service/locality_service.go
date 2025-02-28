package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type LocalityService struct {
	repo repository.LocalityRepository
}

func NewLocalityService(repo repository.LocalityRepository) *LocalityService {
	return &LocalityService{repo: repo}
}

func (repo *LocalityService) CreateLocality(locality models.Locality) (models.Locality, error) {
	locality, err := repo.repo.CreateLocality(locality)
	if err != nil {
		return models.Locality{}, error_management.HandleErrorLocalitiesService(err)
	}
	return locality, nil
}

func (repo *LocalityService) GetSellersByLocality(localityId int) (models.LocalitySellers, error) {
	localitySellers, err := repo.repo.GetSellersByLocality(localityId)
	if err != nil {
		return models.LocalitySellers{}, error_management.HandleErrorLocalitiesService(err)
	}

	return localitySellers, nil
}

func (ls *LocalityService) GetCarriersByAllLocalities() ([]models.LocalityCarriers, error) {
	carriersByLocalities, err := ls.repo.GetCarriersByAllLocalities()
	if err != nil {
		return nil, error_management.HandleErrorLocalitiesService(err)
	}

	return carriersByLocalities, nil
}

func (ls *LocalityService) GetCarriersByLocality(id int) (models.LocalityCarriers, error) {
	localityCarriers, err := ls.repo.GetCarriersByLocality(id)
	if err != nil {
		return models.LocalityCarriers{}, error_management.HandleErrorLocalitiesService(err)
	}
	return localityCarriers, nil
}
