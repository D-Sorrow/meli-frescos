package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductService interface {
	GetProducts() map[int]models.Product
}
