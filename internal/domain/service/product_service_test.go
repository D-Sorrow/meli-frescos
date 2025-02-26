package service_test

import (
	models2 "github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repository2 "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	service2 "github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/models"
	"github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProductService_SaveProduct_Create(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("SaveProduct", mock.Anything).Return(nil)

	serviceTest := service.NewProductService(mockRepository)

	err := serviceTest.SaveProduct(models.ReturnMockProductModel())

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)

}

func TestProductService_SaveProduct_Conflict(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("SaveProduct", mock.Anything).Return(repository2.ErrRepositoryProductAlreadyExists)

	serviceTest := service.NewProductService(mockRepository)

	err := serviceTest.SaveProduct(models.ReturnMockProductModel())

	mockRepository.AssertExpectations(t)
	assert.Equal(t, err, service2.ErrServiceProductAlreadyExists)
}
func TestProductService_GetProducts(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("GetProducts").Return(models.ReturnProductModelMap(), nil)

	serviceTest := service.NewProductService(mockRepository)
	products, err := serviceTest.GetProducts()

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, len(products), len(models.ReturnProductModelMap()))
}
func TestProductService_GetProductByID_NonExistent(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("GetProductByID", 1).Return(models2.Product{}, repository2.ErrRepositoryProductNotFound)

	serviceTest := service.NewProductService(mockRepository)
	_, err := serviceTest.GetProductByID(1)

	mockRepository.AssertExpectations(t)
	assert.Equal(t, err, repository2.ErrRepositoryProductNotFound)
}

func TestProductService_GetProductByID_Existent(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("GetProductByID", 1).Return(models.ReturnMockProductModel(), nil)

	serviceTest := service.NewProductService(mockRepository)
	product, err := serviceTest.GetProductByID(1)

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, product, models.ReturnMockProductModel())
}
func TestProductService_UpdateProduct_Existent(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("UpdateProduct", 1, mock.Anything).Return(nil)
	mockRepository.On("GetProductByID", 1).Return(models.ReturnMockProductModel(), nil)
	att := models.ReturnMockProductModel()

	serviceTest := service.NewProductService(mockRepository)
	product, err := serviceTest.UpdateProduct(1, map[string]any{
		"description": att.Attributes.Description,
	})

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, models.ReturnMockProductModel(), product)
}

func TestProductService_UpdateProduct_NonExistent(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("UpdateProduct", 1, mock.Anything).Return(nil)
	mockRepository.On("GetProductByID", 1).Return(models2.Product{}, repository2.ErrRepositoryProductNotFound)
	att := models.ReturnMockProductModel()

	serviceTest := service.NewProductService(mockRepository)
	_, err := serviceTest.UpdateProduct(1, map[string]any{
		"description": att.Attributes.Description,
	})

	mockRepository.AssertNotCalled(t, "UpdateProduct")
	assert.Equal(t, err, service2.ErrServiceProductNotFound)
}

func TestProductService_DeleteProduct_NonExistent(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("DeleteProduct", 1).Return(repository2.ErrRepositoryProductNotFound)

	serviceTest := service.NewProductService(mockRepository)
	err := serviceTest.DeleteProduct(1)

	mockRepository.AssertExpectations(t)
	assert.Equal(t, err, repository2.ErrRepositoryProductNotFound)
}
func TestProductService_DeleteProduct_Existent(t *testing.T) {
	mockRepository := new(repository.ProductRepositoryMock)
	mockRepository.On("DeleteProduct", 1).Return(nil)

	serviceTest := service.NewProductService(mockRepository)
	err := serviceTest.DeleteProduct(1)

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)
}
