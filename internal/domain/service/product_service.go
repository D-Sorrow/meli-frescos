package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	er "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"net/http"
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
		return models.Product{}, er.Error{Code: http.StatusNotFound, Message: er.ProductNotFound}
	}

	return product, nil
}

func (p ProductService) SaveProduct(productSave models.Product) error {
	isCreated := p.repo.SaveProduct(productSave)
	if isCreated == false {
		return er.Error{Code: http.StatusNotFound, Message: er.ProductIsAlreadyExist}
	}
	return nil
}
func (p ProductService) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	product, errUpdate := p.repo.UpdateProduct(id, attributes)
	if errUpdate != nil {
		return models.Product{}, errUpdate
	}
	return product, nil
}
func (p ProductService) DeleteProduct(id int) error {
	isDeleted := p.repo.DeleteProduct(id)
	if isDeleted == false {
		return er.Error{Code: http.StatusNotFound, Message: er.ProductNotFound}
	}
	return nil
}
