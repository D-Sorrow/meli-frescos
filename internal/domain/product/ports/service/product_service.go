package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/product/models"
)

type ProductService interface {
	GetProducts() map[int]models.Product
	GetProductByID(id int) (models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(id int, attributes map[string]any) (models.Product, error)
	DeleteProduct(id int) error
}
