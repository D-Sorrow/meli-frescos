package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	handler_errors "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type CarryHandler struct {
	service service.CarrierServiceInterface
}

func NewCarryHandler(service service.CarrierServiceInterface) *CarryHandler {
	return &CarryHandler{service: service}
}

func (ch *CarryHandler) GetAllCarriers(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		carriers, err := ch.service.GetAllCarriers()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  "Server Error",
				Data: nil,
			})
			return nil
		}

		carriersDto := mappers.MapperToCarriersDto(carriers)

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Carriers got successfully",
			Data: carriersDto,
		})

		return nil
	}
}

func (ch *CarryHandler) CreateCarrier(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqBody := dto.CarrierDto{}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "Bad request",
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

		newCarrier, err := ch.service.CreateCarrier(mappers.MapperToCarrierModel(reqBody))
		if err != nil {
			handler_err := handler_errors.HandleErrorCarrier(err, ctx)
			response.JSON(w, handler_err.Code, dto.ResponseDTO{
				Code: handler_err.Code,
				Msg:  handler_err.Message,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "carrier created successsfully",
			Data: mappers.MapperToCarrierDto(newCarrier),
		})

		return nil
	}
}
