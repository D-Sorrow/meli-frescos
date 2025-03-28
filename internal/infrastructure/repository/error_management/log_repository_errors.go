package error_management

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
)

func HandleLogRepositoryError(err error) error {
	return repository.ErrLogUnexpectedError
}
