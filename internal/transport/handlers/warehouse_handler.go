package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/validation"
	"github.com/melisource/fury_go-core/pkg/web"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	handler_errors "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
)

type WarehouseHandler struct {
	service service.WarehouseServiceInterface
}

func NewWarehouseHandler(sevice service.WarehouseServiceInterface) *WarehouseHandler {
	return &WarehouseHandler{service: sevice}
}

func (wh *WarehouseHandler) GetWarehouses(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		warehouses, err := wh.service.GetWarehouses()
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err, ctx)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}

		warehousesDto := mappers.MapperToWarehousesDto(warehouses)

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Warehouses got successfully",
			Data: warehousesDto,
		})

		return nil
	}
}

func (wh *WarehouseHandler) GetWarehouseById(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(
				handler_errors.ErrWarehouseIdNotValid, ctx,
			)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}

		warehouse, err := wh.service.GetWarehouseById(id)
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err, ctx)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}

		warehouseJSON := mappers.MapperToWarehouseDto(warehouse)

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Warehouse got successfully",
			Data: warehouseJSON,
		})

		return nil
	}
}

func (wh *WarehouseHandler) CreateWarehouse(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqBody := dto.WarehouseDto{}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  "Unprocessable Entity",
				Data: nil,
			})
			return nil
		}

		if err := reqBody.Validate(); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}

		newWarehouse, err := wh.service.CreateWarehouse(mappers.MapperToWarehouseModel(reqBody))
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err, ctx)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}
		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Warehouse created successsfully",
			Data: mappers.MapperToWarehouseDto(newWarehouse),
		})

		return nil
	}
}

func (wh *WarehouseHandler) PatchWarehouse(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "invalid id",
				Data: nil,
			})
			return nil
		}

		jsonBody, err := validation.ValidatePatchRequestBody(r)
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: jsonBody,
			})
			return nil
		}
		err = validation.ValidatePatchValues(jsonBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}

		warehouse, err := wh.service.PatchWarehouse(id, jsonBody)
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err, ctx)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Warehouse updated",
			Data: mappers.MapperToWarehouseDto(warehouse),
		})

		return nil
	}
}

func (wh *WarehouseHandler) DeleteWarehouse(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "invalid id",
				Data: nil,
			})
			return nil
		}

		err = wh.service.DeleteWarehouse(id)
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err, ctx)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  fmt.Sprintf("Werehouse with id %d deleted", id),
			Data: nil,
		})

		return nil
	}
}
