package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
)

type InboundOrderHandler struct {
	service   service.InboundOrderService
	validator *validator.Validate
}

func NewInboundOrderHandler(service service.InboundOrderService) *InboundOrderHandler {
	return &InboundOrderHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (handler *InboundOrderHandler) CreateInboundOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inboundOrderToCreate dto.InboundOrderRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&inboundOrderToCreate); err != nil {
			handler.handleError(w, error_management.ErrInboundOrderBodyDecoding)
			return
		}

		if err := handler.validator.Struct(inboundOrderToCreate); err != nil {
			handler.handleError(w, err)
			return
		}

		inboundOrderModel := mappers.InboundOrderRequestDTOToModel(inboundOrderToCreate)
		err := handler.service.CreateInboundOrder(inboundOrderModel)

		if err != nil {
			handler.handleError(w, err)
			return
		}

		inboundOrderResponseDto := mappers.InboundOrderModelToResponseDTO(*inboundOrderModel)
		handler.respondWithJSON(w, http.StatusCreated, "Success", inboundOrderResponseDto)
	}
}

func (handler *InboundOrderHandler) respondWithJSON(w http.ResponseWriter, code int, msg string, data interface{}) {
	response.JSON(w, code, dto.ResponseDTO{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (handler *InboundOrderHandler) handleError(w http.ResponseWriter, err error) {
	inboundOrderError := error_management.HandleErrorInboundOrder(err)
	handler.respondWithJSON(w, inboundOrderError.Code, inboundOrderError.Message, nil)
}
