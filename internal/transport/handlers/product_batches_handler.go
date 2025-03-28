package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	serr "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	herr "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type ProductBatchesHandler struct {
	s        service.ProductBatchesRepository
	validate *validator.Validate
}

func NewProductBatches(s service.ProductBatchesRepository) *ProductBatchesHandler {
	return &ProductBatchesHandler{s: s, validate: validator.New()}
}

func (h ProductBatchesHandler) AddProductBatches(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var ProductBatchesDto dto.ProductBatchesDto

		if err := json.NewDecoder(r.Body).Decode(&ProductBatchesDto); err != nil {
			herr.ResponseErrorProductBatches(err, w, ctx)
			return nil
		}

		if err := h.validate.Struct(ProductBatchesDto); err != nil {
			herr.ResponseErrorProductBatches(err, w, ctx)
			return nil
		}

		product, err := h.s.AddProductBatches(mappers.MapperToProductBatches(ProductBatchesDto))
		if err != nil {
			if errors.Is(err, serr.ErrProductBatchesAlredyExists) {
				herr.ResponseErrorProductBatches(herr.LocalityAlreadyExists, w, ctx)
				return nil
			}
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "ProductBatches Created",
			Data: mappers.MapperToProductBatchesDTO(product),
		})

		return nil
	}
}

func (h *ProductBatchesHandler) GetById(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := chi.URLParam(r, "id")
		idInt, idErr := strconv.Atoi(id)
		if idErr != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  herr.InvalidID,
			})
			return nil
		}

		productBatch, getByIdErr := h.s.GetById(idInt)
		if getByIdErr != nil {
			if errors.Is(getByIdErr, service.ErrProductBatchOrderDoesNotExist) {
				getByIdErr = herr.HandlerError{
					Code: http.StatusNotFound,
					Msg:  fmt.Sprintf(getByIdErr.Error(), idInt),
				}
			}

			herr.HandlerResponseError(getByIdErr, &w, ctx)
			return nil
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Get product by ID successful",
			Data: mappers.MapperToProductBatchesDTO2(&productBatch),
		})

		return nil
	}
}

func (h *ProductBatchesHandler) Create(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var ProductBatchesDto dto.ProductBatchesDtoReport

		if jsonErr := json.NewDecoder(r.Body).Decode(&ProductBatchesDto); jsonErr != nil {
			response.JSON(w, http.StatusBadRequest,
				dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  herr.InvalidJSON,
				})
			return nil
		}

		ProductBatches, createErr := h.s.Create(
			*mappers.ProductBatchesCreateDTOToPProductBatchesFKs(&ProductBatchesDto),
		)
		if createErr != nil {
			if errors.Is(createErr, service.ErrForeignKeysNotValidProductBatches) {
				createErr = herr.HandlerError{
					Code: http.StatusConflict,
					Msg:  createErr.Error(),
				}
			}

			herr.HandlerResponseError(createErr, &w, ctx)
			return nil
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Create ProductBatch order successful",
			Data: mappers.MapperToProductBatchesDTO2(&ProductBatches),
		})

		return nil
	}
}
