package errormanagement

import (
	"errors"
)

var ErrIdNotFound = errors.New("id not found")
var ErrIdDuplicate = errors.New("id duplicate")
var ErrWarehouseCodeDuplicate = errors.New("warehouse code duplicate")
