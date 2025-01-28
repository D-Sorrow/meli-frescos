package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/ports/repository"
	"github.com/D-Sorrow/meli-frescos/pkg/models"
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

	}
	return repo.repository.GetSellers()
}

func (repo *SellerService) GetSellerById(id int) (models.Seller, error) {
	seller, err := repo.repository.GetSellerById(id)
	if err != nil {
		return models.Seller{}, errors.New("user doesnt exist")
	}
	return seller, nil
}
