package error_management

import "errors"

var ErrLocalityNotFound = errors.New("not found error")
var ErrLocalityAlreadyExists = errors.New("locality already exists")
var ErrLocalityCannotBeCreated = errors.New("locality cant be created")
var ErrGetAllLocalities = errors.New("could not get all localities")
var ErrProvinceNotFound = errors.New("not found error")
