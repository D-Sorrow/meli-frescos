package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
)

type ProductRecordService struct {
	repository repository.ProductRecordRepository
}

func NewProductRecordService(repository repository.ProductRecordRepository) *ProductRecordService {
	return &ProductRecordService{
		repository: repository,
	}
}
func (service *ProductRecordService) SaveProductRecord(productRecord models.ProductRecord) error {
	err := service.repository.SaveProductRecord(productRecord)
	if err != nil {
		return err
	}
	return nil
}
