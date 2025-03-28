package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) GetProducts() (map[int]models.Product, error) {
	productMap := make(map[int]models.Product)
	productEntity := entities.ProductEntity{}
	rows, err := p.db.Query(productEntity.GetAllProducts())
	if err != nil {
		log.Println(err)
		return nil, repository.ErrRepositoryProductUnknown
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Attributes.Description,
			&product.Attributes.ExpirationRate,
			&product.Attributes.FreezingRate,
			&product.Attributes.Dimensions.Height,
			&product.Attributes.Dimensions.Length,
			&product.Attributes.NetWeight,
			&product.Attributes.ProductCode,
			&product.Attributes.TemperatureFreezing,
			&product.Attributes.Dimensions.Width,
			&product.Attributes.ProductTypeId,
			&product.SellerId,
		)
		if err != nil {
			return nil, repository.ErrRepositoryProductUnknown
		}
		productMap[product.Id] = product
	}
	return productMap, nil

}

func (p ProductRepository) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	productEntity := entities.ProductEntity{}
	row := p.db.QueryRow(productEntity.GetProductById(id))
	err := row.Scan(
		&product.Id,
		&product.Attributes.Description,
		&product.Attributes.ExpirationRate,
		&product.Attributes.FreezingRate,
		&product.Attributes.Dimensions.Height,
		&product.Attributes.Dimensions.Length,
		&product.Attributes.NetWeight,
		&product.Attributes.ProductCode,
		&product.Attributes.TemperatureFreezing,
		&product.Attributes.Dimensions.Width,
		&product.Attributes.ProductTypeId,
		&product.SellerId,
	)

	if err != nil {
		return models.Product{}, repository.ErrRepositoryProductNotFound
	}
	return product, nil
}

func (p ProductRepository) SaveProduct(productSave models.Product) error {

	var productEntity entities.ProductEntity
	var errSql *mysql.MySQLError
	_, err := p.db.Exec(
		productEntity.SaveProduct(),
		productSave.Attributes.Description,
		productSave.Attributes.ExpirationRate,
		productSave.Attributes.FreezingRate,
		productSave.Attributes.Dimensions.Height,
		productSave.Attributes.Dimensions.Length,
		productSave.Attributes.NetWeight,
		productSave.Attributes.ProductCode,
		productSave.Attributes.TemperatureFreezing,
		productSave.Attributes.Dimensions.Width,
		productSave.Attributes.ProductTypeId,
		productSave.SellerId,
	)
	if err != nil {
		log.Println(err)
		if errors.As(err, &errSql) {
			switch errSql.Number {
			case 1062:
				return repository.ErrRepositoryProductAlreadyExists
			default:
				return repository.ErrRepositoryProductUnknown
			}
		}
	}
	return nil
}
func (p ProductRepository) UpdateProduct(id int, attributes map[string]any) error {
	var productEntity entities.ProductEntity
	var errSql *mysql.MySQLError

	query, values := productEntity.UpdateProduct(attributes, id)
	result, err := p.db.Exec(query, values...)
	if errors.As(err, &errSql) {
		switch errSql.Number {
		case 1062:
			return repository.ErrRepositoryProductAlreadyExists
		default:
			return repository.ErrRepositoryProductUnknown
		}
	}
	if row, _ := result.RowsAffected(); row == 0 {
		return repository.ErrRepositoryProductNotFound
	}
	return nil
}
func (p ProductRepository) DeleteProduct(id int) error {
	var productEntity entities.ProductEntity
	result, err := p.db.Exec(productEntity.DeleteProduct(id))
	if result != nil {
		if count, _ := result.RowsAffected(); count == 0 {
			return repository.ErrRepositoryProductNotFound
		}
	}
	if err != nil {
		return repository.ErrRepositoryProductUnknown
	}
	return nil
}
