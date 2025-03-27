package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

func HandleLogServiceError(err error) error {
	return service.ErrLogUnexpectedError
}
