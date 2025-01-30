package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/validation"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"

	serviceErrors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type WarehouseHandler struct {
	service service.WarehouseServiceInterface
}

func NewWarehouseHandler(sevice service.WarehouseServiceInterface) *WarehouseHandler {
	return &WarehouseHandler{service: sevice}
}

func (wh *WarehouseHandler) GetWarehouses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses := wh.service.GetWarehouses()

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
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "invalid id",
				Data: nil,
			})
			return
		}

		warehouse, err := wh.service.GetWarehouseById(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  err.Error(),
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
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "Bad request",
				Data: nil,
			})
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
			switch {
			case err.Error() == serviceErrors.ErrIdDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrWarehouseCodeDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			}
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
			switch {
			case err.Error() == serviceErrors.ErrIdNotFound().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrWarehouseCodeDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			}
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
			if err.Error() == serviceErrors.InternalServerErr().Error() {
				response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
					Data: nil,
				})
				return
			}
			response.JSON(w, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  err.Error(),
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
