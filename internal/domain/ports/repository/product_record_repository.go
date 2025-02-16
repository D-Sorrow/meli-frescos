package repository

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRecordRepository interface {
	SaveProductRecord(productRecord models.ProductRecord) error
}

var (
	CodeSaveErr = errors.New("001")
)
