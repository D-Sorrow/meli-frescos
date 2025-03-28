package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type ProductRepository interface {
	GetProducts() (map[int]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(id int, attributes map[string]any) error
	DeleteProduct(id int) error
}

var (
	ErrRepositoryProductNotFound      = errors.New("product not found")
	ErrRepositoryProductAlreadyExists = errors.New("product already exists")
	ErrRepositoryProductUnknown       = errors.New("product unknown")
)
