package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRepository struct {
	productMap map[int]models.Product
}

func NewProductRepository(productMap map[int]models.Product) *ProductRepository {
	return &ProductRepository{productMap: productMap}
}

func (p ProductRepository) GetProducts() map[int]models.Product {
	return p.productMap
}
