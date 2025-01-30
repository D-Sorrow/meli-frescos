package error_management

import (
	"errors"
)

func InternalServerErr() error {
	return errors.New("internal server error")
}

func ErrIdNotFound() error {
	return errors.New("id not found")
}

func ErrIdDuplicate() error {
	return errors.New("id already exists")
}

func ErrWarehouseCodeDuplicate() error {
	return errors.New("warehouse code already exists")
}
