package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
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
