package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/validation"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type BuyerHandler struct {
	service service.BuyerService
}

func NewBuyerHandler(service service.BuyerService) *BuyerHandler {
	return &BuyerHandler{service: service}
}

func (b *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buyers, getAllErr := b.service.GetAll()
		if getAllErr != nil {
			if errors.Is(getAllErr, service.NoRegisteredBuyersYet) {
				getAllErr = handler_errors.HandlerError{
					Code: http.StatusOK,
					Msg:  service.NoRegisteredBuyersYet.Error(),
				}
			}

			handler_errors.HandlerResponseError(getAllErr, &w)
			return
		}

		data := make([]dto.BuyerDTO, 0)
		for _, value := range buyers {
			data = append(data, *mappers.BuyerToBuyerDTO(&value))
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get all buyers successful",
			Data: data,
		})
	}
}
func (b *BuyerHandler) GetById() http.HandlerFunc {
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

		buyer, getByIdErr := b.service.GetById(idInt)
		if getByIdErr != nil {
			if errors.Is(getByIdErr, service.BuyerDoesNotExist) {
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
			Msg:  "Get buyer by ID successful",
			Data: mappers.BuyerToBuyerDTO(&buyer),
		})
	}
}

func (b *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buyerCreateDTO dto.BuyerCreateDTO

		if jsonErr := json.NewDecoder(r.Body).Decode(&buyerCreateDTO); jsonErr != nil {
			response.JSON(w, http.StatusBadRequest,
				dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  handler_errors.InvalidJSON,
				})
			return
		}

		validator := validation.BuyerValidator()

		if err := validator.Struct(&buyerCreateDTO); err != nil {
			errs := validation.MapValidatorErrors(err, buyerCreateDTO)
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  handler_errors.InvalidBuyerCreate,
				Data: errs,
			})
			return
		}

		newBuyer, createErr := b.service.Create(*mappers.BuyerCreateDTOToBuyerAttributes(&buyerCreateDTO))
		if createErr != nil {
			if errors.Is(createErr, service.BuyerAlreadyExists) {
				createErr = handler_errors.HandlerError{
					Code: http.StatusConflict,
					Msg:  fmt.Sprintf(createErr.Error(), *buyerCreateDTO.CardNumberID),
				}
			}

			handler_errors.HandlerResponseError(createErr, &w)
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Create buyer successful",
			Data: mappers.BuyerToBuyerDTO(&newBuyer),
		})
	}
}

func (b *BuyerHandler) Patch() http.HandlerFunc {
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

		var buyerPatchDTO dto.BuyerPatchDTO

		if jsonErr := json.NewDecoder(r.Body).Decode(&buyerPatchDTO); jsonErr != nil {
			response.JSON(w, http.StatusBadRequest,
				dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  handler_errors.InvalidJSON,
				})
			return
		}

		validator := validation.BuyerValidator()

		if err := validator.Struct(&buyerPatchDTO); err != nil {
			errs := validation.MapValidatorErrors(err, buyerPatchDTO)
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  handler_errors.InvalidBuyerPatch,
				Data: errs,
			})
			return
		}

		updatedBuyer, updatedErr := b.service.Patch(idInt, *mappers.BuyerPatchDTOToBuyerPatchAttributes(&buyerPatchDTO))
		if updatedErr != nil {
			if errors.Is(updatedErr, service.BuyerAlreadyExists) {
				updatedErr = handler_errors.HandlerError{
					Code: http.StatusConflict,
					Msg:  fmt.Sprintf(updatedErr.Error(), *buyerPatchDTO.CardNumberID),
				}
			} else if errors.Is(updatedErr, service.BuyerDoesNotExist) {
				updatedErr = handler_errors.HandlerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(updatedErr.Error(), idInt),
				}
			}

			handler_errors.HandlerResponseError(updatedErr, &w)
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Update buyer successful",
			Data: mappers.BuyerToBuyerDTO(&updatedBuyer),
		})
	}
}

func (b *BuyerHandler) Delete() http.HandlerFunc {
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

		deleteErr := b.service.Delete(idInt)
		if deleteErr != nil {
			if errors.Is(deleteErr, service.BuyerDoesNotExist) {
				deleteErr = handler_errors.HandlerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(deleteErr.Error(), idInt),
				}
			} else if errors.Is(deleteErr, service.CannotDeleteBuyerWithOrders) {
				deleteErr = handler_errors.HandlerError{
					Code: http.StatusConflict,
					Msg:  deleteErr.Error(),
				}
			}

			handler_errors.HandlerResponseError(deleteErr, &w)
			return
		}

		response.JSON(w, http.StatusNoContent, dto.ResponseDTO{
			Code: http.StatusNoContent,
			Msg:  "Delete buyer successful",
		})
	}
}
func (b *BuyerHandler) GetReportPurchaseOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var idP *int = nil
		id := r.URL.Query().Get("id")

		if id != "" {
			idInt, idErr := strconv.Atoi(id)
			if idErr != nil {
				response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  handler_errors.InvalidID,
				})
				return
			}

			idP = &idInt
		}

		report, getReportErr := b.service.GetReportPurchaseOrders(idP)
		if getReportErr != nil {
			if errors.Is(getReportErr, service.BuyerHasNoOrders) {
				getReportErr = handler_errors.HandlerError{
					Code: http.StatusOK,
					Msg:  service.BuyerHasNoOrders.Error(),
				}
			} else if errors.Is(getReportErr, service.BuyerDoesNotExist) {
				getReportErr = handler_errors.HandlerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(getReportErr.Error(), *idP),
				}
			}

			handler_errors.HandlerResponseError(getReportErr, &w)
			return
		}

		data := make([]dto.ReportPurchaseOrdersDTO, 0)
		for _, value := range report {
			data = append(data, *mappers.ReportPurchaseOrdersToReportPurchaseOrdersDTO(&value))
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get all orders",
			Data: data,
		})
	}
}
