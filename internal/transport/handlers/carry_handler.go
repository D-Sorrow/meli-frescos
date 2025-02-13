package handlers

import (
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/bootcamp-go/web/response"
)

type CarryHandler struct {
	service service.CarryServiceInterface
}

func NewCarryHandler(service service.CarryServiceInterface) *CarryHandler {
	return &CarryHandler{service: service}
}

func (ch *CarryHandler) GetTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, "Hellow Carries")
	}
}
