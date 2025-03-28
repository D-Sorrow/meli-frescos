package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
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
