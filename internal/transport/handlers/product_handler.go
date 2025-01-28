package handlers

import (
	"encoding/json"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	mapper "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	serv service.ProductService
}

func NewProductHandler(serv service.ProductService) *ProductHandler {
	return &ProductHandler{serv: serv}
}

func (hand *ProductHandler) GetProducts() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		mapProduct := hand.serv.GetProducts()
		mapProductDto := mapper.MapperToProductsDto(mapProduct)

		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Products successfully retrieved",
			Data: mapProductDto,
		})
	}
}

func (hand *ProductHandler) GetProductByID() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, errConv := strconv.Atoi(chi.URLParam(request, "id"))

		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
			return
		}

		product, err := hand.serv.GetProductByID(id)

		productDto := mapper.MapperToProductDto(product)
		if err != nil {
			response.JSON(writer, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Product successfully retrieved",
			Data: productDto,
		})

	}
}

func (hand *ProductHandler) SaveProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var product dto.ProductDto

		if err := json.NewDecoder(request.Body).Decode(&product); err != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		errValidate := product.Validate()
		if errValidate != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  errValidate.Error(),
				Data: nil,
			})
			return
		}

		errSave := hand.serv.SaveProduct(mapper.MapperToProductModel(product))
		if errSave != nil {
			response.JSON(writer, http.StatusConflict, dto.ResponseDTO{
				Code: http.StatusConflict,
				Msg:  errSave.Error(),
				Data: nil,
			})
			return
		}
	}
}

func (hand *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, errConv := strconv.Atoi(chi.URLParam(request, "id"))
		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
			return
		}

		attributes := make(map[string]interface{})

		if err := json.NewDecoder(request.Body).Decode(&attributes); err != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		product, err := hand.serv.UpdateProduct(id, attributes)
		productDto := mapper.MapperToProductDto(product)
		if err != nil {
			response.JSON(writer, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Product successfully updated",
			Data: productDto,
		})

	}
}

func (hand *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, errConv := strconv.Atoi(chi.URLParam(request, "id"))
		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
		}
		errDelete := hand.serv.DeleteProduct(id)
		if errDelete != nil {
			response.JSON(writer, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  errDelete.Error(),
				Data: nil,
			})
		}

		response.JSON(writer, http.StatusNoContent, dto.ResponseDTO{
			Code: http.StatusNoContent,
			Msg:  "Product successfully deleted",
			Data: nil,
		})
	}
}
