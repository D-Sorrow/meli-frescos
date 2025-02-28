package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type SellerService struct {
	repository repository.SellerRepository
}

func NewSellerService(repository repository.SellerRepository) *SellerService {
	return &SellerService{repository: repository}
}

func (repo *SellerService) GetSellers() (map[int]models.Seller, error) {
	_, err := repo.repository.GetSellers()
	if err != nil {
		return map[int]models.Seller{}, service_errors.HandleErrorSellerService(err)
	}
	return repo.repository.GetSellers()
}

func (repo *SellerService) GetSellerById(id int) (models.Seller, error) {
	seller, err := repo.repository.GetSellerById(id)
	if err != nil {
		return models.Seller{}, service_errors.HandleErrorSellerService(err)
	}
	return seller, nil
}

func (repo *SellerService) CreateSeller(seller models.Seller) (models.Seller, error) {
	seller, err := repo.repository.CreateSeller(seller)
	if err != nil {
		return models.Seller{}, service_errors.HandleErrorSellerService(err)
	}
	return seller, nil
}

func (repo *SellerService) UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error) {
	value, err := repo.repository.UpdateSeller(id, seller)
	if err != nil {
		return models.Seller{}, service_errors.HandleErrorSellerService(err)
	}
	return value, nil
}

func (repo *SellerService) DeleteSeller(id int) error {
	err := repo.repository.DeleteSeller(id)
	if err != nil {
		return service_errors.HandleErrorSellerService(err)
	}
	return nil
}
