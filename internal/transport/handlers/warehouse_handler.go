package handlers

import (
	// "encoding/json"
	// "fmt"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/validation"

	// "github.com/D-Sorrow/meli-frescos/internal/domain/validation"

	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type WarehouseHandler struct {
	service service.WarehouseServiceInterface
}

func NewWarehouseHandler(sevice service.WarehouseServiceInterface) *WarehouseHandler {
	return &WarehouseHandler{service: sevice}
}

func (wh *WarehouseHandler) GetWarehouses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := wh.service.GetWarehouses()
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}

		warehousesDto := mappers.MapperToWarehousesDto(warehouses)

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Warehouses got successfully",
			Data: warehousesDto,
		})
	}
}

func (wh *WarehouseHandler) GetWarehouseById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(handler_errors.ErrWarehouseIdNotValid)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}

		warehouse, err := wh.service.GetWarehouseById(id)
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}

		warehouseJSON := mappers.MapperToWarehouseDto(warehouse)

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Warehouse got successfully",
			Data: warehouseJSON,
		})
	}
}

func (wh *WarehouseHandler) CreateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := dto.WarehouseDto{}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  "Unprocessable Entity",
				Data: nil,
			})
			return
		}

		if err := reqBody.Validate(); err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		newWarehouse, err := wh.service.CreateWarehouse(mappers.MapperToWarehouseModel(reqBody))
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}
		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Warehouse created successsfully",
			Data: newWarehouse,
		})
	}
}

func (wh *WarehouseHandler) PatchWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "invalid id",
				Data: nil,
			})
			return
		}

		jsonBody, err := validation.ValidatePatchRequestBody(r)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: jsonBody,
			})
			return
		}
		err = validation.ValidatePatchValues(jsonBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		warehouse, err := wh.service.PatchWarehouse(id, jsonBody)
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Warehouse updated",
			Data: warehouse,
		})
	}
}

func (wh *WarehouseHandler) DeleteWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "invalid id",
				Data: nil,
			})
			return
		}

		err = wh.service.DeleteWarehouse(id)
		if err != nil {
			handler_err := handler_errors.HandleErrorWarehouse(err)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  fmt.Sprintf("Werehouse with id %d deleted", id),
			Data: nil,
		})
	}
}
