package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entity"
)

func ToProductEntity(product models.Product) entity.ProductEntity {
	return entity.ProductEntity{
		Id:                  product.Id,
		SellerId:            product.SellerId,
		ProductCode:         product.Attributes.ProductCode,
		Description:         product.Attributes.Description,
		NetWeight:           product.Attributes.NetWeight,
		ExpirationRate:      product.Attributes.ExpirationRate,
		TemperatureFreezing: product.Attributes.TemperatureFreezing,
		FreezingRate:        product.Attributes.FreezingRate,
		ProductTypeId:       product.Attributes.ProductTypeId,
		Width:               product.Attributes.Dimensions.Width,
		Height:              product.Attributes.Dimensions.Height,
		Length:              product.Attributes.Dimensions.Length,
	}
}
