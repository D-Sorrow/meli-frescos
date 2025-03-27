package service

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRecordService interface {
	SaveProductRecord(productRecord models.ProductRecord) (models.ProductRecord, error)
	GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error)
}

var (
	ErrServiceProductRecordNotFound      = errors.New("product record not found")
	ErrServiceProductRecordBusinessRules = errors.New("product record business rules error")
	ErrServiceProductRecordUnknown       = errors.New("err unknown product record")
)
