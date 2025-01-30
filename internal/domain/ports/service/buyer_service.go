package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type BuyerService interface {
	GetAll() (map[int]models.Buyer, error)
	GetById(id int) (models.Buyer, error)
	Create(buyer models.BuyerAttributes) (models.Buyer, error)
	Patch(id int, buyerToPatch models.BuyerPatchAttributes) (models.Buyer, error)
	Delete(id int) error
}
