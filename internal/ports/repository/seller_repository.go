package repository

import "github.com/D-Sorrow/meli-frescos/pkg/models"

type SellerRepository interface {
	GetSellers() (map[int]models.Seller, error)
}
