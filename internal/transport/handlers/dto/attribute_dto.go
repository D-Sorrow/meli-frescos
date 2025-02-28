package dto

import "errors"

type AttributeDto struct {
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

func (attribute *AttributeDto) Validation() error {
	if attribute.Description != nil {
		if *attribute.Description == "" {
			return errors.New(DescriptionInvalid)
		}
	}
	if attribute.ExpirationRate != nil {
		if *attribute.ExpirationRate < 0 {
			return errors.New(ExpirationInvalid)
		}
	}
	if attribute.FreezingRate != nil {
		if *attribute.FreezingRate >= 0 {
			return errors.New(FreezingInvalid)
		}
	}
	if attribute.Height != nil {
		if *attribute.Height <= 0 {
			return errors.New(DimensionsInvalid)
		}
	}
	if attribute.Length != nil {
		if *attribute.Length <= 0 {
			return errors.New(DimensionsInvalid)
		}
	}
	if attribute.Width != nil {
		if *attribute.Width <= 0 {
			return errors.New(DimensionsInvalid)
		}
	}
	if attribute.NetWeight != nil {
		if *attribute.NetWeight <= 0 {
			return errors.New(NetWeightInvalid)
		}
	}
	if attribute.ProductCode != nil {
		if *attribute.ProductCode == "" {
			return errors.New(ProductCodeInvalid)
		}
	}
	if attribute.ProductTypeId != nil {
		if *attribute.ProductTypeId <= 0 {
			return errors.New(ProductTypeInvalid)
		}
	}
	if attribute.SellerId != nil {
		if *attribute.SellerId <= 0 {
			return errors.New(SellerIdInvalid)
		}
	}
	return nil
}
