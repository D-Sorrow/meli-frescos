package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/product/models"
)

type SellerRepository interface {
	GetSellers() (map[int]models.Seller, error)
}
