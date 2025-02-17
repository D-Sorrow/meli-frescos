package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func ToProductRecordModel(product *dto.ProductRecordDto) models.ProductRecord {
	return models.ProductRecord{
		ProductId:      *product.ProductId,
		PurchasePrice:  *product.PurchasePrice,
		SalePrice:      *product.SalePrice,
		LastUpdateTime: *product.LastUpdateTime,
	}
}
func ToProductRecordResponseSlice(productRecordMap map[int]models.ProductRecordResponse) []models.ProductRecordResponse {
	var productRecords []models.ProductRecordResponse
	for _, productRecord := range productRecordMap {
		productRecords[productRecord.ProductId] = productRecord
	}
	return productRecords
}
