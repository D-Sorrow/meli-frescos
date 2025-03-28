package handlers

import (
	"context"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	handler_errors "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type OrderStatusHandler struct {
	service service.OrderStatusService
}

func NewOrderStatusHandler(service service.OrderStatusService) *OrderStatusHandler {
	return &OrderStatusHandler{service: service}
}

func (b *OrderStatusHandler) GetAll(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		orderStatus, getAllErr := b.service.GetAll()
		if getAllErr != nil {
			getAllErr = handler_errors.HandleHandlerError(getAllErr)
			handlerErr := handler_errors.HandleOrderStatusHandlerError(getAllErr, nil, nil, ctx)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return nil
		}

		data := make([]dto.OrderStatusDTO, 0)
		for _, value := range orderStatus {
			data = append(data, *mappers.OrderStatusToOrderStatusDTO(&value))
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get all order statuses successful",
			Data: data,
		})

		return nil
	}
}
