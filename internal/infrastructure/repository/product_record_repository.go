package repository

import (
	"database/sql"
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	recordRepo "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entity"
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
