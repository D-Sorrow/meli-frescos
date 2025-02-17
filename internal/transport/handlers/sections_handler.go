package handlers

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	mapper "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/go-chi/chi/v5"

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
	return func(writer http.ResponseWriter, r *http.Request) {
		mapSections := hand.serv.GetSections()
		mapSectionsDto := mapper.MapperToSectionsDto(mapSections)
		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Products successfully retrieved",
			Data: mapSectionsDto,
		})
	}
}

func (hand *SectionstHandler) GetSectionsById() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		id, errConv := strconv.Atoi(chi.URLParam(r, "id"))
		if errConv != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  errConv.Error(),
				Data: nil,
			})
			return
		}

		sections, err := hand.serv.GetSectionsById(id)

		sectionsDto := mapper.MapperToSectionDto(sections)
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
			Data: sectionsDto,
		})

	}
}

func (hand *SectionstHandler) SaveSections() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		var sections dto.SectionsDto

		if err := json.NewDecoder(r.Body).Decode(&sections); err != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		errValidate := sections.Validate()
		if errValidate != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  errValidate.Error(),
				Data: nil,
			})
			return
		}

		errSave := hand.serv.SaveSections(mapper.MapperToSectionsModel(sections))
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

func (hand *SectionstHandler) DeleteSections() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		errD := hand.serv.DeleteSections(id)
		if errD != nil {
			response.JSON(writer, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  errD.Error(),
				Data: nil,
			})
			return
		}

		response.JSON(writer, http.StatusNoContent, dto.ResponseDTO{
			Code: http.StatusNoContent,
			Msg:  "Product successfully deleted",
			Data: nil,
		})
	}
}

//func (hand *SectionstHandler) UpdateSections() http.HandlerFunc {
/*	return func(writer http.ResponseWriter, r *http.Request) {
	id, errConv := strconv.Atoi(chi.URLParam(r, "id"))
	if errConv != nil {

		response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
			Code: http.StatusBadRequest,
			Msg:  "id must be a number",
			Data: nil,
		})
		return
	}

	var sectionsDto dto.SectionsUpdateDto

	if errConv := json.NewDecoder(r.Body).Decode(&sectionsDto); errConv != nil {
		response.JSON(w)
		return
	}
}*/
//	panic("impl")
//}
