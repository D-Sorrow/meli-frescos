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
	//TODO
	return models.LocalitySellers{}, nil
}
