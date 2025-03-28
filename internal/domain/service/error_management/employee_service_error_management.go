package error_management

import (
	"errors"
	"strconv"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandleErrorEmployeeService(err error) error {
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrEmployeeNotFound):
			return service.ErrEmployeeNotFound
		case errors.Is(err, strconv.ErrSyntax):
			return service.ErrEmployeeDecodingError
		default:
			return service.ErrEmployeeServiceDefault
		}
	}
	return nil
}
