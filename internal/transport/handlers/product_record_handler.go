package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
)

type ProductRecordHandler struct {
	service service.ProductRecordService
}

func NewProductRecordHandler(service service.ProductRecordService) *ProductRecordHandler {
	return &ProductRecordHandler{
		service: service,
	}
}

func (hand *ProductRecordHandler) SaveProductRecord(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var data map[string]dto.ProductRecordDto

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			errSpe := error_management.HandlerErrProductRecord(err, ctx)
			response.JSON(w, http.StatusInternalServerError, dto.ResponseDTO{
				Code: http.StatusInternalServerError,
				Msg:  errSpe.Error(),
				Data: nil,
			})
			return nil
		}
		var productRecord = dto.ProductRecordDto{
			ProductId:      data["data"].ProductId,
			SalePrice:      data["data"].SalePrice,
			PurchasePrice:  data["data"].PurchasePrice,
			LastUpdateTime: data["data"].LastUpdateTime,
		}
		errValidate := productRecord.Validation()
		if errValidate != nil {
			response.JSON(w, http.StatusUnprocessableEntity, dto.ResponseDTO{
				Code: http.StatusUnprocessableEntity,
				Msg:  errValidate.Error(),
				Data: nil,
			})
			return nil
		}

		record, errSave := hand.service.SaveProductRecord(
			mappers.ToProductRecordModel(&productRecord),
		)
		if errSave != nil {
			errSpe := error_management.HandlerErrProductRecord(errSave, ctx)
			response.JSON(w, errSpe.Code, dto.ResponseDTO{
				Code: errSpe.Code,
				Msg:  errSpe.Error(),
				Data: nil,
			})
			return nil
		}
		response.JSON(w, http.StatusCreated, dto.ResponseDTO{
			Code: http.StatusCreated,
			Msg:  "Product record saved",
			Data: record,
		})

		return nil
	}
}
func (hand *ProductRecordHandler) GetProductRecord(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.URL.Query().Get("id")
		var productId int
		productId, _ = strconv.Atoi(id)
		records, errGet := hand.service.GetProductRecord(productId)
		if errGet != nil {
			errSpe := error_management.HandlerErrProductRecord(errGet, ctx)
			response.JSON(w, errSpe.Code, dto.ResponseDTO{
				Code: errSpe.Code,
				Msg:  errSpe.Error(),
				Data: nil,
			})
			return nil
		}

		response.JSON(w, http.StatusOK, dto.ResponseDTO{
			Code: http.StatusOK,
			Data: records,
		})

		return nil
	}
}
