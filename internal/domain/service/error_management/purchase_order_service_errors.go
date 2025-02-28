package error_management

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
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
