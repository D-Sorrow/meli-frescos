package handlers

import (
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type HandlerSeller struct {
	service service.SellerService
}

func NewHandlerService(service service.SellerService) *HandlerSeller {
	return &HandlerSeller{service: service}
}

func (hand *HandlerSeller) GetSellers() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		mapSeller, _ := hand.service.GetSellers()
		response.JSON(writer, http.StatusOK, mapSeller)
	}
}

func (hand *HandlerSeller) GetSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "id must be a number",
			})
			return
		}
		seller, err := hand.service.GetSellerById(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "user not found",
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    seller,
		})

	}
}
