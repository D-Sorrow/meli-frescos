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
	"github.com/go-playground/validator/v10"
)

type LocalityHandler struct {
	service  service.LocalityService
	validate *validator.Validate
}

func NewLocalityHandler(service service.LocalityService) *LocalityHandler {
	return &LocalityHandler{service: service, validate: validator.New()}
}

func (hand LocalityHandler) CreateLocality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var localityDto dto.LocalityDto

		if err := json.NewDecoder(r.Body).Decode(&localityDto); err != nil {
			handler_errors.ResponseErrorLocality(err, w)
			return
		}

		if err := hand.validate.Struct(localityDto); err != nil {
			handler_errors.ResponseErrorLocality(err, w)
			return
		}

		locality, err := hand.service.CreateLocality(mappers.MapperToLocality(localityDto))
		if err != nil {
			if errors.Is(err, service_errors.ErrLocalityAlreadyExists) {
				handler_errors.ResponseErrorLocality(handler_errors.ErrLocalityAlreadyExists, w)
				return
			}
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Locality Created",
			Data: mappers.MapperToLocalityDTO(locality),
		})

	}
}

func (hand LocalityHandler) GetSellersByLocality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return
		}
		localitySellers, err := hand.service.GetSellersByLocality(id)
		if err != nil {
			if errors.Is(err, service_errors.ErrNotFound) {
				handler_errors.ResponseErrorLocality(handler_errors.ErrNotFound, w)
				return
			}
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToLocalitySellersDTO(localitySellers),
		})
	}
}
