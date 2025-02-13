package repository

import "database/sql"

type CarryRepository struct {
	db *sql.DB
}

func NewCarryRepository(db *sql.DB) *CarryRepository {
	return &CarryRepository{db: db}
}
