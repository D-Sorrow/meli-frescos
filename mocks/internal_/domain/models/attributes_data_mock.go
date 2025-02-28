package models

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

func ReturnAttributesModel() models.ProductAttribute {
	return models.ProductAttribute{
		Description:         description,
		ExpirationRate:      expirationRate,
		FreezingRate:        freezingRate,
		ProductCode:         productCode,
		NetWeight:           netWeight,
		TemperatureFreezing: temperatureFreezing,
		ProductTypeId:       productTypeId,
		Dimensions: models.Dimensions{
			Width:  width,
			Height: height,
			Length: length,
		},
	}
}
