package repository

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRepository interface {
	GetProducts() (map[int]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(id int, attributes map[string]any) (models.Product, error)
	DeleteProduct(id int) error
}

var (
	CodeQueryConsult        = errors.New("001")
	CodeGetProduct          = errors.New("002")
	CodeNoRowsAffected      = errors.New("003")
	CodeDelete              = errors.New("005")
	CodeDeleteIsNotPossible = errors.New("006")
)
