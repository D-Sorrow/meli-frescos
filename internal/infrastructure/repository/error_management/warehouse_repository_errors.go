package errormanagement

import (
	"errors"
)

var ErrIdNotFound = errors.New("id not found")
var ErrIdDuplicate = errors.New("id duplicate")
var ErrWarehouseCode = errors.New("warehouse code duplicate")
