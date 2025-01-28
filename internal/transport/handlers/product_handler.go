package handlers

import (
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
