package error_management

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

const CodeUseCaseError = "004"
const CodeDuplicatedCodeProduct = "007"

type ErrServiceProduct struct {
	Code string
}

func (e *ErrServiceProduct) Error() string {
	return e.Code
}

func HandlerServiceProductError(err error) error {
	switch {
	case errors.Is(err, repository.ErrRepositoryProductNotFound):
		return service.ErrServiceProductNotFound
	case errors.Is(err, repository.ErrRepositoryProductAlreadyExists):
		return service.ErrServiceProductAlreadyExists

	default:
		return service.ErrServiceProductUnknown
	}
}
