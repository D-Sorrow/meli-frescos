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
	serviceErrors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
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
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  "Server Error",
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
				Msg:  "id not found",
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
			switch {
			case err.Error() == serviceErrors.ErrIdDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  "id already exists",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrWarehouseCodeDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  "warehouse code already exists",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrEntityId().Error():
				response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  "entity id faild",
					Data: nil,
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
					Code: http.StatusInternalServerError,
					Msg:  "internal server error",
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
				response.JSON(w, http.StatusNotFound, dto.ResponseDTO{
					Code: http.StatusNotFound,
					Msg:  "id not found",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrWarehouseCodeDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  "warehouse code already exists",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrUpdateBySameData().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  "enter different data to update",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrEntityId().Error():
				response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  "entity id faild",
					Data: nil,
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
					Code: http.StatusInternalServerError,
					Msg:  "internal server error",
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
			switch {
			case err.Error() == serviceErrors.ErrIdNotFound().Error():
				response.JSON(w, http.StatusNotFound, dto.ResponseDTO{
					Code: http.StatusNotFound,
					Msg:  "id not found",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrFKConstraintFail().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  fmt.Sprintf("This warehouse with id %d is use for other entities", id),
					Data: nil,
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
					Code: http.StatusInternalServerError,
					Msg:  "internal server error",
					Data: nil,
				})
				return
			}
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  fmt.Sprintf("Werehouse with id %d deleted", id),
			Data: nil,
		})
	}
}
