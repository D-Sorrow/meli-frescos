package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/product/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/entity"
)

func JsonToProductModel(productJSON []entity.ProductJSON) map[int]models.Product {
	products := make(map[int]models.Product)

	for _, productJSON := range productJSON {
		products[productJSON.Id] = models.Product{
			Id:       productJSON.Id,
			SellerId: productJSON.SellerId,
			Attributes: models.ProductAttribute{
				ProductCode:         productJSON.ProductCode,
				Description:         productJSON.Description,
				NetWeight:           productJSON.NetWeight,
				ExpirationRate:      productJSON.ExpirationRate,
				TemperatureFreezing: productJSON.TemperatureFreezing,
				FreezingRate:        productJSON.FreezingRate,
				ProductTypeId:       productJSON.ProductTypeId,
				Dimensions: models.Dimensions{
					Width:  productJSON.Width,
					Height: productJSON.Height,
					Length: productJSON.Length,
				},
			},
		}
	}

	return products
}
