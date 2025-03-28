package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
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
