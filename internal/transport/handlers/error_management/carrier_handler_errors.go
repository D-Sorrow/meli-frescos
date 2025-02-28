package error_management

import (
	"errors"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
)

const (
	msgCarrierInternalError          = "internal server error"
	msgCarrierNotFound               = "carrier not found"
	msgCarrierIdDuplicated           = "carrier id already exists"
	msgCarrierCidDuplicated          = "carrier code already exists"
	msgCarrierLocalityIdNotFound     = "locality id not found"
	msgCarrierGetUpdatedOrCreateItem = "the carrier was created or updated but could not be displayed"
	msgCarrierFKConstraintFail       = "This carrier cannot be deleted, another entity is using it"
)

type HandlerErrorWarehouse struct {
	Code    int
	Message string
}

var warehouseHandlerErrors = map[error]HandlerErrorWarehouse{
	service.ErrCarrierServiceDefault: {
		Code:    http.StatusInternalServerError,
		Message: msgCarrierInternalError,
	},
	service.ErrCarrierNotFound: {
		Code:    http.StatusNotFound,
		Message: msgCarrierNotFound,
	},
	service.ErrCarrierCidDuplicate: {
		Code:    http.StatusConflict,
		Message: msgCarrierCidDuplicated,
	},
	service.ErrCarrierLocalityIdNotFound: {
		Code:    http.StatusBadRequest,
		Message: msgCarrierLocalityIdNotFound,
	},
	service.ErrCarrierGetUpdatedOrCreatedItem: {
		Code:    http.StatusInternalServerError,
		Message: msgCarrierGetUpdatedOrCreateItem,
	},
	service.ErrCarrierFKConstraintFail: {
		Code:    http.StatusConflict,
		Message: msgCarrierFKConstraintFail,
	},
}

func HandleErrorWarehouse(err error) HandlerErrorWarehouse {
	switch {
	case errors.Is(err, service.ErrCarrierNotFound):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierCidDuplicate):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierLocalityIdNotFound):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierGetUpdatedOrCreatedItem):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierFKConstraintFail):
		return warehouseHandlerErrors[err]
	default:
		return warehouseHandlerErrors[service.ErrCarrierServiceDefault]
	}
}
