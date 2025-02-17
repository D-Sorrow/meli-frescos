package service

import (
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

func (p ProductService) GetProducts() (map[int]models.Product, error) {
	return p.repo.GetProducts()
}

func (p ProductService) GetProductByID(id int) (models.Product, error) {
	return p.repo.GetProductByID(id)
}

func (p ProductService) SaveProduct(productSave models.Product) error {
	err := productSave.ValidateProduct()
	if err != nil {
		return &error_management.ErrService{
			Code: error_management.CodeUseCaseError,
		}
	}
	return p.repo.SaveProduct(productSave)
}
func (p ProductService) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	return p.repo.UpdateProduct(id, attributes)
}
func (p ProductService) DeleteProduct(id int) error {
	return p.repo.DeleteProduct(id)
}
