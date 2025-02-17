package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

type ProductBatchesRepository interface {
	AddProductBatches(productBatches models.ProductBatches) (models.ProductBatches, error)
	//GetProductsBatchByAllSections() ([]models.SectionsProductsBatchs, error)
	//GetProductsBatchBySections(id int) (models.SectionsProductsBatchs, error)
	GetById(id int) (entities.ProductBatchesEntity, error)
	Create(purchaseOrder entities.ProductBatchesEntity) (entities.ProductBatchesEntity, error)
}

var (
	ErrForeignKeysNotValid        = errors.New("The foreign keys of the requested Productbatches order are not valid")
	ErrProductBatchNotFoundWithID = errors.New("The requested  Productbatches order does not exist with ID provided")
)
