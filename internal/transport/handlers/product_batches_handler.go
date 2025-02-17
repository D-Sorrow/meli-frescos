package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	serr "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	herr "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ProductBatchesHandler struct {
	s        service.ProductBatchesRepository
	validate *validator.Validate
}

func NewProductBatches(s service.ProductBatchesRepository) *ProductBatchesHandler {
	return &ProductBatchesHandler{s: s, validate: validator.New()}
}

func (h ProductBatchesHandler) AddProductBatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ProductBatchesDto dto.ProductBatchesDto

		if err := json.NewDecoder(r.Body).Decode(&ProductBatchesDto); err != nil {
			herr.ResponseErrorProductBatches(err, w)
			return
		}

		if err := h.validate.Struct(ProductBatchesDto); err != nil {
			herr.ResponseErrorProductBatches(err, w)
			return
		}

		product, err := h.s.AddProductBatches(mappers.MapperToProductBatches(ProductBatchesDto))
		if err != nil {
			if errors.Is(err, serr.ErrProductBatchesAlredyExists) {
				herr.ResponseErrorProductBatches(herr.ErrLocalityAlreadyExists, w)
				return
			}
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "ProductBatches Created",
			Data: mappers.MapperToProductBatchesDTO(product),
		})

	}
}

func (h *ProductBatchesHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, idErr := strconv.Atoi(id)
		if idErr != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  herr.InvalidID,
			})
			return
		}

		productBatch, getByIdErr := h.s.GetById(idInt)
		if getByIdErr != nil {
			if errors.Is(getByIdErr, service.ProductBatchOrderDoesNotExist) {
				getByIdErr = herr.HandlerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(getByIdErr.Error(), idInt),
				}
			}

			herr.HandlerResponseError(getByIdErr, &w)
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get product by ID successful",
			Data: mappers.MapperToProductBatchesDTO2(&productBatch),
		})
	}
}

func (h *ProductBatchesHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ProductBatchesDto dto.ProductBatchesDtoReport

		if jsonErr := json.NewDecoder(r.Body).Decode(&ProductBatchesDto); jsonErr != nil {
			response.JSON(w, http.StatusBadRequest,
				dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  herr.InvalidJSON,
				})
			return
		}

		ProductBatches, createErr := h.s.Create(*mappers.ProductBatchesCreateDTOToPProductBatchesFKs(&ProductBatchesDto))
		if createErr != nil {
			if errors.Is(createErr, service.ForeignKeysNotValidProductBatches) {
				createErr = herr.HandlerError{
					Code: http.StatusConflict,
					Msg:  createErr.Error(),
				}
			}

			herr.HandlerResponseError(createErr, &w)
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Create ProductBatch order successful",
			Data: mappers.MapperToProductBatchesDTO2(&ProductBatches),
		})
	}
}
