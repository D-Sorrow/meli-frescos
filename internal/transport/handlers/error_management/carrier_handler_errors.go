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

type HandlerErrorCarrier struct {
	Code    int
	Message string
}

var carrierHandlerErrors = map[error]HandlerErrorWarehouse{
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

func HandleErrorCarrier(err error) HandlerErrorWarehouse {
	switch {
	case errors.Is(err, service.ErrCarrierNotFound):
		return carrierHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierCidDuplicate):
		return carrierHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierLocalityIdNotFound):
		return carrierHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierGetUpdatedOrCreatedItem):
		return carrierHandlerErrors[err]
	case errors.Is(err, service.ErrCarrierFKConstraintFail):
		return carrierHandlerErrors[err]
	default:
		return carrierHandlerErrors[service.ErrCarrierServiceDefault]
	}
}
