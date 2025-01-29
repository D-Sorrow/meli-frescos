package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type SellerService interface {
	GetSellers() (map[int]models.Seller, error)
}
