package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandlerErrServiceProductRecord(err error) error {
	if errors.Is(err, repository.ErrRepositoryProductRecordNotFound) {
		return service.ErrServiceProductRecordNotFound
	}
	return service.ErrServiceProductRecordUnknown
}
