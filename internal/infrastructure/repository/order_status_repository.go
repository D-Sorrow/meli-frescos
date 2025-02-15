package repository

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
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
		return
	}
	defer rows.Close()

	for rows.Next() {
		var buyer entities.OrderStatusEntity
		err = rows.Scan(&buyer.ID,
			&buyer.Description,
		)

		if err != nil {
			return
		}

		buyers = append(buyers, buyer)
	}

	if len(buyers) == 0 {
		err = repository.ErrNoRegisteredOrderStatusesYet
		return
	}

	return
}
