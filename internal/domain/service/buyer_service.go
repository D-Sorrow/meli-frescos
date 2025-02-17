package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
)

type BuyerService struct {
	repo repository.BuyerRepository
}

func NewBuyerService(repository repository.BuyerRepository) *BuyerService {
	return &BuyerService{repo: repository}
}

func (b *BuyerService) GetAll() (buyers []models.Buyer, err error) {
	buyers = make([]models.Buyer, 0)

	buyerEntities, err := b.repo.GetAll()
	if err != nil {
		if errors.Is(err, repository.ErrNoRegisteredBuyersYet) {
			err = service.NoRegisteredBuyersYet
		}

		return
	}

	for _, buyerEntity := range buyerEntities {
		buyers = append(buyers, *mappers.BuyerEntityToBuyer(&buyerEntity))
	}

	return
}

func (b *BuyerService) GetById(id int) (buyer models.Buyer, err error) {
	buyerEntity, err := b.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrBuyerNotFoundWithID) {
			err = service.BuyerDoesNotExist
		}

		return
	}

	buyer = *mappers.BuyerEntityToBuyer(&buyerEntity)

	return
}

func (b *BuyerService) Create(buyer models.BuyerAttributes) (newBuyer models.Buyer, err error) {
	buyerEntity := mappers.BuyerAttributesToBuyerEntity(&buyer)

	newBuyerEntity, err := b.repo.Create(*buyerEntity)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateCardNumberID) {
			err = service.BuyerAlreadyExists
		}

		return
	}

	newBuyer = *mappers.BuyerEntityToBuyer(&newBuyerEntity)

	return
}

func (b *BuyerService) Patch(id int, buyer models.BuyerAttributes) (updatedBuyer models.Buyer, err error) {
	buyerEntity := mappers.BuyerAttributesToBuyerEntity(&buyer)

	updatedBuyerEntity, err := b.repo.Patch(id, *buyerEntity)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateCardNumberID) {
			err = service.BuyerAlreadyExists
		} else if errors.Is(err, repository.ErrBuyerNotFoundWithID) {
			err = service.BuyerDoesNotExist
		}

		return
	}

	updatedBuyer = *mappers.BuyerEntityToBuyer(&updatedBuyerEntity)

	return
}

func (b *BuyerService) Delete(id int) (err error) {
	err = b.repo.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrBuyerNotFoundWithID) {
			err = service.BuyerDoesNotExist
		} else if errors.Is(err, repository.ErrCannotDeleteBuyerWithOrders) {
			err = service.CannotDeleteBuyerWithOrders
		}

		return
	}

	return
}

func (b *BuyerService) GetReportPurchaseOrders(buyerID *int) (report []models.ReportPurchaseOrders, err error) {
	report = make([]models.ReportPurchaseOrders, 0)

	reportEntities, err := b.repo.GetReportPurchaseOrders(buyerID)
	if err != nil {
		if errors.Is(err, repository.ErrBuyerNotFoundOrBuyerHasNoOrders) {
			if buyerID == nil {
				err = service.BuyerHasNoOrders
				return
			} else {
				err = service.BuyerDoesNotExist
				return
			}
		}

		return
	}

	for _, reportEntity := range reportEntities {
		report = append(report, *mappers.ReportPurchaseOrdersEntityToReportPurchaseOrders(&reportEntity))
	}

	return
}
