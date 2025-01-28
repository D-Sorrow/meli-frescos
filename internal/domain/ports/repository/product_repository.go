package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRepository interface {
	GetProducts() map[int]models.Product
}
