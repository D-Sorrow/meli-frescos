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
	CodeSaveErr                        = errors.New("001")
	CodeGetErr                         = errors.New("003")
	ErrRepositoryProductRecordNotFound = errors.New("product record not found")
	ErrRepositoryProductRecordUnknown  = errors.New("product record unknown error")
)
