package handlers

import (
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
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
