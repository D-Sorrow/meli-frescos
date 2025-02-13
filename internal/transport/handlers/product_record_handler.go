package handlers

import (
	"encoding/json"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"net/http"
)

type ProductRecordHandler struct {
	service service.ProductRecordService
}

func NewProductRecordHandler(service service.ProductRecordService) *ProductRecordHandler {
	return &ProductRecordHandler{
		service: service,
	}
}

func (hand *ProductRecordHandler) SaveProductRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data map[string]dto.ProductRecordDto
		var productRecord dto.ProductRecordDto

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {

		}
		errValidate := productRecord.Validation()
		if errValidate != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		errSave := hand.service.SaveProductRecord(mappers.ToProductRecordModel(&productRecord))
		if errSave != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  errSave.Error(),
				Data: nil,
			})
			return
		}
		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Product record saved",
			Data: nil,
		})
		return
	}
}
