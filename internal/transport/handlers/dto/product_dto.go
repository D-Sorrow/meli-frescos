package dto

import (
	err "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"net/http"
)

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
		return err.Error{Code: http.StatusBadRequest, Message: err.DescriptionInvalid}
	}
	// More validation rules can be added here
	if p.ExpirationRate < 0 {
		return err.Error{Code: http.StatusBadRequest, Message: err.ExpirationInvalid}
	}
	if p.FreezingRate <= 0 {
		return err.Error{Code: http.StatusBadRequest, Message: err.FreezingInvalid}
	}
	if p.Height <= 0 || p.Length <= 0 || p.Width <= 0 {
		return err.Error{Code: http.StatusBadRequest, Message: err.DimensionsInvalid}
	}
	if p.NetWeight <= 0 {
		return err.Error{Code: http.StatusBadRequest, Message: err.NetWeightInvalid}
	}
	if p.ProductCode == "" {
		return err.Error{Code: http.StatusBadRequest, Message: err.ProductCodeInvalid}
	}
	if p.ProductTypeId <= 0 {
		return err.Error{Code: http.StatusBadRequest, Message: err.ProductTypeInvalid}
	}
	if p.SellerId <= 0 {
		return err.Error{Code: http.StatusBadRequest, Message: err.SellerIdInvalid}
	}
	return nil
}
