package repository

import (
	"database/sql"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	recordRepo "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entity"
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

func (repository *ProductRecordRepository) SaveProductRecord(productRecord models.ProductRecord) error {
	var productRecordEntity entity.ProductRecordEntity
	_, err := repository.db.Exec(productRecordEntity.SaveProductRecord(), productRecord.LastUpdateTime, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)
	if err != nil {
		return recordRepo.CodeSaveErr
	}
	return nil
}
func (repository *ProductRecordRepository) GetProductRecord(productId int) (map[int]models.ProductRecordResponse, error) {
	var productRecordEntity entity.ProductRecordEntity
	productRecordMap := make(map[int]models.ProductRecordResponse)

	rows, err := repository.db.Query(productRecordEntity.GetRecord(productId))
	if err != nil {
		log.Println(err)
		return nil, recordRepo.CodeGetErr
	}
	defer rows.Close()

	for rows.Next() {
		var productRecord models.ProductRecordResponse
		err := rows.Scan(&productRecord.ProductId, &productRecord.Description, &productRecord.RecordsCount)
		if err != nil {
			return nil, recordRepo.CodeGetErr
		}
		productRecordMap[productRecord.ProductId] = productRecord
	}
	return productRecordMap, nil
}
