package dto

import (
	"errors"
)

const descriptionInvalid = "description must not be empty"
const expirationInvalid = "rates must not be null"
const freezingInvalid = "freezingRate must not be null"
const dimensionsInvalid = "dimensions must not be null"
const netWeightInvalid = "net weight must be null"
const productCodeInvalid = "productCode must not be empty"
const productTypeInvalid = "productTypeId must not be null"
const sellerIdInvalid = "sellerId must not be null"

type ProductDto struct {
	Id                  *int     `json:"id"`
	Description         *string  `json:"description"`
	ExpirationRate      *int     `json:"expiration_rate"`
	FreezingRate        *int     `json:"freezing_rate"`
	Height              *float64 `json:"height"`
	Length              *float64 `json:"length"`
	NetWeight           *float64 `json:"netweight"`
	ProductCode         *string  `json:"product_code"`
	TemperatureFreezing *float64 `json:"recommended_freezing_temperature"`
	Width               *float64 `json:"width"`
	ProductTypeId       *int     `json:"product_type_id"`
	SellerId            *int     `json:"seller_id"`
}

func (p *ProductDto) Validate() error {
	if p.Description == nil {
		return errors.New(descriptionInvalid)
	}
	if p.ExpirationRate == nil {
		return errors.New(expirationInvalid)
	}
	if p.FreezingRate == nil {
		return errors.New(freezingInvalid)
	}
	if p.Height == nil || p.Length == nil || p.Width == nil {
		return errors.New(dimensionsInvalid)
	}
	if p.NetWeight == nil {
		return errors.New(netWeightInvalid)
	}
	if p.ProductCode == nil {
		return errors.New(productCodeInvalid)
	}
	if p.ProductTypeId == nil {
		return errors.New(productTypeInvalid)
	}
	if p.SellerId == nil {
		return errors.New(sellerIdInvalid)
	}
	return nil
}
