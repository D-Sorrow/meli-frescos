package handlers

import (
	"errors"
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
			if errors.Is(getAllErr, service.NoRegisteredOrderStatusesYet) {
				getAllErr = handler_errors.HandlerError{
					Code: http.StatusOK,
					Msg:  service.NoRegisteredOrderStatusesYet.Error(),
				}
			}

			handler_errors.HandlerResponseError(getAllErr, &w)
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
