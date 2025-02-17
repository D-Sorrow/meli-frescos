package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  handler_errors.InvalidID,
			})
			return
		}

		purchaseOrder, getByIdErr := b.service.GetById(idInt)
		if getByIdErr != nil {
			if errors.Is(getByIdErr, service.PurchaseOrderDoesNotExist) {
				getByIdErr = handler_errors.HandlerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(getByIdErr.Error(), idInt),
				}
			}

			handler_errors.HandlerResponseError(getByIdErr, &w)
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
			response.JSON(w, http.StatusBadRequest,
				dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  handler_errors.InvalidJSON,
				})
			return
		}

		newPurchaseOrder, createErr := b.service.Create(*mappers.PurchaseOrderCreateDTOToPurchaseOrderAttributesFKs(&purchaseOrderCreateDTO))
		if createErr != nil {
			if errors.Is(createErr, service.ForeignKeysNotValid) {
				createErr = handler_errors.HandlerError{
					Code: http.StatusConflict,
					Msg:  createErr.Error(),
				}
			}

			handler_errors.HandlerResponseError(createErr, &w)
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Create purchase order successful",
			Data: mappers.PurchaseOrderToPurchaseOrderDTO(&newPurchaseOrder),
		})
	}
}
