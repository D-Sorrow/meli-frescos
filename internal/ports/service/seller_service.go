package service

import "github.com/D-Sorrow/meli-frescos/pkg/models"

type SellerService interface {
	GetSellers() (map[int]models.Seller, error)
	GetSellerById(id int) (models.Seller, error)
}
