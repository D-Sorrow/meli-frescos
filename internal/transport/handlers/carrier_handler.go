package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
)

type CarryHandler struct {
	service service.CarrierServiceInterface
}

func NewCarryHandler(service service.CarrierServiceInterface) *CarryHandler {
	return &CarryHandler{service: service}
}

func (ch *CarryHandler) GetAllCarriers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		carriers, err := ch.service.GetAllCarriers()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  "Server Error",
				Data: nil,
			})
			return
		}

		carriersDto := mappers.MapperToCarriersDto(carriers)

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Carriers got successfully",
			Data: carriersDto,
		})
	}
}

func (ch *CarryHandler) CreateCarrier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := dto.CarrierDto{}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "Bad request",
				Data: nil,
			})
			return
		}

		if err := reqBody.Validate(); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		newCarrier, err := ch.service.CreateCarrier(mappers.MapperToCarrierModel(reqBody))
		if err != nil {
			handler_err := handler_errors.HandleErrorCarrier(err)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "carrier created successsfully",
			Data: mappers.MapperToCarrierDto(newCarrier),
		})
	}
}
