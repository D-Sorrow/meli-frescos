package dto

import (
	"errors"
)

const IdInvalid = "ID must be positive"
const DescriptionInvalid = "Description must not be empty."
const ExpirationInvalid = "Rates must not be negative"
const FreezingInvalid = "FreezingRate must not be negative"
const DimensionsInvalid = "Invalid dimensions"
const NetWeightInvalid = "Net weight must be positive"
const ProductCodeInvalid = "ProductCode must not be empty"
const ProductTypeInvalid = "ProductTypeId must not be negative"
const SellerIdInvalid = "SellerId must not be negative"

type ProductDto struct {
	Id                  int     `json:"id"`
	Description         string  `json:"description"`
	ExpirationRate      int     `json:"expiration_rate"`
	FreezingRate        int     `json:"freezing_rate"`
	Height              float64 `json:"height"`
	Length              float64 `json:"length"`
	NetWeight           float64 `json:"netweight"`
	ProductCode         string  `json:"product_code"`
	TemperatureFreezing float64 `json:"recommended_freezing_temperature"`
	Width               float64 `json:"width"`
	ProductTypeId       int     `json:"product_type_id"`
	SellerId            int     `json:"seller_id"`
}

func (p *ProductDto) Validate() error {
	if p.Description == "" {
		return errors.New(DescriptionInvalid)
	}
	if p.ExpirationRate < 0 {
		return errors.New(ExpirationInvalid)
	}
	if p.FreezingRate <= 0 {
		return errors.New(FreezingInvalid)
	}
	if p.Height <= 0 || p.Length <= 0 || p.Width <= 0 {
		return errors.New(DimensionsInvalid)
	}
	if p.NetWeight <= 0 {
		return errors.New(NetWeightInvalid)
	}
	if p.ProductCode == "" {
		return errors.New(ProductCodeInvalid)
	}
	if p.ProductTypeId <= 0 {
		return errors.New(ProductTypeInvalid)
	}
	if p.SellerId <= 0 {
		return errors.New(SellerIdInvalid)
	}
	return nil
}
