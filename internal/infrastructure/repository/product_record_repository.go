package repository

import (
	"database/sql"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type ProductRecordRepository struct {
	db *sql.DB
}

func NewProductRecordRepository(db *sql.DB) *ProductRecordRepository {
	return &ProductRecordRepository{
		db: db,
	}
}

func (repository *ProductRecordRepository) SaveProductRecord(productRecord models.ProductRecord) error {
	query := "INSERT INTO melifresh.product_records" +
		"(last_update_date, purchase_price, sale_price, product_id)" +
		"VALUES (?, ?, ?, ?)"

	_, err := repository.db.Exec(query, productRecord.LastUpdateTime, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)
	if err != nil {
		return err
	}
	return nil
}
