package repository

import (
	"database/sql"
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	recordRepo "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	"github.com/go-sql-driver/mysql"
	"log"
)

type ProductRecordRepository struct {
	db *sql.DB
}

func NewProductRecordRepository(db *sql.DB) *ProductRecordRepository {
	return &ProductRecordRepository{
		db: db,
	}
}

func (repository *ProductRecordRepository) SaveProductRecord(productRecord models.ProductRecord) (models.ProductRecord, error) {
	var productRecordEntity entities.ProductRecordEntity
	var sqlErr *mysql.MySQLError
	_, err := repository.db.Exec(productRecordEntity.SaveProductRecord(), productRecord.LastUpdateTime, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)
	if err != nil {
		errors.As(err, &sqlErr)
		switch sqlErr.Number {
		case 1452:
			return models.ProductRecord{}, recordRepo.ErrRepositoryProductRecordNotFound
		}
	}
	return productRecord, nil
}
func (repository *ProductRecordRepository) GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error) {
	var productRecordEntity entities.ProductRecordEntity
	productRecordMap := make(map[int]models.ProductRecordResponse)

	rows, err := repository.db.Query(productRecordEntity.GetRecord(productId))
	if err != nil {
		log.Println(err)
		return nil, recordRepo.ErrRepositoryProductRecordNotFound
	}
	defer rows.Close()

	for rows.Next() {
		var productRecord models.ProductRecordResponse
		err := rows.Scan(&productRecord.ProductId, &productRecord.Description, &productRecord.RecordsCount)
		if err != nil {
			return nil, recordRepo.ErrRepositoryProductRecordNotFound
		}
		productRecordMap[productRecord.ProductId] = productRecord
	}
	return productRecordMap, nil
}
