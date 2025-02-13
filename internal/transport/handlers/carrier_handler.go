package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"

	serviceErrors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type CarryHandler struct {
	service service.CarrierServiceInterface
}

func NewCarryHandler(service service.CarrierServiceInterface) *CarryHandler {
	return &CarryHandler{service: service}
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
			switch {
			case err.Error() == serviceErrors.ErrIdDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  "id already exists",
					Data: nil,
				})
				return
			case err.Error() == serviceErrors.ErrCarrierCidDuplicate().Error():
				response.JSON(w, http.StatusConflict, dto.ResponseDTO{
					Code: http.StatusConflict,
					Msg:  "carrier cid already exists",
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
			Msg:  "carrier created successsfully",
			Data: newCarrier,
		})
	}
}
