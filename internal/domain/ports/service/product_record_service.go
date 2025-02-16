package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRecordService interface {
	SaveProductRecord(productRecord models.ProductRecord) error
	GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error)
}
