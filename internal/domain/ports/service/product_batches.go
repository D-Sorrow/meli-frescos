package service

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type ProductBatchesRepository interface {
	AddProductBatches(productBatches models.ProductBatches) (models.ProductBatches, error)
	//GetProductsBatchByAllSections() ([]models.SectionsProductsBatchs, error)
	//GetProductsBatchBySections(id int) (models.SectionsProductsBatchs, error)
	GetById(id int) (models.ProductBatches3, error)
	Create(purchaseOrder models.ProductBatches2AttributesFks) (models.ProductBatches3, error)
}

var (
	ErrForeignKeysNotValidProductBatches = errors.New(
		"the foreign keys of the requested ProductBatch order are not valid",
	)
	ErrProductBatchOrderDoesNotExist = errors.New(
		"the requested ProductBatch order does not exist in the database for the ID: %d",
	)
)
