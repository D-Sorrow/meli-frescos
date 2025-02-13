package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type ProductRecordRepository interface {
	SaveProductRecord(productRecord models.ProductRecord) error
}
