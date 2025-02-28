package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRecordRepository interface {
	SaveProductRecord(productRecord models.ProductRecord) error
	GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error)
}

var (
	ErrCodeSaveErr = errors.New("001")
	ErrCodeGetErr  = errors.New("003")
)
