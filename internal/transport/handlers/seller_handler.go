package handlers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/bootcamp-go/web/response"
	"net/http"
)

type HandlerSeller struct {
	service service.SellerService
}

func NewHandlerService(service service.SellerService) *HandlerSeller {
	return &HandlerSeller{service: service}
}

func (hand *HandlerSeller) GetSeller() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		mapSeller, _ := hand.service.GetSellers()
		response.JSON(writer, http.StatusOK, mapSeller)
	}
}
