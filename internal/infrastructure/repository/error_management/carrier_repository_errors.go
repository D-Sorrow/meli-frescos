package error_management

import (
	"errors"
)

var (
	ErrCarrierCidDuplicate = errors.New("carrier cid exists")
)
