package error_management

import (
	"errors"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
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
