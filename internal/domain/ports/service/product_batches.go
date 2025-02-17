package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductBatchesRepository interface {
	AddProductBatches(productBatches models.ProductBatches) (models.ProductBatches, error)
	//GetProductsBatchByAllSections() ([]models.SectionsProductsBatchs, error)
	//GetProductsBatchBySections(id int) (models.SectionsProductsBatchs, error)
	GetById(id int) (models.ProductBatches3, error)
	Create(purchaseOrder models.ProductBatches2AttributesFks) (models.ProductBatches3, error)
}

var (
	ForeignKeysNotValid           = errors.New("The foreign keys of the requested ProductBatch order are not valid")
	ProductBatchOrderDoesNotExist = errors.New("The requested ProductBatch order does not exist in the database for the ID: %d")
)
