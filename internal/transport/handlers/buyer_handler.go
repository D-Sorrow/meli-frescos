package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	service_mappers "github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
	"github.com/D-Sorrow/meli-frescos/internal/domain/validation"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	handler_mappers "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
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
			if errors.Is(getAllErr, service_errors.NoRegisteredBuyersYet) {
				getAllErr = handler_errors.BuyerError{
					Code: http.StatusOK,
					Msg:  service_errors.NoRegisteredBuyersYet.Error(),
				}
			}

			handler_errors.BuyerResponseError(getAllErr, &w)
			return
		}

		data := make([]dto.BuyerDTO, 0)
		for _, value := range buyers {
			data = append(data, *service_mappers.BuyerToBuyerDTO(&value))
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
			if errors.Is(getByIdErr, service_errors.BuyerDoesNotExist) {
				getByIdErr = handler_errors.BuyerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(getByIdErr.Error(), idInt),
				}
			}

			handler_errors.BuyerResponseError(getByIdErr, &w)
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get buyer by ID successful",
			Data: service_mappers.BuyerToBuyerDTO(&buyer),
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

		newBuyer, createErr := b.service.Create(*handler_mappers.BuyerCreateDTOToBuyerAttributes(&buyerCreateDTO))
		if createErr != nil {
			if errors.Is(createErr, service_errors.BuyerAlreadyExists) {
				createErr = handler_errors.BuyerError{
					Code: http.StatusConflict,
					Msg:  fmt.Sprintf(createErr.Error(), newBuyer.CardNumberID),
				}
			}

			handler_errors.BuyerResponseError(createErr, &w)
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Create buyer successful",
			Data: service_mappers.BuyerToBuyerDTO(&newBuyer),
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

		updatedBuyer, updatedErr := b.service.Patch(idInt, *handler_mappers.BuyerPatchDTOToBuyerPatchAttributes(&buyerPatchDTO))
		if updatedErr != nil {
			if errors.Is(updatedErr, service_errors.BuyerAlreadyExists) {
				updatedErr = handler_errors.BuyerError{
					Code: http.StatusConflict,
					Msg:  fmt.Sprintf(updatedErr.Error(), updatedBuyer.CardNumberID),
				}
			} else if errors.Is(updatedErr, service_errors.BuyerDoesNotExist) {
				updatedErr = handler_errors.BuyerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(updatedErr.Error(), idInt),
				}
			}

			handler_errors.BuyerResponseError(updatedErr, &w)
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Update buyer successful",
			Data: service_mappers.BuyerToBuyerDTO(&updatedBuyer),
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
			if errors.Is(deleteErr, service_errors.BuyerDoesNotExist) {
				deleteErr = handler_errors.BuyerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(deleteErr.Error(), idInt),
				}
			}

			handler_errors.BuyerResponseError(deleteErr, &w)
			return
		}

		response.JSON(w, http.StatusNoContent, dto.ResponseDTO{
			Code: http.StatusNoContent,
			Msg:  "Delete buyer successful",
		})
	}
}
