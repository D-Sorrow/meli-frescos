package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

type ProductBatchesRepository interface {
	AddProductBatches(productBatches models.ProductBatches) (models.ProductBatches, error)
	//GetProductsBatchByAllSections() ([]models.SectionsProductsBatchs, error)
	//GetProductsBatchBySections(id int) (models.SectionsProductsBatchs, error)
	GetById(id int) (entities.ProductBatchesEntity, error)
	Create(purchaseOrder entities.ProductBatchesEntity) (entities.ProductBatchesEntity, error)
}

var (
	ErrForeignKeysNotValidProductBatches = errors.New(
		"the foreign keys of the requested Productbatches order are not valid",
	)
	ErrProductBatchNotFoundWithID = errors.New(
		"the requested  Productbatches order does not exist with ID provided",
	)
)
