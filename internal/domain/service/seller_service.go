package service

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	repository_errors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type SellerService struct {
	repository repository.SellerRepository
}

func NewSellerService(repository repository.SellerRepository) *SellerService {
	return &SellerService{repository: repository}
}

func (repo *SellerService) GetSellers() (map[int]models.Seller, error) {
	_, err := repo.repository.GetSellers()
	if errors.Is(err, repository_errors.ErrNotFound) {
		return nil, service_errors.ErrNotFound
	}
	return repo.repository.GetSellers()
}

func (repo *SellerService) GetSellerById(id int) (models.Seller, error) {
	seller, err := repo.repository.GetSellerById(id)
	if errors.Is(err, repository_errors.ErrNotFound) {
		return models.Seller{}, service_errors.ErrNotFound
	}
	return seller, nil
}

func (repo *SellerService) CreateSeller(seller models.Seller) (models.Seller, error) {
	seller, err := repo.repository.CreateSeller(seller)
	if errors.Is(err, repository_errors.ErrAlreadyExists) {
		return models.Seller{}, service_errors.ErrAlreadyExists
	}
	return seller, nil
}

func (repo *SellerService) UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error) {
	value, err := repo.repository.UpdateSeller(id, seller)
	if err != nil {
		return models.Seller{}, service_errors.ErrNotFound
	}
	return value, nil
}

func (repo *SellerService) DeleteSeller(id int) error {
	err := repo.repository.DeleteSeller(id)
	if err != nil {
		return service_errors.ErrNotFound
	}
	return nil
}
