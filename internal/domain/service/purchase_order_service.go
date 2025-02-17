package service

import (
	"errors"
	"time"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
	"github.com/google/uuid"
)

type PurchaseOrderService struct {
	repo repository.PurchaseOrderRepository
}

func NewPurchaseOrderService(repository repository.PurchaseOrderRepository) *PurchaseOrderService {
	return &PurchaseOrderService{repo: repository}
}

func (p *PurchaseOrderService) GetById(id int) (purchaseOrder models.PurchaseOrder, err error) {
	purchaseOrderEntity, err := p.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrPurchaseOrderNotFoundWithID) {
			err = service.PurchaseOrderDoesNotExist
		}

		return
	}

	purchaseOrder = *mappers.PurchaseOrderEntityToPurchaseOrder(&purchaseOrderEntity)

	return
}

func (p *PurchaseOrderService) Create(buyer models.PurchaseOrderAttributesFKs) (newPurchaseOrder models.PurchaseOrder, err error) {
	newUUID := uuid.New().String()
	utcNow := time.Now().UTC()
	buyer.PurchaseOrderAttributes.TrackingCode = newUUID
	buyer.PurchaseOrderAttributes.OrderDate = utcNow.Format("2006-01-02 15:04:05")

	purchaseOrderEntity := mappers.PurchaseOrderAttributesFKsToPurchaseOrderEntity(&buyer)

	newPurchaseOrderEntity, err := p.repo.Create(*purchaseOrderEntity)
	if err != nil {
		if errors.Is(err, repository.ErrForeignKeysNotValid) {
			err = service.ForeignKeysNotValid
		}
		return
	}

	newPurchaseOrder = *mappers.PurchaseOrderEntityToPurchaseOrder(&newPurchaseOrderEntity)

	return
}
