package error_management

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

var inboundOrderServiceErrors = map[error]error{
	repository.ErrInboundOrderNumberAlreadyExists:    service.ErrInboundOrderNumberAlreadyExists,
	repository.ErrInboundOrderEmployeeIdNotFound:     service.ErrInboundOrderEmployeeIdNotFound,
	repository.ErrInboundOrderProductBatchIdNotFound: service.ErrInboundOrderProductBatchIdNotFound,
	repository.ErrInboundOrderWareHouseIdNotFound:    service.ErrInboundOrderWareHouseIdNotFound,
	repository.ErrInboundOrderLastInsertId:           service.ErrInboundOrderLastInsertId,
	repository.ErrInboundOrderRepositoryGeneric:      service.ErrInboundOrderServiceGeneric,
	repository.ErrInboundOrderDateInvalid:            service.ErrInboundOrderDateInvalid,
}

func HandleInboundOrderServiceError(err error) error {
	if e, exists := inboundOrderServiceErrors[err]; exists {
		return e
	}
	return inboundOrderServiceErrors[repository.ErrInboundOrderRepositoryGeneric]
}
