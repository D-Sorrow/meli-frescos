package repository

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
)

type SellerRepository interface {
	GetSellers() (map[int]models.Seller, error)
	GetSellerById(id int) (models.Seller, error)
	CreateSeller(seller models.Seller) (models.Seller, error)
	UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error)
	DeleteSeller(id int) error
}

var (
	ErrSellerAlreadyExists     = errors.New("seller already exists")
	ErrSellerNotFound          = errors.New("seller not found")
	ErrSellerRepositoryGeneric = errors.New("internal server error")
)
