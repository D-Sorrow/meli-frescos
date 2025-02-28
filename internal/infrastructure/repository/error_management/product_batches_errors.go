package error_management

import "errors"

var ErrProductBatchesNotFound = errors.New("not found error")
var ErrProductBatchesAlredyExists = errors.New("productBatches alredy exist")
var ErrProductBatchesNotCreated = errors.New("productBatchess cant be created")
