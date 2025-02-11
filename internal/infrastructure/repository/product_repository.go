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
	productMap := make(map[int]models.Product)
	rows, err := p.db.Query("SELECT id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id FROM melifresh.products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Attributes.Description, &product.Attributes.ExpirationRate, &product.Attributes.FreezingRate, &product.Attributes.Dimensions.Height, &product.Attributes.Dimensions.Length, &product.Attributes.NetWeight, &product.Attributes.ProductCode, &product.Attributes.TemperatureFreezing, &product.Attributes.Dimensions.Width, &product.Attributes.ProductTypeId, &product.SellerId)
		if err != nil {
			panic(err)
		}
		productMap[product.Id] = product
	}
	return productMap

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
