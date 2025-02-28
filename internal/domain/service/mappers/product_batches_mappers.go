package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

func ProductBatchesToProductBatchesEntity(m *models.ProductBatches3) *entities.ProductBatchesEntity {
	return &entities.ProductBatchesEntity{
		Id:                 m.ID,
		BatchNumber:        m.ProductBatches2Attributes.BatchNumber,
		CurrentQuantity:    m.ProductBatches2Attributes.CurrentQuantity,
		CurrentTemperature: m.ProductBatches2Attributes.CurrentTemperature,
		DueDate:            m.ProductBatches2Attributes.DueDate,
		InitialQuantity:    m.ProductBatches2Attributes.InitialQuantity,
		ManufacturingDate:  m.ProductBatches2Attributes.ManufacturingDate,
		ManufacturingHour:  m.ProductBatches2Attributes.ManufacturingHour,
		MinumumTemperature: m.ProductBatches2Attributes.MinumumTemperature,
		ProductId:          m.ProductBatches2FKs.ProductId,
		SectionId:          m.ProductBatches2FKs.SectionId,
	}
}

func ProductBatches2AttributesFksToProductBatchesEntity(m *models.ProductBatches2AttributesFks) *entities.ProductBatchesEntity {
	return &entities.ProductBatchesEntity{
		Id:                 m.Id,
		BatchNumber:        m.BatchNumber,
		CurrentQuantity:    m.CurrentQuantity,
		CurrentTemperature: m.CurrentTemperature,
		DueDate:            m.DueDate,
		InitialQuantity:    m.InitialQuantity,
		ManufacturingDate:  m.ManufacturingDate,
		ManufacturingHour:  m.ManufacturingHour,
		MinumumTemperature: m.MinumumTemperature,
		ProductId:          m.ProductId,
		SectionId:          m.SectionId,
	}
}

func ProductBatchesEntityToProductBatches(m *entities.ProductBatchesEntity) *models.ProductBatches3 {
	return &models.ProductBatches3{
		ID: m.Id,
		ProductBatches2Attributes: models.ProductBatches2Attributes{
			BatchNumber:        m.BatchNumber,
			CurrentQuantity:    m.CurrentQuantity,
			CurrentTemperature: m.CurrentTemperature,
			DueDate:            m.DueDate,
			InitialQuantity:    m.InitialQuantity,
			ManufacturingDate:  m.ManufacturingDate,
			ManufacturingHour:  m.ManufacturingHour,
			MinumumTemperature: m.MinumumTemperature,
		},
		ProductBatches2FKs: models.ProductBatches2FKs{
			ProductId: m.ProductId,
			SectionId: m.SectionId,
		},
	}
}

func ReportProductBatchesEntityToReportProductBatches(m *entities.ReportProductBatchEntity) *models.ReportProductBatch {
	return &models.ReportProductBatch{
		Id:                  m.Id,
		Warehouse_id:        m.Warehouse_id,
		Product_type_id:     m.Product_type_id,
		Section_number:      m.Section_number,
		Current_temperature: m.Current_temperature,
		Minimum_temperature: m.Minimum_temperature,
		Current_capacity:    m.Current_capacity,
		Minimum_capacity:    m.Minimum_capacity,
		Maximum_capacity:    m.Maximum_capacity,
		ProductBatchesCount: m.ProductBatchesCount,
	}
}
