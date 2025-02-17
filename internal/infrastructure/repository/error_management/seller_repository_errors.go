package error_management

import "errors"

var ErrSellerNotFound = errors.New("not found error")
var ErrSellerAlreadyExists = errors.New("seller already exists")
var ErrSellerCannotBeDeleted = errors.New("seller cant be deleted")
var ErrSellerCannotBeUpdated = errors.New("seller cant be updated")
var ErrSellerCannotBeCreated = errors.New("seller cant be created")
