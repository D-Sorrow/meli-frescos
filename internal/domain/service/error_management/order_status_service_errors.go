package error_management

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

func HandleOrderStatusServiceError(err error) error {
	switch {
	case errors.Is(err, repository.ErrOrderStatusNoRegisteredOrderStatusesYet):
		return service.ErrOrderStatusNoRegisteredOrderStatusesYet
	default:
		return service.ErrOrderStatusUnexpectedError
	}
}
