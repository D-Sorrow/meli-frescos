package error_management

import "errors"

var ErrNotFound = errors.New("not found error")
var ErrAlreadyExists = errors.New("seller already exists")
