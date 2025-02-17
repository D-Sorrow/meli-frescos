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

func ErrEntityId() error {
	return errors.New("entity id faild")
}

func ErrUpdateBySameData() error {
	return errors.New("enter different data to update")
}

func ErrFKConstraintFail() error {
	return errors.New("foreign key constraint fails")
}
