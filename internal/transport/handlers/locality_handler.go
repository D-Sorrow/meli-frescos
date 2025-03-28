package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type LocalityHandler struct {
	service  service.LocalityService
	validate *validator.Validate
}

func NewLocalityHandler(service service.LocalityService) *LocalityHandler {
	return &LocalityHandler{service: service, validate: validator.New()}
}

func (hand LocalityHandler) CreateLocality(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var localityDto dto.LocalityDto

		if err := json.NewDecoder(r.Body).Decode(&localityDto); err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return nil
		}

		if err := hand.validate.Struct(localityDto); err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return nil
		}

		locality, err := hand.service.CreateLocality(mappers.MapperToLocality(localityDto))
		if err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Locality Created",
			Data: mappers.MapperToLocalityDTO(locality),
		})

		return nil
	}
}

func (hand LocalityHandler) GetSellersByLocality(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return nil
		}
		localitySellers, err := hand.service.GetSellersByLocality(id)
		if err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return nil
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToLocalitySellersDTO(localitySellers),
		})

		return nil
	}
}

func (hand LocalityHandler) GetCarriersByLocality(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
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
				return nil
			}

			carriersByLocalitiesDto := mappers.MapperToLocalitiesCarriersDTO(carriersByLocalities)
			response.JSON(w, http.StatusOK, dto.ResponseDTO{
				Code: http.StatusOK,
				Msg:  "success",
				Data: carriersByLocalitiesDto,
			})
			return nil
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return nil
		}

		localityCarriers, err := hand.service.GetCarriersByLocality(id)
		if err != nil {
			localityError := error_management.HandleErrorLocality(err, ctx)
			response.JSON(w, localityError.Code, dto.ResponseDTO{
				Code: localityError.Code,
				Msg:  localityError.Msg,
				Data: nil,
			})
			return nil
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToLocalityCarrierDTO(localityCarriers),
		})

		return nil
	}
}
