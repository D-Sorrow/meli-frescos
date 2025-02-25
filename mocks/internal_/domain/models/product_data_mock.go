package internal

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

var (
	description         = "Test Product"
	expirationRate      = 30
	freezingRate        = -10
	height              = 10.5
	length              = 5.0
	netWeight           = 1.5
	productCode         = "TP001"
	temperatureFreezing = -18.0
	width               = 15.0
	productTypeId       = 2
	sellerId            = 3
)

func ReturnMockProduct() dto.ProductDto {

	return dto.ProductDto{
		Description:         &description,
		ExpirationRate:      &expirationRate,
		FreezingRate:        &freezingRate,
		Height:              &height,
		Length:              &length,
		NetWeight:           &netWeight,
		ProductCode:         &productCode,
		TemperatureFreezing: &temperatureFreezing,
		Width:               &width,
		ProductTypeId:       &productTypeId,
		SellerId:            &sellerId,
	}
}

func ReturnMockProductModel() models.Product {
	return models.Product{
		SellerId: sellerId,
		Attributes: models.ProductAttribute{
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
		},
	}
}

func ReturnProductModelMap() map[int]models.Product {
	return map[int]models.Product{
		1: models.Product{
			SellerId: sellerId,
			Attributes: models.ProductAttribute{
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
			},
		},
	}
}
