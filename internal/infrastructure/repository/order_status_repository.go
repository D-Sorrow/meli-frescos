package repository

import (
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/error_management"
)

type OrderStatusRepository struct {
	db *sql.DB
}

func NewOrderStatusRepository(db *sql.DB) *OrderStatusRepository {
	return &OrderStatusRepository{db: db}
}

func (b *OrderStatusRepository) GetAll() (buyers []entities.OrderStatusEntity, err error) {
	buyers = make([]entities.OrderStatusEntity, 0)

	query, args := (&entities.OrderStatusEntity{}).GetAllQuery()

	rows, err := b.db.Query(query, args...)
	if err != nil {
		err = error_management.HandleRepositoryError(repository.ErrOrderStatusUnexpectedError, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var buyer entities.OrderStatusEntity
		err = rows.Scan(&buyer.ID,
			&buyer.Description,
		)

		if err != nil {
			err = error_management.HandleRepositoryError(
				repository.ErrOrderStatusUnexpectedError,
				err,
			)
			return
		}

		buyers = append(buyers, buyer)
	}

	if len(buyers) == 0 {
		err = repository.ErrOrderStatusNoRegisteredOrderStatusesYet
		return
	}

	return
}
