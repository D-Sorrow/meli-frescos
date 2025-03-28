package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type HandlerSeller struct {
	service  service.SellerService
	validate *validator.Validate
}

func NewHandlerService(service service.SellerService) *HandlerSeller {
	return &HandlerSeller{service: service, validate: validator.New()}
}

func (hand *HandlerSeller) GetSellers(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		mapSeller, err := hand.service.GetSellers()
		if err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
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

		return nil
	}
}

func (hand *HandlerSeller) GetSeller(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return nil
		}
		seller, err := hand.service.GetSellerById(id)
		if err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}
		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "success",
			Data: mappers.MapperToSellerDTO(seller),
		})

		return nil
	}
}

func (hand *HandlerSeller) CreateSeller(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var sellerDto dto.SellerDto

		if err := json.NewDecoder(r.Body).Decode(&sellerDto); err != nil {
			if errors.Is(err, io.EOF) {
				response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
					Code: http.StatusBadRequest,
					Msg:  "Bad Request - body can not be null",
					Data: nil,
				})
			} else {
				sellerError := error_management.HandleErrorSeller(err, ctx)
				response.JSON(w, sellerError.Code, dto.ResponseDTO{
					Code: sellerError.Code,
					Msg:  sellerError.Msg,
					Data: nil,
				})
			}
			return nil
		}

		if err := hand.validate.Struct(sellerDto); err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}

		seller, err := hand.service.CreateSeller(mappers.MapperToSeller(sellerDto))
		if err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "User Created",
			Data: mappers.MapperToSellerDTO(seller),
		})

		return nil
	}
}

func (hand *HandlerSeller) UpdateSeller(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return nil
		}

		var sellerDto dto.SellerUpdateDto

		if err := json.NewDecoder(r.Body).Decode(&sellerDto); err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}

		if err := hand.validate.Struct(sellerDto); err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}

		seller, err := hand.service.UpdateSeller(id, mappers.MapperToSellerPatch(sellerDto))
		if err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "User Updated",
			Data: mappers.MapperToSellerDTO(seller),
		})

		return nil
	}
}

func (hand *HandlerSeller) DeleteSeller(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  "id must be a number",
				Data: nil,
			})
			return nil
		}

		err = hand.service.DeleteSeller(id)
		if err != nil {
			sellerError := error_management.HandleErrorSeller(err, ctx)
			response.JSON(w, sellerError.Code, dto.ResponseDTO{
				Code: sellerError.Code,
				Msg:  sellerError.Msg,
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusNoContent, dto.ResponseDTO{})

		return nil
	}
}
