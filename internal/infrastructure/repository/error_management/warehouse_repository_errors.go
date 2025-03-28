package error_management

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
)

var (
	ErrIdNotFound             = errors.New("id not found")
	ErrIdDuplicate            = errors.New("id duplicate")
	ErrWarehouseCodeDuplicate = errors.New("warehouse code duplicate")
	ErrDataBase               = errors.New("database error")
	ErrLocalityId             = errors.New("locality id not found")
	ErrUpdateBySameData       = errors.New("enter different values to update")
	ErrFKConstraintFail       = errors.New("foreign key constraint fails")
)

func HandleWarehouseRepositoryError(err error) error {
	switch e := err.(type) {
	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			return repository.ErrWarehouseCodeDuplicate
		case 1452:
			return repository.ErrWarehouseLocalityIdNotFound
		case 1451:
			return repository.ErrWarehouseFKConstraintFail
		default:
			return repository.ErrWarehouseDataBase
		}
	default:
		return repository.ErrWarehouseDataBase
	}
}
