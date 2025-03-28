package error_management

import (
	"errors"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
)

func HandleErrorWarehouseService(err error) error {
	switch {
	case errors.Is(err, repository.ErrWarehouseNotFound):
		return service.ErrWarehouseNotFound
	case errors.Is(err, repository.ErrWarehouseIdDuplicate):
		return service.ErrWarehouseIdDuplicate
	case errors.Is(err, repository.ErrWarehouseCodeDuplicate):
		return service.ErrWarehouseCodeDuplicate
	case errors.Is(err, repository.ErrWarehouseLocalityIdNotFound):
		return service.ErrWarehouseLocalityIdNotFound
	case errors.Is(err, repository.ErrWarehouseUpdateBySameData):
		return service.ErrWarehouseUpdateBySameData
	case errors.Is(err, repository.ErrWarehouseGetUpdatedOrCreatedItem):
		return service.ErrWarehouseGetUpdatedOrCreatedItem
	case errors.Is(err, repository.ErrWarehouseFKConstraintFail):
		return service.ErrWarehouseFKConstraintFail
	default:
		return service.ErrWarehouseServiceDefault
	}
}

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
