package error_management

import (
	"errors"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

func HandlerErrServiceProductRecord(err error) error {
	if errors.Is(err, repository.ErrRepositoryProductRecordNotFound) {
		return service.ErrServiceProductRecordNotFound
	}
	return service.ErrServiceProductRecordUnknown
}
