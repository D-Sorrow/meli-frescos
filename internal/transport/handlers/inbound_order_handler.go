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
			inboundOrderError := error_management.HandleErrorInboundOrder(error_management.ErrInboundOrderBodyDecoding)
			response.JSON(w, inboundOrderError.Code, inboundOrderError.Message)
			return
		}

		if err := handler.validator.Struct(inboundOrderToCreate); err != nil {
			inboundOrderError := error_management.HandleErrorInboundOrder(err)
			response.JSON(w, inboundOrderError.Code, inboundOrderError.Message)
			return
		}

		inboundOrderModel := mappers.InboundOrderRequestDTOToModel(inboundOrderToCreate)
		err := handler.service.CreateInboundOrder(inboundOrderModel)

		if err != nil {
			inboundOrderError := error_management.HandleErrorInboundOrder(err)
			response.JSON(w, inboundOrderError.Code, inboundOrderError.Message)
			return
		}

		inboundOrderResponseDto := mappers.InboundOrderModelToResponseDTO(*inboundOrderModel)
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": inboundOrderResponseDto,
		})
	}
}
