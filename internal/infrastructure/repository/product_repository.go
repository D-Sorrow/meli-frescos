package repository

import (
	"database/sql"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) GetProducts() map[int]models.Product {
	panic("implement me")
}

func (p ProductRepository) GetProductByID(id int) (models.Product, bool) {
	/*product, ok := p.productMap[id]
	if !ok {
		return models.Product{}, false
	}
	return product, true*/
	panic("implement me")
}

func (p ProductRepository) SaveProduct(productSave models.Product) bool {
	/*indexToSave := 0

	for _, product := range p.productMap {
		if product.Attributes.ProductCode == productSave.Attributes.ProductCode {
			return false
		}
		indexToSave += 1
	}
	indexToSave++
	productSave.Id = indexToSave
	p.productMap[indexToSave] = productSave
	return true*/
	panic("implement me")
}
func (p ProductRepository) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	/*product, ok := p.productMap[id]
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
	return product, nil*/
	panic("implement me")
}
func (p ProductRepository) DeleteProduct(id int) bool {
	panic("implement me")
}
func (p ProductRepository) validateCode(code string) bool {
	panic("implement me")
}
