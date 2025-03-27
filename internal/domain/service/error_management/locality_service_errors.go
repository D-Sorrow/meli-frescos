package error_management

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

func HandleErrorLocalitiesService(err error) error {
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrLocalityNotFound):
			return service.ErrLocalityNotFound
		case errors.Is(err, repository.ErrLocalityAlreadyExists):
			return service.ErrLocalityAlreadyExists
		case errors.Is(err, repository.ErrGetAllLocalities):
			return service.ErrGetAllLocalities
		case errors.Is(err, repository.ErrProvinceNotFound):
			return service.ErrProvinceNotFound
		default:
			return service.ErrLocalityRepositoryGeneric
		}
	}
	return nil
}
