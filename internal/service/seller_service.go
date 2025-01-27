package service

import (
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
