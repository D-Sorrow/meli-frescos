package handlers

import (
	"encoding/json"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"net/http"
	"strconv"
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

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		var productRecord = dto.ProductRecordDto{
			ProductId:      data["data"].ProductId,
			SalePrice:      data["data"].SalePrice,
			PurchasePrice:  data["data"].PurchasePrice,
			LastUpdateTime: data["data"].LastUpdateTime,
		}
		errValidate := productRecord.Validation()
		if errValidate != nil {
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  errValidate.Error(),
				Data: nil,
			})
			return
		}

		errSave := hand.service.SaveProductRecord(mappers.ToProductRecordModel(&productRecord))
		if errSave != nil {
			errSpe := error_management.HandlerErrProductRecord(errSave)
			response.JSON(w, errSpe.Code, dto.ResponseDTO{
				Code: errSpe.Code,
				Msg:  errSpe.Error(),
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
func (hand *ProductRecordHandler) GetProductRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var productId int
		productId, _ = strconv.Atoi(id)
		records, errGet := hand.service.GetProductRecord(productId)
		if errGet != nil {
			errSpe := error_management.HandlerErrProductRecord(errGet)
			response.JSON(w, errSpe.Code, dto.ResponseDTO{
				Code: errSpe.Code,
				Msg:  errSpe.Error(),
				Data: nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Data: records,
		})
		return
	}
}
