package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type BuyerRepository struct {
	db map[int]models.Buyer
}

func NewBuyerRepository(db map[int]models.Buyer) *BuyerRepository {
	defaultDb := make(map[int]models.Buyer)
	if db != nil {
		defaultDb = db
	}

	return &BuyerRepository{db: defaultDb}
}

func (b *BuyerRepository) GetAll() (buyers map[int]models.Buyer, err error) {
	buyers = make(map[int]models.Buyer)

	for key, value := range b.db {
		buyers[key] = value
	}

	if len(buyers) == 0 {
		err = error_management.ErrNoRegisteredBuyersYet
	}

	return
}

func (b *BuyerRepository) GetById(id int) (buyer models.Buyer, err error) {
	buyer, ok := b.db[id]
	if !ok {
		err = error_management.ErrBuyerNotFoundWithID
		return
	}

	return
}

func (b *BuyerRepository) Create(buyer models.BuyerAttributes) (newBuyer models.Buyer, err error) {
	maxID := 0
	for key, value := range b.db {
		if value.BuyerAttributes.CardNumberID == buyer.CardNumberID {
			err = error_management.ErrDuplicateCardNumberID
			return
		}

		if key > maxID {
			maxID = key
		}
	}

	newBuyer = models.Buyer{
		ID:              maxID + 1,
		BuyerAttributes: buyer,
	}
	b.db[newBuyer.ID] = newBuyer

	return
}

func (b *BuyerRepository) Patch(id int, buyerToPatch models.BuyerPatchAttributes) (updatedBuyer models.Buyer, err error) {
	updatedBuyer, ok := b.db[id]

	if !ok {
		err = error_management.ErrBuyerNotFoundWithID
		return
	}

	if buyerToPatch.CardNumberID != nil {
		for _, value := range b.db {
			if value.BuyerAttributes.CardNumberID == *buyerToPatch.CardNumberID {
				err = error_management.ErrDuplicateCardNumberID
				return
			}
		}
		updatedBuyer.CardNumberID = *buyerToPatch.CardNumberID
	}

	if buyerToPatch.FirstName != nil {
		updatedBuyer.BuyerAttributes.FirstName = *buyerToPatch.FirstName
	}

	if buyerToPatch.LastName != nil {
		updatedBuyer.BuyerAttributes.LastName = *buyerToPatch.LastName
	}

	b.db[id] = updatedBuyer

	return
}

func (b *BuyerRepository) Delete(id int) (err error) {
	_, ok := b.db[id]

	if !ok {
		err = error_management.ErrBuyerNotFoundWithID
		return
	}

	delete(b.db, id)
	return
}
