package dto

import (
	"errors"
)

const DescriptionInvalid = "Description must not be empty."
const ExpirationInvalid = "Rates must not be null"
const FreezingInvalid = "FreezingRate must not be null"
const DimensionsInvalid = "Dimensions must not be null"
const NetWeightInvalid = "Net weight must be null"
const ProductCodeInvalid = "ProductCode must not be empty"
const ProductTypeInvalid = "ProductTypeId must not be null"
const SellerIdInvalid = "SellerId must not be null"

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
		return errors.New(DescriptionInvalid)
	}
	if p.ExpirationRate == nil {
		return errors.New(ExpirationInvalid)
	}
	if p.FreezingRate == nil {
		return errors.New(FreezingInvalid)
	}
	if p.Height == nil || p.Length == nil || p.Width == nil {
		return errors.New(DimensionsInvalid)
	}
	if p.NetWeight == nil {
		return errors.New(NetWeightInvalid)
	}
	if p.ProductCode == nil {
		return errors.New(ProductCodeInvalid)
	}
	if p.ProductTypeId == nil {
		return errors.New(ProductTypeInvalid)
	}
	if p.SellerId == nil {
		return errors.New(SellerIdInvalid)
	}
	return nil
}
