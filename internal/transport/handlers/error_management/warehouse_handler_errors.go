package error_management

import (
	"context"
	"errors"
	"net/http"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/middlewares"
)

const (
	msgWarehouseInternalError          = "internal server error"
	msgWarehouseNotFound               = "warehouse not found"
	msgWarehouseIdDuplicated           = "warehouse id already exists"
	msgWarehouseCodeDuplicated         = "warehouse code already exists"
	msgWarehouseLocalityIdNotFound     = "locality id not found"
	msgWarehouseUpdateBySameData       = "enter different data to update"
	msgWarehouseGetUpdatedOrCreateItem = "the warehouse was created or updated but could not be displayed"
	msgWarehouseFKConstraintFail       = "This warehouse cannot be deleted, another entity is using it"
)

var (
	ErrWarehouseIdNotValid = errors.New("invalid id")
)

type HandlerErrorWarehouse struct {
	Code    int
	Message string
}

var warehouseHandlerErrors = map[error]HandlerErrorWarehouse{
	service.ErrWarehouseServiceDefault: {
		Code:    http.StatusInternalServerError,
		Message: msgWarehouseInternalError,
	},
	service.ErrWarehouseNotFound: {
		Code:    http.StatusNotFound,
		Message: msgWarehouseNotFound,
	},
	service.ErrWarehouseIdDuplicate: {
		Code:    http.StatusConflict,
		Message: msgWarehouseIdDuplicated,
	},
	service.ErrWarehouseCodeDuplicate: {
		Code:    http.StatusConflict,
		Message: msgWarehouseCodeDuplicated,
	},
	service.ErrWarehouseLocalityIdNotFound: {
		Code:    http.StatusBadRequest,
		Message: msgWarehouseLocalityIdNotFound,
	},
	service.ErrWarehouseUpdateBySameData: {
		Code:    http.StatusConflict,
		Message: msgWarehouseUpdateBySameData,
	},
	service.ErrWarehouseUpdateBySameData: {
		Code:    http.StatusConflict,
		Message: msgWarehouseUpdateBySameData,
	},
	service.ErrWarehouseGetUpdatedOrCreatedItem: {
		Code:    http.StatusInternalServerError,
		Message: msgWarehouseGetUpdatedOrCreateItem,
	},
	service.ErrWarehouseFKConstraintFail: {
		Code:    http.StatusConflict,
		Message: msgWarehouseFKConstraintFail,
	},
	ErrWarehouseIdNotValid: {
		Code:    http.StatusBadRequest,
		Message: ErrWarehouseIdNotValid.Error(),
	},
}

func HandleErrorWarehouse(err error, ctx *context.Context) HandlerErrorWarehouse {
	*ctx = context.WithValue(*ctx, middlewares.AppErrorKey, err)

	switch {
	case errors.Is(err, service.ErrWarehouseNotFound):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrWarehouseIdDuplicate):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrWarehouseCodeDuplicate):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrWarehouseLocalityIdNotFound):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrWarehouseUpdateBySameData):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrWarehouseGetUpdatedOrCreatedItem):
		return warehouseHandlerErrors[err]
	case errors.Is(err, service.ErrWarehouseFKConstraintFail):
		return warehouseHandlerErrors[err]
	case errors.Is(err, ErrWarehouseIdNotValid):
		return warehouseHandlerErrors[err]
	default:
		return warehouseHandlerErrors[service.ErrWarehouseServiceDefault]
	}
}
