package handlers

import (
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	mapper "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
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

		fmt.Println("holaa")
		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Products successfully retrieved",
			Data: mapProductDto,
		})
	}
}
