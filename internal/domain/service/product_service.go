package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (p ProductService) GetProducts() (map[int]models.Product, error) {
	return p.repo.GetProducts()
}

func (p ProductService) GetProductByID(id int) (models.Product, error) {
	product, err := p.repo.GetProductByID(id)
	if err != nil {
		return models.Product{}, error_management.HandlerServiceProductError(err)
	}
	return product, nil
}

func (p ProductService) SaveProduct(productSave models.Product) error {
	err := productSave.ValidateProduct()
	if err != nil {
		return service.ErrServiceProductBusinessRules
	}
	err = p.repo.SaveProduct(productSave)
	if err != nil {
		return error_management.HandlerServiceProductError(err)
	}
	return nil
}
func (p ProductService) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	err := p.repo.UpdateProduct(id, attributes)
	if err != nil {
		return models.Product{}, error_management.HandlerServiceProductError(err)
	}
	product, err := p.repo.GetProductByID(id)
	if err != nil {
		return models.Product{}, error_management.HandlerServiceProductError(err)
	}
	return product, nil
}
func (p ProductService) DeleteProduct(id int) error {
	return p.repo.DeleteProduct(id)
}
