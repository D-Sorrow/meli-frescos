package handlers

import (
	"context"
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	mapper "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"

	"github.com/bootcamp-go/web/response"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
)

type SectionstHandler struct {
	serv service.SectionsService
}

func NewSectionsHandler(serv service.SectionsService) *SectionstHandler {
	return &SectionstHandler{serv: serv}
}

func (hand *SectionstHandler) GetSections(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		mapSections := hand.serv.GetSections()
		mapSectionsDto := mapper.MapperToSectionsDto(mapSections)
		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Products successfully retrieved",
			Data: mapSectionsDto,
		})

		return nil
	}
}

func (hand *SectionstHandler) GetSectionsById(ctx *context.Context) web.Handler {
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

		sections, err := hand.serv.GetSectionsById(id)

		sectionsDto := mapper.MapperToSectionDto(sections)
		if err != nil {
			response.JSON(writer, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}

		response.JSON(writer, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Msg:  "Product successfully retrieved",
			Data: sectionsDto,
		})

		return nil
	}
}

func (hand *SectionstHandler) SaveSections(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		var sections dto.SectionsDto

		if err := json.NewDecoder(r.Body).Decode(&sections); err != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}

		errValidate := sections.Validate()
		if errValidate != nil {
			response.JSON(writer, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  errValidate.Error(),
				Data: nil,
			})
			return nil
		}

		errSave := hand.serv.SaveSections(mapper.MapperToSectionsModel(sections))
		if errSave != nil {
			response.JSON(writer, http.StatusConflict, dto.ResponseDTO{
				Code: http.StatusConflict,
				Msg:  errSave.Error(),
				Data: nil,
			})
			return nil
		}

		return nil
	}
}

func (hand *SectionstHandler) DeleteSections(ctx *context.Context) web.Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(writer, http.StatusBadRequest, dto.ResponseDTO{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
				Data: nil,
			})
			return nil
		}
		errD := hand.serv.DeleteSections(id)
		if errD != nil {
			response.JSON(writer, http.StatusNotFound, dto.ResponseDTO{
				Code: http.StatusNotFound,
				Msg:  errD.Error(),
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
