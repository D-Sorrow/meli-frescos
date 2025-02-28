package repository

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRecordRepository interface {
	SaveProductRecord(productRecord models.ProductRecord) (models.ProductRecord, error)
	GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error)
}

var (
	ErrRepositoryProductRecordNotFound = errors.New("product record not found")
)
