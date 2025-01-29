package repository_errors

import "errors"

var ErrNotFound = errors.New("not found error")
var ErrAlreadyExists = errors.New("seller already exists")
