package repository

import (
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/error_management"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (b *LogRepository) Register(log entities.LogEntity) (newLog entities.LogEntity, err error) {
	query, args := log.GetCreateQuery()
	result, err := b.db.Exec(query, args...)

	if err != nil {
		err = error_management.HandleRepositoryError(
			error_management.HandleLogRepositoryError(err),
			err,
		)
		return
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		err = error_management.HandleRepositoryError(
			repository.ErrLogUnexpectedError,
			err,
		)
		return
	}

	log.ID = lastId
	newLog = log
	return
}
