package service

import (
	"time"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
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
		err = error_management.HandleServiceError(
			error_management.HandlePurchaseOrderServiceError(err),
			err,
		)
		return
	}

	purchaseOrder = *mappers.PurchaseOrderEntityToPurchaseOrder(&purchaseOrderEntity)

	return
}

func (p *PurchaseOrderService) Create(
	buyer models.PurchaseOrderAttributesFKs,
	utcNow time.Time,
	newUUID string,
) (newPurchaseOrder models.PurchaseOrder, err error) {
	buyer.PurchaseOrderAttributes.TrackingCode = newUUID
	buyer.PurchaseOrderAttributes.OrderDate = utcNow.Format("2006-01-02 15:04:05")

	purchaseOrderEntity := mappers.PurchaseOrderAttributesFKsToPurchaseOrderEntity(&buyer)

	newPurchaseOrderEntity, err := p.repo.Create(*purchaseOrderEntity)
	if err != nil {
		err = error_management.HandleServiceError(
			error_management.HandlePurchaseOrderServiceError(err),
			err,
		)
		return
	}

	newPurchaseOrder = *mappers.PurchaseOrderEntityToPurchaseOrder(&newPurchaseOrderEntity)

	return
}
