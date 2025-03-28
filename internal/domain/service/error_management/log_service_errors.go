package error_management

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandleLogServiceError(err error) error {
	return service.ErrLogUnexpectedError
}
