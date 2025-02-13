package repository

import (
	"database/sql"
	"fmt"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/mappers"
	"strings"

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

func (p ProductRepository) GetProductByID(id int) (models.Product, error) {
	var product models.Product

	row := p.db.QueryRow("SELECT id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id FROM melifresh.products WHERE  id = ?", id)
	err := row.Scan(&product.Id, &product.Attributes.Description, &product.Attributes.ExpirationRate, &product.Attributes.FreezingRate, &product.Attributes.Dimensions.Height, &product.Attributes.Dimensions.Length, &product.Attributes.NetWeight, &product.Attributes.ProductCode, &product.Attributes.TemperatureFreezing, &product.Attributes.Dimensions.Width, &product.Attributes.ProductTypeId, &product.SellerId)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, err
		}
	}
	return product, nil
}

func (p ProductRepository) SaveProduct(productSave models.Product) error {
	productEntity := mappers.ToProductEntity(productSave)
	query := "INSERT INTO melifresh.products" +
		"(description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := p.db.Exec(query, productEntity.Description, productEntity.ExpirationRate, productEntity.FreezingRate, productEntity.Height, productEntity.Length, productEntity.NetWeight, productEntity.ProductCode, productEntity.TemperatureFreezing, productEntity.Width, productEntity.ProductTypeId, productEntity.SellerId)

	if err != nil {
		return err
	}
	return nil
}
func (p ProductRepository) UpdateProduct(id int, attributes map[string]any) (models.Product, error) {
	setParts := make([]string, 0, len(attributes))
	values := make([]interface{}, 0, len(attributes))

	for columna, valor := range attributes {
		setParts = append(setParts, fmt.Sprintf("%s = ?", columna))
		values = append(values, valor)
	}

	setClause := strings.Join(setParts, ", ")

	sqlQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d ", "melifresh.products", setClause, id)

	_, err := p.db.Exec(sqlQuery, values...)

	if err != nil {
		return models.Product{}, err
	}
	product, errGet := p.GetProductByID(id)

	if errGet != nil {
		return models.Product{}, errGet
	}
	return product, nil
}
func (p ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM melifresh.products WHERE id = ?"

	err := p.db.QueryRow(query, id)

	if err.Err() != nil {
		return err.Err()
	}
	return nil
}
