package models

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
