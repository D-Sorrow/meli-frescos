package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type ProductRecordRepository interface {
	SaveProductRecord(productRecord models.ProductRecord) (models.ProductRecord, error)
	GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error)
}

var (
	ErrRepositoryProductRecordNotFound = errors.New("product record not found")
)
