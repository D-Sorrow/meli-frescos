package service_test

import (
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	service3 "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	service2 "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	repository_mock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductRecordService_SaveProductRecord(t *testing.T) {
	t.Run("save product record success", func(t *testing.T) {
		productRecordMock := models.ProductRecord{
			ProductId:      1,
			PurchasePrice:  19.00,
			SalePrice:      19.00,
			LastUpdateTime: "2026-01-02",
		}
		mockRepository := new(repository_mock.ProductRecordRepositoryMock)
		mockRepository.On("SaveProductRecord", productRecordMock).Return(productRecordMock, nil)

		service := service2.NewProductRecordService(mockRepository)
		proRecord, err := service.SaveProductRecord(productRecordMock)

		assert.Nil(t, err)
		assert.Equal(t, productRecordMock, proRecord)
		mockRepository.AssertExpectations(t)
	})
	t.Run("save product record error by validation", func(t *testing.T) {
		productRecordMock := models.ProductRecord{
			ProductId:      1,
			PurchasePrice:  -19.00,
			SalePrice:      19.00,
			LastUpdateTime: "2026-01-02",
		}
		mockRepository := new(repository_mock.ProductRecordRepositoryMock)

		service := service2.NewProductRecordService(mockRepository)
		record, err := service.SaveProductRecord(productRecordMock)

		assert.NotNil(t, err)
		assert.Equal(t, service3.ErrServiceProductRecordBusinessRules, err)
		assert.Equal(t, models.ProductRecord{}, record)
		mockRepository.AssertExpectations(t)
	})
	t.Run("save product record error by validation", func(t *testing.T) {
		productRecordMock := models.ProductRecord{
			ProductId:      1,
			PurchasePrice:  19.00,
			SalePrice:      19.00,
			LastUpdateTime: "2026-01-02",
		}
		mockRepository := new(repository_mock.ProductRecordRepositoryMock)
		mockRepository.On("SaveProductRecord", productRecordMock).
			Return(models.ProductRecord{}, repository.ErrRepositoryProductRecordNotFound)

		service := service2.NewProductRecordService(mockRepository)
		record, err := service.SaveProductRecord(productRecordMock)

		assert.NotNil(t, err)
		assert.Equal(t, service3.ErrServiceProductRecordNotFound, err)
		assert.Equal(t, models.ProductRecord{}, record)
		mockRepository.AssertExpectations(t)
	})
}

func TestProductRecordService_GetProductRecord(t *testing.T) {
	t.Run("get product record success", func(t *testing.T) {
		mockRepository := new(repository_mock.ProductRecordRepositoryMock)
		mockRepository.On("GetProductRecord",
			mock.AnythingOfType("int")).Return(map[int]models.ProductRecordResponse{}, nil)

		service := service2.NewProductRecordService(mockRepository)
		mapProductRecord, err := service.GetProductRecord(1)

		assert.Nil(t, err)
		assert.Equal(t, map[int]models.ProductRecordResponse{}, mapProductRecord)
		mockRepository.AssertExpectations(t)
	})

	t.Run("get product record error", func(t *testing.T) {
		mockRepository := new(repository_mock.ProductRecordRepositoryMock)
		mockRepository.On(
			"GetProductRecord",
			mock.AnythingOfType(
				"int",
			),
		).Return(map[int]models.ProductRecordResponse{}, repository.ErrRepositoryProductRecordNotFound)

		service := service2.NewProductRecordService(mockRepository)
		mapProductRecord, err := service.GetProductRecord(1)

		assert.NotNil(t, err)
		assert.Equal(t, service3.ErrServiceProductRecordNotFound, err)
		assert.Equal(t, map[int]models.ProductRecordResponse{}, mapProductRecord)
		mockRepository.AssertExpectations(t)
	})
}
