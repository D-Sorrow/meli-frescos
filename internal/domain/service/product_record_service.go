package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	service2 "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
)

type ProductRecordService struct {
	repository repository.ProductRecordRepository
}

func NewProductRecordService(repository repository.ProductRecordRepository) *ProductRecordService {
	return &ProductRecordService{
		repository: repository,
	}
}

func (service *ProductRecordService) SaveProductRecord(
	productRecord models.ProductRecord,
) (models.ProductRecord, error) {
	errServ := productRecord.ValidateProductRecord()
	if errServ != nil {
		return models.ProductRecord{}, service2.ErrServiceProductRecordBusinessRules
	}
	productRecordResponse, errServ := service.repository.SaveProductRecord(productRecord)
	if errServ != nil {
		return models.ProductRecord{}, error_management.HandlerErrServiceProductRecord(errServ)
	}
	return productRecordResponse, nil
}

func (service *ProductRecordService) GetProductRecord(
	productId int,
) (map[int]models.ProductRecordResponse, error) {
	mapProductRecord, errSer := service.repository.GetProductRecord(productId)
	if errSer != nil {
		return map[int]models.ProductRecordResponse{}, error_management.HandlerErrServiceProductRecord(
			errSer,
		)
	}
	return mapProductRecord, nil
}
