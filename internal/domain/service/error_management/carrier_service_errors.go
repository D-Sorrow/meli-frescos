package error_management

import "errors"

func ErrCarrierCidDuplicate() error {
	return errors.New("carrier cid already exists")
}
