package dto

import "errors"

type ProductRecordDto struct {
	ProductId      *int     `json:"product_id"`
	SalePrice      *float64 `json:"sale_price"`
	PurchasePrice  *float64 `json:"purchase_price"`
	LastUpdateTime *string  `json:"last_update_date"`
}

const (
	ErrPurchasePrice  = "purchase price must be greater than zero"
	ErrSalePrice      = "sale price must be greater than zero"
	ErrProductId      = "productId is not valid"
	ErrLastUpdateTime = "lastUpdateTime is not valid"
)

func (pr ProductRecordDto) Validation() error {
	if pr.PurchasePrice == nil {
		return errors.New(ErrPurchasePrice)
	} else if pr.SalePrice == nil {
		return errors.New(ErrSalePrice)
	} else if pr.LastUpdateTime == nil {
		return errors.New(ErrLastUpdateTime)
	} else if pr.ProductId == nil {
		return errors.New(ErrProductId)
	}
	return nil
}
