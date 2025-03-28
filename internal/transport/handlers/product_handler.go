package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	dto "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	mapper "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type ProductHandler struct {
	serv service.ProductService
}

func NewProductHandler(serv service.ProductService) *ProductHandler {
	return &ProductHandler{serv: serv}
}

func (hand *ProductHandler) GetProducts(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		mapProduct, errGet := hand.serv.GetProducts()
		if errGet != nil {
			errSpe := error_management.HandlerErrorProduct(errGet, ctx)
			response.JSON(writer, errSpe.GetCode(), errSpe)
			return nil
		}
		mapProductDto := mapper.MapperToProductsDto(mapProduct)

		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Products successfully retrieved",
			Data: mapProductDto,
		})

		return nil
	}
}

func (hand *ProductHandler) GetProductByID(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		id, errConv := strconv.Atoi(chi.URLParam(r, "id"))

		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
			return nil
		}

		product, err := hand.serv.GetProductByID(id)

		productDto := mapper.MapperToProductDto(product)
		if err != nil {
			errSpe := error_management.HandlerErrorProduct(err, ctx)
			response.JSON(writer, errSpe.GetCode(), dto.ResponseDTO{
				Code: errSpe.GetCode(),
				Msg:  errSpe.Message,
				Data: nil,
			})
			return nil
		}

		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Product successfully retrieved",
			Data: productDto,
		})

		return nil
	}
}

func (hand *ProductHandler) SaveProduct(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		var product dto.ProductDto

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}

		errValidate := product.Validate()
		if errValidate != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  errValidate.Error(),
				Data: nil,
			})
			return nil
		}

		errSave := hand.serv.SaveProduct(mapper.MapperToProductModel(&product))
		if errSave != nil {
			errSpe := error_management.HandlerErrorProduct(errSave, ctx)
			response.JSON(writer, errSpe.GetCode(), dto.ResponseDTO{
				Code: errSpe.Code,
				Msg:  errSpe.Message,
				Data: nil,
			})
			return nil
		}
		response.JSON(writer, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Product successfully saved",
			Data: nil,
		})

		return nil
	}
}

func (hand *ProductHandler) UpdateProduct(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		id, errConv := strconv.Atoi(chi.URLParam(r, "id"))
		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
			return nil
		}

		var att dto.AttributeDto

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&att); err != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}

		if errValidation := att.Validation(); errValidation != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errValidation.Error(),
				Data: nil,
			})
			return nil
		}

		product, errUpdate := hand.serv.UpdateProduct(id, mapper.ModelToMap(att))
		productDto := mapper.MapperToProductDto(product)

		if errUpdate != nil {
			errSpe := error_management.HandlerErrorProduct(errUpdate, ctx)
			response.JSON(writer, errSpe.Code, dto.ResponseDTO{
				Code: errSpe.Code,
				Msg:  errSpe.Error(),
				Data: nil,
			})
			return nil
		}
		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Product successfully updated",
			Data: productDto,
		})

		return nil
	}
}

func (hand *ProductHandler) DeleteProduct(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		id, errConv := strconv.Atoi(chi.URLParam(r, "id"))
		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
			return nil
		}
		errDelete := hand.serv.DeleteProduct(id)
		if errDelete != nil {
			errSp := error_management.HandlerErrorProduct(errDelete, ctx)
			response.JSON(writer, errSp.Code, dto.ResponseDTO{
				Code: errSp.Code,
				Msg:  errSp.Error(),
				Data: nil,
			})
			return nil
		}

		response.JSON(writer, http.StatusNoContent, dto.ResponseDTO{
			Code: http.StatusNoContent,
			Msg:  "Product successfully deleted",
			Data: nil,
		})

		return nil
	}
}
