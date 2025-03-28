package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandleOrderStatusServiceError(err error) error {
	switch {
	case errors.Is(err, repository.ErrOrderStatusNoRegisteredOrderStatusesYet):
		return service.ErrOrderStatusNoRegisteredOrderStatusesYet
	default:
		return service.ErrOrderStatusUnexpectedError
	}
}
