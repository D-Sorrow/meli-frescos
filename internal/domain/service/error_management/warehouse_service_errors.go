package errormanagement

import "fmt"

type ErrIdNotFound struct {
	Id int
}

func (e *ErrIdNotFound) Error() error {
	return fmt.Errorf("id %d not found", e.Id)
}

type ErrIdDuplicate struct {
	Id int
}

func (e *ErrIdDuplicate) Error() error {
	return fmt.Errorf("id %d already exists", e.Id)
}

type ErrWarehouseCodeDuplicate struct {
	WarehouseCode string
}

func (e *ErrWarehouseCodeDuplicate) Error() error {
	return fmt.Errorf("warehouse code %s already exists", e.WarehouseCode)
}
