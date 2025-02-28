package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	er "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
	"github.com/go-sql-driver/mysql"
)

type ProductBatchesRepository struct {
	db *sql.DB
}

func NewProductBatchesRepository(db *sql.DB) *ProductBatchesRepository {
	return &ProductBatchesRepository{db: db}
}

func (p *ProductBatchesRepository) AddProductBatches(productBatches models.ProductBatches) (models.ProductBatches, error) {

	result, err := p.db.Exec("INSERT INTO product_batches (batch_number,current_quantity,current_temperature,due_date,initial_quantity,manufacturing_date,manufacturing_hour,minimum_temperature,product_id,section_id) values (?,?,?,?,?,?,?,?,?,?)", productBatches.BatchNumber, productBatches.CurrentQuantity, productBatches.CurrentTemperature, productBatches.DueDate, productBatches.InitialQuantity, productBatches.ManufacturingDate, productBatches.ManufacturingHour, productBatches.MinumumTemperature, productBatches.ProductId, productBatches.SectionId)

	if err != nil {
		log.Print(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return models.ProductBatches{}, er.ErrProductBatchesAlredyExists
		}
		return models.ProductBatches{}, er.ErrProductBatchesNotCreated
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return models.ProductBatches{}, er.ErrProductBatchesNotCreated
	}

	productBatches.Id = int(lastInsertId)
	return productBatches, nil
}

func (p *ProductBatchesRepository) GetById(id int) (productBatches entities.ProductBatchesEntity, err error) {
	query, args := productBatches.GetByIdQuery(id)

	err = p.db.QueryRow(query, args...).Scan(
		&productBatches.Id,
		&productBatches.BatchNumber,
		&productBatches.CurrentQuantity,
		&productBatches.CurrentTemperature,
		&productBatches.DueDate,
		&productBatches.InitialQuantity,
		&productBatches.ManufacturingDate,
		&productBatches.ManufacturingHour,
		&productBatches.MinumumTemperature,
		&productBatches.ProductId,
		&productBatches.SectionId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repository.ErrProductBatchNotFoundWithID
			return
		}
	}

	return
}

func (b *ProductBatchesRepository) Create(productBatches entities.ProductBatchesEntity) (newProductBatches entities.ProductBatchesEntity, err error) {
	query, args := productBatches.GetCreateQuery()
	result, err := b.db.Exec(query, args...)

	if err != nil {
		var mySqlErr *mysql.MySQLError
		if errors.As(err, &mySqlErr) && mySqlErr.Number == 1452 {
			err = repository.ErrForeignKeysNotValid
			return
		}

		return
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return
	}

	newProductBatches, err = b.GetById(int(lastId))
	return
}
