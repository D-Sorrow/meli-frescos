package service

import (
	er "github.com/D-Sorrow/meli-frescos/internal/domain/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (p ProductService) GetProducts() map[int]models.Product {
	return p.repo.GetProducts()
}

func (p ProductService) GetProductByID(id int) (models.Product, error) {
	product, isFalse := p.repo.GetProductByID(id)

	if isFalse == false {
		return models.Product{}, er.ProductError{Code: er.CodeNotFound, Message: er.ProductNotFound}
	}

	return product, nil
}

func (p ProductService) SaveProduct(productSave models.Product) error {
	isCreated := p.repo.SaveProduct(productSave)
	if isCreated == false {
		return er.ProductError{Code: er.CodeNotFound, Message: er.ProductIsAlreadyExist}
	}
	return nil
}
func (p ProductService) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	return p.repo.UpdateProduct(id, attributes)
}
func (p ProductService) DeleteProduct(id int) error {
	isDeleted := p.repo.DeleteProduct(id)
	if isDeleted == false {
		return er.ProductError{Code: er.CodeNotFound, Message: er.ProductNotFound}
	}
	return nil
}
