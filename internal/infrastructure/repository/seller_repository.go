package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repository_errors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
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
	for key, value := range r.db {
		v[key] = value
	}
	if len(v) <= 0 {
		return nil, repository_errors.ErrNotFound
	}
	return
}

func (r *SellerRepository) GetSellerById(id int) (models.Seller, error) {

	for _, value := range r.db {
		if value.Id == id {
			return value, nil
		}
	}

	return models.Seller{}, repository_errors.ErrNotFound
}

func (r *SellerRepository) CreateSeller(seller models.Seller) (models.Seller, error) {

	var newId int = 1

	for _, value := range r.db {
		if value.Id >= newId {
			newId = value.Id + 1
		}
		if value.Cid == seller.Cid {
			return models.Seller{}, repository_errors.ErrAlreadyExists
		}
	}
	seller.Id = newId
	r.db[seller.Id] = seller

	return seller, nil
}

func (r *SellerRepository) UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error) {

	value, exist := r.db[id]
	if !exist {
		return models.Seller{}, repository_errors.ErrNotFound
	}

	if seller.Cid != nil {
		value.Cid = *seller.Cid
	}
	if seller.Address != nil {
		value.Address = *seller.Address
	}
	if seller.CompanyName != nil {
		value.CompanyName = *seller.CompanyName
	}
	if seller.Telephone != nil {
		value.Telephone = *seller.Telephone
	}

	r.db[id] = value

	return r.db[id], nil

}
