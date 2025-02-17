package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToProductBatches(productBatches dto.ProductBatchesDto) models.ProductBatches {
	return models.ProductBatches{
		Id:                 productBatches.Id,
		BatchNumber:        productBatches.BatchNumber,
		CurrentQuantity:    productBatches.CurrentQuantity,
		CurrentTemperature: productBatches.CurrentTemperature,
		DueDate:            productBatches.DueDate,
		InitialQuantity:    productBatches.InitialQuantity,
		ManufacturingDate:  productBatches.ManufacturingDate,
		ManufacturingHour:  productBatches.ManufacturingHour,
		MinumumTemperature: productBatches.MinumumTemperature,
		ProductId:          productBatches.ProductId,
		SectionId:          productBatches.SectionId,
	}
}

func MapperToProductBatchesDTO(productBatches models.ProductBatches) dto.ProductBatchesDto {
	return dto.ProductBatchesDto{
		Id:                 productBatches.Id,
		BatchNumber:        productBatches.BatchNumber,
		CurrentQuantity:    productBatches.CurrentQuantity,
		CurrentTemperature: productBatches.CurrentTemperature,
		DueDate:            productBatches.DueDate,
		InitialQuantity:    productBatches.InitialQuantity,
		ManufacturingDate:  productBatches.ManufacturingDate,
		ManufacturingHour:  productBatches.ManufacturingHour,
		MinumumTemperature: productBatches.MinumumTemperature,
		ProductId:          productBatches.ProductId,
		SectionId:          productBatches.SectionId,
	}
}

func MapperToProductBatchesDTO2(productBatches *models.ProductBatches3) dto.ProductBatchesDto {
	return dto.ProductBatchesDto{
		Id:                 productBatches.ID,
		BatchNumber:        productBatches.ProductBatches2Attributes.BatchNumber,
		CurrentQuantity:    productBatches.ProductBatches2Attributes.CurrentQuantity,
		CurrentTemperature: productBatches.ProductBatches2Attributes.CurrentTemperature,
		DueDate:            productBatches.ProductBatches2Attributes.DueDate,
		InitialQuantity:    productBatches.ProductBatches2Attributes.InitialQuantity,
		ManufacturingDate:  productBatches.ProductBatches2Attributes.ManufacturingDate,
		ManufacturingHour:  productBatches.ProductBatches2Attributes.ManufacturingHour,
		MinumumTemperature: productBatches.ProductBatches2Attributes.MinumumTemperature,
		ProductId:          productBatches.ProductBatches2FKs.ProductId,
		SectionId:          productBatches.ProductBatches2FKs.SectionId,
	}
}

func ProductBatchesCreateDTOToPProductBatchesFKs(productBatches *dto.ProductBatchesDtoReport) *models.ProductBatches2AttributesFks {
	return &models.ProductBatches2AttributesFks{
		ProductBatches2FKs: models.ProductBatches2FKs{

			ProductId: productBatches.ProductId,
			SectionId: productBatches.SectionId,
		},
	}
}
