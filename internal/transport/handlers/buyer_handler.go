package handlers

import (
	"encoding/json"
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
			getAllErr = handler_errors.HandleHandlerError(getAllErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(getAllErr, nil, nil)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
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
			handler_errors.HandleHandlerError(idErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidID,
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

		buyer, getByIdErr := b.service.GetById(idInt)
		if getByIdErr != nil {
			getByIdErr = handler_errors.HandleHandlerError(getByIdErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
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
			Msg:  "Get buyer by ID successful",
			Data: mappers.BuyerToBuyerDTO(&buyer),
		})
	}
}

func (b *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buyerCreateDTO dto.BuyerCreateDTO

		if jsonErr := json.NewDecoder(r.Body).Decode(&buyerCreateDTO); jsonErr != nil {
			handler_errors.HandleHandlerError(jsonErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidJSON,
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

		validator := validation.BuyerValidator()

		if err := validator.Struct(&buyerCreateDTO); err != nil {
			err = handler_errors.HandleHandlerError(err)
			errs := validation.MapValidatorErrors(err, buyerCreateDTO)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidCreateDTO,
				errs,
				nil,
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
		}

		newBuyer, createErr := b.service.Create(
			*mappers.BuyerCreateDTOToBuyerAttributes(&buyerCreateDTO),
		)
		if createErr != nil {
			createErr = handler_errors.HandleHandlerError(createErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				createErr,
				nil,
				map[string]interface{}{
					"CardNumberID": *buyerCreateDTO.CardNumberID,
				},
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
			handler_errors.HandleHandlerError(idErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidID,
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

		var buyerPatchDTO dto.BuyerPatchDTO

		if jsonErr := json.NewDecoder(r.Body).Decode(&buyerPatchDTO); jsonErr != nil {
			handler_errors.HandleHandlerError(jsonErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidJSON,
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

		validator := validation.BuyerValidator()

		if err := validator.Struct(&buyerPatchDTO); err != nil {
			err = handler_errors.HandleHandlerError(err)
			errs := validation.MapValidatorErrors(err, buyerPatchDTO)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidPatchDTO,
				errs,
				nil,
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
			return
		}

		updatedBuyer, updatedErr := b.service.Patch(
			idInt,
			*mappers.BuyerPatchDTOToBuyerPatchAttributes(&buyerPatchDTO),
		)
		if updatedErr != nil {
			updatedErr = handler_errors.HandleHandlerError(updatedErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				updatedErr,
				nil,
				map[string]interface{}{
					"ID":           idInt,
					"CardNumberID": *buyerPatchDTO.CardNumberID,
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
			handler_errors.HandleHandlerError(idErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				handler_errors.ErrBuyerInvalidID,
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

		deleteErr := b.service.Delete(idInt)
		if deleteErr != nil {
			deleteErr = handler_errors.HandleHandlerError(deleteErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				deleteErr,
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
				handler_errors.HandleHandlerError(idErr)
				handlerErr := handler_errors.HandleBuyerHandlerError(
					handler_errors.ErrBuyerInvalidID,
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

			idP = &idInt
		}

		report, getReportErr := b.service.GetReportPurchaseOrders(idP)
		if getReportErr != nil {
			getReportErr = handler_errors.HandleHandlerError(getReportErr)
			handlerErr := handler_errors.HandleBuyerHandlerError(
				getReportErr,
				nil,
				map[string]interface{}{
					"ID": *idP,
				},
			)
			response.JSON(w, handlerErr.Code, dto.ResponseDTO{
				Code: handlerErr.Code,
				Msg:  handlerErr.Msg,
				Data: handlerErr.Data,
			})
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
