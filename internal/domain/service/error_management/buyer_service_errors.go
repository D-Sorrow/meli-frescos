package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandleBuyerServiceError(err error) error {
	switch {
	case errors.Is(err, repository.ErrBuyerNoRegisteredBuyersYet):
		return service.ErrBuyerNoRegisteredBuyersYet
	case errors.Is(err, repository.ErrBuyerDuplicateCardNumberID):
		return service.ErrBuyerAlreadyExists
	case errors.Is(err, repository.ErrBuyerNotFoundWithID):
		return service.ErrBuyerDoesNotExist
	case errors.Is(err, repository.ErrBuyerCannotDeleteBuyerWithOrders):
		return service.ErrBuyerCannotDeleteBuyerWithOrders
	case errors.Is(err, repository.ErrBuyerHasNoOrders):
		return service.ErrBuyerHasNoOrders
	default:
		return service.ErrBuyerUnexpectedError
	}
}
