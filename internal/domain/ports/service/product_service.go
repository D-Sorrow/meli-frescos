package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductService interface {
	GetProducts() (map[int]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(id int, attributes map[string]any) (models.Product, error)
	DeleteProduct(id int) error
}
