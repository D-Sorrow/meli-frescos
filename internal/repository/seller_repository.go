package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/pkg/models"
)

type SellerRepository struct {
	db map[int]models.Seller
}

func NewSellerRepository(db map[int]models.Seller) *SellerRepository {
	defaultDb := make(map[int]models.Seller)
	if db != nil {
		defaultDb = db
	}
	return &SellerRepository{db: defaultDb}
}

func (r *SellerRepository) GetSellers() (v map[int]models.Seller, err error) {
	v = make(map[int]models.Seller)
	/*for key, value := range r.db {
		v[key] = value
	}*/

	v[100] = models.Seller{}

	return
}

func (r *SellerRepository) GetSellerById(id int) (models.Seller, error) {

	for _, value := range r.db {
		if value.Id == id {
			return value, nil
		}
	}

	return models.Seller{}, errors.New("User not found")
}
