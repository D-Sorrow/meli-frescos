package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
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
	errServ := productRecord.ValidateProductRecord()
	if errServ != nil {
		return error_management.CodeErrBusiness
	}
	return service.repository.SaveProductRecord(productRecord)
}
