package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandlePurchaseOrderServiceError(err error) error {
	switch {
	case errors.Is(err, repository.ErrPurchaseOrderFKWareHouseIdNotValid):
		return service.ErrPurchaseOrderFKWareHouseIdNotValid
	case errors.Is(err, repository.ErrPurchaseOrderFKBuyerIdNotValid):
		return service.ErrPurchaseOrderFKBuyerIdNotValid
	case errors.Is(err, repository.ErrPurchaseOrderFKOrderStatusIdNotValid):
		return service.ErrPurchaseOrderFKOrderStatusIdNotValid
	case errors.Is(err, repository.ErrPurchaseOrderFKCarrierIdNotValid):
		return service.ErrPurchaseOrderFKCarrierIdNotValid
	case errors.Is(err, repository.ErrPurchaseOrderNotFoundWithID):
		return service.ErrPurchaseOrderDoesNotExist
	default:
		return service.ErrPurchaseOrderUnexpectedError
	}
}
