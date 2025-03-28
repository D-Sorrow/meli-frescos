package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
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

func (handler *InboundOrderHandler) CreateInboundOrder(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var inboundOrderToCreate dto.InboundOrderRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&inboundOrderToCreate); err != nil {
			handler.handleError(w, error_management.ErrInboundOrderBodyDecoding, ctx)
			return nil
		}

		if err := handler.validator.Struct(inboundOrderToCreate); err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		inboundOrderModel := mappers.InboundOrderRequestDTOToModel(inboundOrderToCreate)
		err := handler.service.CreateInboundOrder(inboundOrderModel)

		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		inboundOrderResponseDto := mappers.InboundOrderModelToResponseDTO(*inboundOrderModel)
		handler.respondWithJSON(w, http.StatusCreated, "Success", inboundOrderResponseDto)
		return nil
	}
}

func (handler *InboundOrderHandler) respondWithJSON(
	w http.ResponseWriter,
	code int,
	msg string,
	data interface{},
) {
	response.JSON(w, code, dto.ResponseDTO{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (handler *InboundOrderHandler) handleError(
	w http.ResponseWriter,
	err error,
	ctx *context.Context,
) {
	inboundOrderError := error_management.HandleErrorInboundOrder(err, ctx)
	handler.respondWithJSON(w, inboundOrderError.Code, inboundOrderError.Message, nil)
}
