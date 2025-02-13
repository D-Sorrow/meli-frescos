package error_management

import (
	"errors"

	"github.com/go-sql-driver/mysql"
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

func MySqlErrors(err mysql.MySQLError) error {
	switch err.Number {
	case 1062:
		return ErrWarehouseCodeDuplicate
	case 1452:
		return ErrLocalityId
	case 1451:
		return ErrFKConstraintFail
	}

	return ErrDataBase
}
