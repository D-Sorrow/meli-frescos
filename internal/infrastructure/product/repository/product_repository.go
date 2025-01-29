package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/product/models"
	err "github.com/D-Sorrow/meli-frescos/internal/infrastructure/product/repository/error_management"
)

type ProductRepository struct {
	productMap map[int]models.Product
}

func NewProductRepository(productMap map[int]models.Product) *ProductRepository {
	return &ProductRepository{productMap: productMap}
}

func (p ProductRepository) GetProducts() map[int]models.Product {
	return p.productMap
}

func (p ProductRepository) GetProductByID(id int) (models.Product, bool) {
	product, ok := p.productMap[id]
	if !ok {
		return models.Product{}, false
	}
	return product, true
}

func (p ProductRepository) SaveProduct(productSave models.Product) bool {
	indexToSave := 0

	for _, product := range p.productMap {
		if product.Attributes.ProductCode == productSave.Attributes.ProductCode {
			return false
		}
		indexToSave += 1
	}
	indexToSave++
	productSave.Id = indexToSave
	p.productMap[indexToSave] = productSave
	return true
}
func (p ProductRepository) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	product, ok := p.productMap[id]
	if !ok {
		return models.Product{}, err.ProductNotExistErr
	}
	for key, value := range attributes {
		switch key {
		case "description":
			product.Attributes.Description = value.(string)
		case "expiration_rate":
			product.Attributes.ExpirationRate = value.(int)
		case "freezing_rate":
			product.Attributes.FreezingRate = value.(int)
		case "height":
			product.Attributes.Dimensions.Height = value.(float64)
		case "length":
			product.Attributes.Dimensions.Length = value.(float64)
		case "netweight":
			product.Attributes.NetWeight = value.(float64)
		case "product_code":
			if p.validateCode(value.(string)) {
				return models.Product{}, err.CodeProductIsExistErr
			}
			product.Attributes.ProductCode = value.(string)
		case "recommended_freezing_temperature":
			product.Attributes.ProductCode = value.(string)
		case "width":
			product.Attributes.Dimensions.Width = value.(float64)
		case "product_type_id":
			product.Attributes.ProductTypeId = value.(int)
		case "seller_id":
			product.SellerId = value.(int)
		default:
			return models.Product{}, err.AttributeIsNotValidErr
		}
	}
	p.productMap[id] = product
	return product, nil
}
func (p ProductRepository) DeleteProduct(id int) bool {
	_, ok := p.productMap[id]
	if !ok {
		return false
	}
	delete(p.productMap, id)
	return true
}
func (p ProductRepository) validateCode(code string) bool {
	for _, product := range p.productMap {
		if product.Attributes.ProductCode == code {
			return true
		}
	}
	return false
}
