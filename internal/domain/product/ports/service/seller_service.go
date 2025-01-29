package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/product/models"
)

type SellerService interface {
	GetSellers() (map[int]models.Seller, error)
}
