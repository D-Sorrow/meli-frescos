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

func (p ProductRepository) GetProductByID(id int) (models.Product, bool) {
	product, ok := p.productMap[id]
	if !ok {
		return models.Product{}, false
	}
	return product, true
}
