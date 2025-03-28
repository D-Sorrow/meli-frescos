package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
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
