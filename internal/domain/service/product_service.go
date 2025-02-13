package service

import (
	er "github.com/D-Sorrow/meli-frescos/internal/domain/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
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
	product, errGet := p.repo.GetProductByID(id)
	if errGet != nil {
		return models.Product{}, error_management.ErrProductNotFound
	}
	return product, nil
}

func (p ProductService) SaveProduct(productSave models.Product) error {
	errSave := p.repo.SaveProduct(productSave)
	if errSave != nil {
		return er.ProductError{Code: er.CodeNotFound, Message: er.ProductIsAlreadyExist}
	}
	return nil
}
func (p ProductService) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	return p.repo.UpdateProduct(id, attributes)
}
func (p ProductService) DeleteProduct(id int) error {
	errDelete := p.repo.DeleteProduct(id)
	if errDelete != nil {
		return er.ProductError{Code: er.CodeNotFound, Message: er.ProductNotFound}
	}
	return nil
}
