package models

import "errors"

type Product struct {
	Id         int
	SellerId   int
	Attributes ProductAttribute
}

type ProductAttribute struct {
	ProductCode         string     `json:"product_code,omitempty"`
	Description         string     `json:"description,omitempty"`
	NetWeight           float64    `json:"netweight,omitempty"`
	ExpirationRate      int        `json:"expiration_rate,omitempty"`
	TemperatureFreezing float64    `json:"recommended_freezing_temperature,omitempty"`
	FreezingRate        int        `json:"freezing_rate,omitempty"`
	ProductTypeId       int        `json:"product_type_id,omitempty"`
	Dimensions          Dimensions `json:"omitempty"`
}

type Dimensions struct {
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Length float64 `json:"length,omitempty"`
}

const expirationInvalid = "rates must not be negative"
const freezingInvalid = "freezingRate must not be positive"
const dimensionsInvalid = "invalid dimensions"
const netWeightInvalid = "net weight must be positive"
const productTypeInvalid = "productTypeId must not be negative"
const sellerIdInvalid = "sellerId must not be negative"

func (p *Product) ValidateProduct() error {
	if p.Attributes.ExpirationRate < 0 {
		return errors.New(expirationInvalid)
	}
	if p.Attributes.FreezingRate >= 0 {
		return errors.New(freezingInvalid)
	}
	if p.Attributes.Dimensions.Height <= 0 || p.Attributes.Dimensions.Length <= 0 || p.Attributes.Dimensions.Width <= 0 {
		return errors.New(dimensionsInvalid)
	}
	if p.Attributes.NetWeight <= 0 {
		return errors.New(netWeightInvalid)
	}
	if p.Attributes.ProductTypeId <= 0 {
		return errors.New(productTypeInvalid)
	}
	if p.SellerId <= 0 {
		return errors.New(sellerIdInvalid)
	}
	return nil
}
