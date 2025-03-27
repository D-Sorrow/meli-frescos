package error_management

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

func HandleErrorSellerService(err error) error {
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrSellerNotFound):
			return service.ErrSellerNotFound
		case errors.Is(err, repository.ErrSellerAlreadyExists):
			return service.ErrSellerAlreadyExists
		default:
			return service.ErrSellerServiceGeneric
		}
	}
	return nil
}
