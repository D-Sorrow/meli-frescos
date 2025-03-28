package error_management

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
)

func HandleInboundOrderRepositoryError(err error) error {
	switch e := err.(type) {
	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			return repository.ErrInboundOrderNumberAlreadyExists
		case 1452:
			if strings.Contains(err.Error(), "FOREIGN KEY (`employe_id`)") {
				return repository.ErrInboundOrderEmployeeIdNotFound
			}
			if strings.Contains(err.Error(), "FOREIGN KEY (`product_batch_id`)") {
				return repository.ErrInboundOrderProductBatchIdNotFound
			}
			if strings.Contains(err.Error(), "FOREIGN KEY (`wareHouse_id`)") {
				return repository.ErrInboundOrderWareHouseIdNotFound
			}
			return repository.ErrInboundOrderRepositoryGeneric
		case 1292:
			return repository.ErrInboundOrderDateInvalid
		default:
			return repository.ErrInboundOrderRepositoryGeneric
		}
	default:
		return repository.ErrInboundOrderRepositoryGeneric
	}
}
