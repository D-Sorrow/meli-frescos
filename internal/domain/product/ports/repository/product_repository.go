package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/product/models"
)

type ProductRepository interface {
	GetProducts() map[int]models.Product
	GetProductByID(id int) (models.Product, bool)
	SaveProduct(product models.Product) bool
	UpdateProduct(id int, attributes map[string]any) (models.Product, error)
	DeleteProduct(id int) bool
}
