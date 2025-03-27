package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
)

func HandleLogRepositoryError(err error) error {
	return repository.ErrLogUnexpectedError
}
