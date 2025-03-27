package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
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

func (hand LocalityHandler) CreateLocality(ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var localityDto dto.LocalityDto

		if err := json.NewDecoder(r.Body).Decode(&localityDto); err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return
		}

		if err := hand.validate.Struct(localityDto); err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return
		}

		locality, err := hand.service.CreateLocality(mappers.MapperToLocality(localityDto))
		if err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Locality Created",
			Data: mappers.MapperToLocalityDTO(locality),
		})

	}
}

func (hand LocalityHandler) GetSellersByLocality(ctx *context.Context) http.HandlerFunc {
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
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToLocalitySellersDTO(localitySellers),
		})
	}
}

func (hand LocalityHandler) GetCarriersByLocality(ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			carriersByLocalities, err := hand.service.GetCarriersByAllLocalities()
			if err != nil {
				localityError := error_management.HandleErrorLocality(err, ctx)
				response.JSON(w, localityError.Code, dto.ResponseDTO{
					Code: localityError.Code,
					Msg:  localityError.Msg,
					Data: nil,
				})
				return
			}

			carriersByLocalitiesDto := mappers.MapperToLocalitiesCarriersDTO(carriersByLocalities)
			response.JSON(w, http.StatusOK, dto.ResponseDTO{
				Code: http.StatusOK,
				Msg:  "success",
				Data: carriersByLocalitiesDto,
			})
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return
		}

		localityCarriers, err := hand.service.GetCarriersByLocality(id)
		if err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToLocalityCarrierDTO(localityCarriers),
		})
	}
}
