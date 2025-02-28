package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/go-sql-driver/mysql"
)

func HandleCarrierRepositoryError(err error) error {
	switch e := err.(type) {
	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			return repository.ErrCarrierCidDuplicate
		case 1452:
			return repository.ErrCarrierLocalityId
		case 1451:
			return repository.ErrCarrierFKConstraintFail
		default:
			return repository.ErrCarrierDataBase
		}
	default:
		return repository.ErrCarrierDataBase
	}
}
