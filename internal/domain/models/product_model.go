package models

type Product struct {
	Id         int
	SellerId   int
	Attributes ProductAttribute
}

type ProductAttribute struct {
	ProductCode         string
	Description         string
	NetWeight           float64
	ExpirationRate      int
	TemperatureFreezing float64
	FreezingRate        int
	ProductTypeId       int
	Dimensions          Dimensions
}

type Dimensions struct {
	Width  float64
	Height float64
	Length float64
}
