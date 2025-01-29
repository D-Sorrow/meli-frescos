package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type SellerService interface {
	GetSellers() (map[int]models.Seller, error)
	GetSellerById(id int) (models.Seller, error)
	CreateSeller(seller models.Seller) (models.Seller, error)
	UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error)
}
