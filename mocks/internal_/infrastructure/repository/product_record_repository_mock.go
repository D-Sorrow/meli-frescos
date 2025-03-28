package repository_mock

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type ProductRecordRepositoryMock struct {
	mock.Mock
}

func (p ProductRecordRepositoryMock) SaveProductRecord(
	productRecord models.ProductRecord,
) (models.ProductRecord, error) {
	args := p.Called(productRecord)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (p ProductRecordRepositoryMock) GetProductRecord(
	productId int,
) (map[int]models.ProductRecordResponse, error) {
	args := p.Called(productId)
	return args.Get(0).(map[int]models.ProductRecordResponse), args.Error(1)
}
