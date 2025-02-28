package repository

import (
	"database/sql"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entity"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) GetProducts() (map[int]models.Product, error) {
	productMap := make(map[int]models.Product)
	productEntity := entity.ProductEntity{}
	rows, err := p.db.Query(productEntity.GetAllProducts())
	if err != nil {
		log.Println(err)
		return nil, repository.CodeQueryConsult
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Attributes.Description, &product.Attributes.ExpirationRate, &product.Attributes.FreezingRate, &product.Attributes.Dimensions.Height, &product.Attributes.Dimensions.Length, &product.Attributes.NetWeight, &product.Attributes.ProductCode, &product.Attributes.TemperatureFreezing, &product.Attributes.Dimensions.Width, &product.Attributes.ProductTypeId, &product.SellerId)
		if err != nil {
			return nil, repository.CodeGetProduct
		}
		productMap[product.Id] = product
	}
	return productMap, nil

}

func (p ProductRepository) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	productEntity := entity.ProductEntity{}
	row := p.db.QueryRow(productEntity.GetProductById(id))
	err := row.Scan(&product.Id, &product.Attributes.Description, &product.Attributes.ExpirationRate, &product.Attributes.FreezingRate, &product.Attributes.Dimensions.Height, &product.Attributes.Dimensions.Length, &product.Attributes.NetWeight, &product.Attributes.ProductCode, &product.Attributes.TemperatureFreezing, &product.Attributes.Dimensions.Width, &product.Attributes.ProductTypeId, &product.SellerId)

	if err != nil {
		return models.Product{}, repository.CodeGetProduct
	}
	return product, nil
}

func (p ProductRepository) SaveProduct(productSave models.Product) error {

	var productEntity entity.ProductEntity
	_, err := p.db.Exec(productEntity.SaveProduct(), productSave.Attributes.Description,
		productSave.Attributes.ExpirationRate, productSave.Attributes.FreezingRate,
		productSave.Attributes.Dimensions.Height, productSave.Attributes.Dimensions.Length,
		productSave.Attributes.NetWeight, productSave.Attributes.ProductCode,
		productSave.Attributes.TemperatureFreezing, productSave.Attributes.Dimensions.Width, productSave.Attributes.ProductTypeId, productSave.SellerId)
	if err != nil {
		log.Println(err)
		return repository.CodeNoRowsAffected
	}
	return nil
}
func (p ProductRepository) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	var productEntity entity.ProductEntity
	query, values := productEntity.UpdateProduct(attributes, id)
	result, _ := p.db.Exec(query, values...)

	if result != nil {
		if count, _ := result.RowsAffected(); count == 0 {
			return models.Product{}, repository.CodeNoRowsAffected
		}
	}
	product, errGet := p.GetProductByID(id)

	if errGet != nil {
		return models.Product{}, errGet
	}
	return product, nil
}
func (p ProductRepository) DeleteProduct(id int) error {
	var productEntity entity.ProductEntity
	result, err := p.db.Exec(productEntity.DeleteProduct(id))
	if result != nil {
		if count, _ := result.RowsAffected(); count == 0 {
			return repository.CodeDelete
		}
	}
	if err != nil {
		return repository.CodeDeleteIsNotPossible
	}
	return nil
}
