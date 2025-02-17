package service

import "github.com/D-Sorrow/meli-frescos/pkg/models"

type SellerService interface {
	GetSellers() (map[int]models.Seller, error)
}
