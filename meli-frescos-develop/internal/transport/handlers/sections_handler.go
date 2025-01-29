package handlers

import (
	"fmt"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	mapper "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"

	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/bootcamp-go/web/response"
)

type SectionstHandler struct {
	serv service.SectionsService
}

func NewSectionsHandler(serv service.SectionsService) *SectionstHandler {
	return &SectionstHandler{serv: serv}
}

func (hand *SectionstHandler) GetSections() http.HandlerFunc {

	fmt.Println("ASdasfasfggg")

	return func(writer http.ResponseWriter, request *http.Request) {
		mapSections := hand.serv.GetSections()
		mapSectionsDto := mapper.MapperToSectionsDto(mapSections)
		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Products successfully retrieved",
			Data: mapSectionsDto,
		})
	}
}
