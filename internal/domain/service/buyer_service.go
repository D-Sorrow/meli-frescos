package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	repo_errors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type BuyerService struct {
	repo repository.BuyerRepository
}

func NewBuyerService(repository repository.BuyerRepository) *BuyerService {
	return &BuyerService{repo: repository}
}

func (b *BuyerService) GetAll() (buyers map[int]models.Buyer, err error) {
	buyers, err = b.repo.GetAll()
	if err != nil {
		if errors.Is(err, repo_errors.ErrNoRegisteredBuyersYet) {
			err = service_errors.NoRegisteredBuyersYet
		}
		return
	}

	return
}

func (b *BuyerService) GetById(id int) (buyer models.Buyer, err error) {
	buyer, err = b.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repo_errors.ErrBuyerNotFoundWithID) {
			err = service_errors.BuyerDoesNotExist
		}
		return
	}

	return
}

func (b *BuyerService) Create(buyer models.BuyerAttributes) (newBuyer models.Buyer, err error) {
	newBuyer, err = b.repo.Create(buyer)
	if err != nil {
		if errors.Is(err, repo_errors.ErrDuplicateCardNumberID) {
			err = service_errors.BuyerAlreadyExists
		}
		return
	}

	return
}

func (b *BuyerService) Patch(id int, buyer models.BuyerPatchAttributes) (updatedBuyer models.Buyer, err error) {
	updatedBuyer, err = b.repo.Patch(id, buyer)
	if err != nil {
		if errors.Is(err, repo_errors.ErrDuplicateCardNumberID) {
			err = service_errors.BuyerAlreadyExists
		} else if errors.Is(err, repo_errors.ErrBuyerNotFoundWithID) {
			err = service_errors.BuyerDoesNotExist
		}
		return
	}

	return
}

func (b *BuyerService) Delete(id int) (err error) {
	err = b.repo.Delete(id)
	if err != nil {
		if errors.Is(err, repo_errors.ErrBuyerNotFoundWithID) {
			err = service_errors.BuyerDoesNotExist
		}
		return
	}

	return
}
