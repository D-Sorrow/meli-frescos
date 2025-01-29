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
			handler_errors.ResponseError(err, w)
			return
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mapSeller,
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
			handler_errors.ResponseError(err, w)
			return
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
			handler_errors.ResponseError(err, w)
			return
		}

		if err := hand.validate.Struct(sellerDto); err != nil {
			handler_errors.ResponseError(err, w)
			return
		}

		seller, err := hand.service.CreateSeller(mappers.MapperToSeller(sellerDto))
		if err != nil {
			handler_errors.ResponseError(err, w)
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "User Created",
			Data: mappers.MapperToSellerDTO(seller),
		})

	}
}
