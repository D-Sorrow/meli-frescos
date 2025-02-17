package entity

import (
	"fmt"
	"strconv"
	"strings"
)

type ProductEntity struct {
	Id                  int     `json:"id"`
	Description         string  `json:"description"`
	ExpirationRate      int     `json:"expiration_rate"`
	FreezingRate        int     `json:"freezing_rate"`
	Height              float64 `json:"height"`
	Length              float64 `json:"length"`
	NetWeight           float64 `json:"netweight"`
	ProductCode         string  `json:"product_code"`
	TemperatureFreezing float64 `json:"recommended_freezing_temperature"`
	Width               float64 `json:"width"`
	ProductTypeId       int     `json:"product_type_id"`
	SellerId            int     `json:"seller_id"`
}

func (entity *ProductEntity) GetAllProducts() string {
	return "SELECT " +
		"id, description, expiration_rate, freezing_rate, " +
		"height, length, netweight, product_code, recommended_freezing_temperature, " +
		"width, product_type_id, seller_id " +
		"FROM melifresh.products"
}
func (entity *ProductEntity) GetProductById(id int) string {
	return "SELECT id, description, expiration_rate, freezing_rate, height, length, netweight, " +
		"product_code, recommended_freezing_temperature, width, product_type_id, seller_id " +
		"FROM melifresh.products WHERE  id = " + strconv.Itoa(id)
}
func (entity *ProductEntity) SaveProduct() string {
	return "INSERT INTO melifresh.products (description, expiration_rate, freezing_rate, height, " +
		"length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, " +
		"seller_id) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func (entity *ProductEntity) UpdateProduct(attributes map[string]any, id int) (string, []interface{}) {
	setParts := make([]string, 0, len(attributes))
	values := make([]interface{}, 0, len(attributes))

	for column, value := range attributes {
		setParts = append(setParts, fmt.Sprintf("%s = ?", column))
		values = append(values, value)
	}

	setClause := strings.Join(setParts, ", ")
	return fmt.Sprintf("UPDATE %s SET %s WHERE id = %d ", "melifresh.products", setClause, id), values
}
func (entity *ProductEntity) DeleteProduct(id int) string {
	return "DELETE FROM melifresh.products WHERE id = " + strconv.Itoa(id)
}
