package models

import (
	"errors"
	"time"
)

type ProductRecord struct {
	ProductId      int
	SalePrice      float64
	PurchasePrice  float64
	LastUpdateTime string
}

var (
	ErrProductId      = errors.New("ProductId is invalid")
	ErrSalePrice      = errors.New("SalePrice is invalid")
	ErrPurchasePrice  = errors.New("PurchasePrice is invalid")
	ErrLastUpdateTime = errors.New("LastUpdateTime is invalid")
	DateFormat        = "2006-01-02"
)

func (pr ProductRecord) ValidateProductRecord() error {
	if pr.ProductId <= 0 {
		return ErrProductId
	} else if pr.SalePrice <= 0 {
		return ErrSalePrice
	} else if pr.PurchasePrice <= 0 {
		return ErrPurchasePrice
	}
	date, _ := time.Parse(DateFormat, pr.LastUpdateTime)
	if date.Before(time.Now()) {
		return ErrLastUpdateTime
	}
	return nil
}
