package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandleErrorCarrierService(err error) error {
	switch {
	case errors.Is(err, repository.ErrCarrierNotFound):
		return service.ErrCarrierNotFound
	case errors.Is(err, repository.ErrCarrierCidDuplicate):
		return service.ErrCarrierCidDuplicate
	case errors.Is(err, repository.ErrCarrierLocalityId):
		return service.ErrCarrierLocalityIdNotFound
	case errors.Is(err, repository.ErrCarrierGetUpdatedOrCreatedItem):
		return service.ErrCarrierGetUpdatedOrCreatedItem
	case errors.Is(err, repository.ErrCarrierFKConstraintFail):
		return service.ErrCarrierFKConstraintFail
	default:
		return service.ErrCarrierServiceDefault
	}
}
