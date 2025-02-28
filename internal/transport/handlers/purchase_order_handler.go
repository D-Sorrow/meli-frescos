package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type PurchaseOrderHandler struct {
	service service.PurchaseOrderService
}

func NewPurchaseOrderHandler(service service.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{service: service}
}

func (b *PurchaseOrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, idErr := strconv.Atoi(id)
		if idErr != nil {
			handler_errors.HandleHandlerError(idErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				handler_errors.ErrPurchaseOrderInvalidID,
				nil,
				nil,
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

func (b *PurchaseOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var purchaseOrderCreateDTO dto.PurchaseOrderCreateDTO

		if jsonErr := json.NewDecoder(r.Body).Decode(&purchaseOrderCreateDTO); jsonErr != nil {
			handler_errors.HandleHandlerError(jsonErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				handler_errors.ErrPurchaseOrderInvalidJSON,
				nil,
				nil,
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
		)
		if createErr != nil {
			createErr = handler_errors.HandleHandlerError(createErr)
			handlerErr := handler_errors.HandlePurchaseOrderHandlerError(
				createErr,
				nil,
				nil,
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
