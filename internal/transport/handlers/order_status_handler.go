package handlers

import (
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
)

type OrderStatusHandler struct {
	service service.OrderStatusService
}

func NewOrderStatusHandler(service service.OrderStatusService) *OrderStatusHandler {
	return &OrderStatusHandler{service: service}
}

func (b *OrderStatusHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderStatus, getAllErr := b.service.GetAll()
		if getAllErr != nil {
			getAllErr = handler_errors.HandleHandlerError(getAllErr)
			handlerErr := handler_errors.HandleOrderStatusHandlerError(getAllErr, nil, nil)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
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
	}
}
