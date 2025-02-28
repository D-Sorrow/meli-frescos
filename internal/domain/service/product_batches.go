package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	er "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
	"github.com/google/uuid"
)

type productBatchesService struct {
	repo repository.ProductBatchesRepository
}

func NewProductBatches(repo repository.ProductBatchesRepository) *productBatchesService {
	return &productBatchesService{repo: repo}
}

func (s *productBatchesService) AddProductBatches(productBatches models.ProductBatches) (models.ProductBatches, error) {
	productBatches, err := s.repo.AddProductBatches(productBatches)
	if errors.Is(err, er.ErrProductBatchesAlredyExists) {
		return models.ProductBatches{}, er.ErrProductBatchesAlredyExists
	}

	return productBatches, nil
}

func (p *productBatchesService) GetById(id int) (productBatches models.ProductBatches3, err error) {
	productBatchesEntity, err := p.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductBatchNotFoundWithID) {
			err = service.ProductBatchOrderDoesNotExist
		}

		return
	}

	productBatches = *mappers.ProductBatchesEntityToProductBatches(&productBatchesEntity)

	return
}

func (p *productBatchesService) Create(product models.ProductBatches2AttributesFks) (newProductBatch models.ProductBatches3, err error) {
	newUUID := uuid.New().String()

	product.ProductBatches2Attributes.BatchNumber = newUUID

	purchaseOrderEntity := mappers.ProductBatches2AttributesFksToProductBatchesEntity(&product)

	newProductBatchEntity, err := p.repo.Create(*purchaseOrderEntity)
	if err != nil {
		if errors.Is(err, repository.ErrForeignKeysNotValid) {
			err = service.ForeignKeysNotValid
		}
		return
	}

	newProductBatch = *mappers.ProductBatchesEntityToProductBatches(&newProductBatchEntity)

	return
}
