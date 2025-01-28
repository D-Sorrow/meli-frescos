package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToProductDto(product models.Product) dto.ProductDto {
	return dto.ProductDto{
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

func MapperToProductsDto(product map[int]models.Product) []dto.ProductDto {
	var productSliceDto []dto.ProductDto
	for _, p := range product {
		productDto := MapperToProductDto(p)
		productSliceDto = append(productSliceDto, productDto)
	}
	return productSliceDto
}

func MapperToProductModel(product dto.ProductDto) models.Product {
	return models.Product{
		Id:       product.Id,
		SellerId: product.SellerId,
		Attributes: models.ProductAttribute{
			ProductCode:         product.ProductCode,
			Description:         product.Description,
			NetWeight:           product.NetWeight,
			ExpirationRate:      product.ExpirationRate,
			TemperatureFreezing: product.TemperatureFreezing,
			FreezingRate:        product.FreezingRate,
			ProductTypeId:       product.ProductTypeId,
			Dimensions: models.Dimensions{
				Width:  product.Width,
				Height: product.Height,
				Length: product.Length,
			},
		},
	}
}
