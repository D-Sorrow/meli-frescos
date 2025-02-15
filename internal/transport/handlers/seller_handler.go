package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	service_errors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	handler_errors "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type HandlerSeller struct {
	service  service.SellerService
	validate *validator.Validate
}

func NewHandlerService(service service.SellerService) *HandlerSeller {
	return &HandlerSeller{service: service, validate: validator.New()}
}

func (hand *HandlerSeller) GetSellers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mapSeller, err := hand.service.GetSellers()
		if err != nil {
			if errors.Is(err, service_errors.ErrNotFound) {
				handler_errors.ResponseErrorSeller(handler_errors.ErrNotFound, w)
				return
			}
		}

		data := make([]dto.SellerDto, 0)
		for _, value := range mapSeller {
			data = append(data, mappers.MapperToSellerDTO(value))
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: data,
		})
	}
}

func (hand *HandlerSeller) GetSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return
		}
		seller, err := hand.service.GetSellerById(id)
		if err != nil {
			if errors.Is(err, service_errors.ErrNotFound) {
				handler_errors.ResponseErrorSeller(handler_errors.ErrNotFound, w)
				return
			}
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToSellerDTO(seller),
		})

	}
}

func (hand *HandlerSeller) CreateSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sellerDto dto.SellerDto

		if err := json.NewDecoder(r.Body).Decode(&sellerDto); err != nil {
			handler_errors.ResponseErrorSeller(err, w)
			return
		}

		if err := hand.validate.Struct(sellerDto); err != nil {
			handler_errors.ResponseErrorSeller(err, w)
			return
		}

		seller, err := hand.service.CreateSeller(mappers.MapperToSeller(sellerDto))
		if err != nil {
			if errors.Is(err, service_errors.ErrAlreadyExists) {
				handler_errors.ResponseErrorSeller(handler_errors.ErrAlreadyExists, w)
				return
			}
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "User Created",
			Data: mappers.MapperToSellerDTO(seller),
		})

	}
}

func (hand *HandlerSeller) UpdateSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return
		}

		var sellerDto dto.SellerUpdateDto

		if err := json.NewDecoder(r.Body).Decode(&sellerDto); err != nil {
			handler_errors.ResponseErrorSeller(err, w)
			return
		}

		if err := hand.validate.Struct(sellerDto); err != nil {
			handler_errors.ResponseErrorSeller(err, w)
			return
		}

		seller, err := hand.service.UpdateSeller(id, mappers.MapperToSellerPatch(sellerDto))
		if err != nil {
			if errors.Is(err, service_errors.ErrNotFound) {
				handler_errors.ResponseErrorSeller(handler_errors.ErrNotFound, w)
				return
			}
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "User Updated",
			Data: mappers.MapperToSellerDTO(seller),
		})

	}
}

func (hand *HandlerSeller) DeleteSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return
		}

		err = hand.service.DeleteSeller(id)
		if err != nil {
			if errors.Is(err, service_errors.ErrNotFound) {
				handler_errors.ResponseErrorSeller(handler_errors.ErrNotFound, w)
				return
			}
		}

		response.JSON(w, http.StatusNoContent, dto.ResponseDTO{})

	}
}
