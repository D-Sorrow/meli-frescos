package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type PurchaseOrderHandler struct {
	service service.PurchaseOrderService
}

func NewPurchaseOrderHandler(service service.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{service: service}
}

func (b *PurchaseOrderHandler) GetById(ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, idErr := strconv.Atoi(id)
		if idErr != nil {
			handler_errors.HandleHandlerError(idErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				handler_errors.ErrPurchaseOrderInvalidID,
				nil,
				nil,
				ctx,
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
		}

		purchaseOrder, getByIdErr := b.service.GetById(idInt)
		if getByIdErr != nil {
			getByIdErr = handler_errors.HandleHandlerError(getByIdErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				getByIdErr,
				nil,
				map[string]interface{}{
					"ID": idInt,
				},
				ctx,
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get purchase order by ID successful",
			Data: mappers.PurchaseOrderToPurchaseOrderDTO(&purchaseOrder),
		})
	}
}

func (b *PurchaseOrderHandler) Create(ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var purchaseOrderCreateDTO dto.PurchaseOrderCreateDTO

		if jsonErr := json.NewDecoder(r.Body).Decode(&purchaseOrderCreateDTO); jsonErr != nil {
			handler_errors.HandleHandlerError(jsonErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				handler_errors.ErrPurchaseOrderInvalidJSON,
				nil,
				nil,
				ctx,
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
		}

		newPurchaseOrder, createErr := b.service.Create(
			*mappers.PurchaseOrderCreateDTOToPurchaseOrderAttributesFKs(&purchaseOrderCreateDTO),
			time.Now().UTC(),
			uuid.New().String(),
		)
		if createErr != nil {
			createErr = handler_errors.HandleHandlerError(createErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				createErr,
				nil,
				nil,
				ctx,
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Create purchase order successful",
			Data: mappers.PurchaseOrderToPurchaseOrderDTO(&newPurchaseOrder),
		})
	}
}
