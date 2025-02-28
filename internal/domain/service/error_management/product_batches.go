package error_management

import "fmt"

type ProductBatchesError struct {
	Msg string
}

func (send *ProductBatchesError) Error() string {
	return fmt.Sprintf("%s", send.Msg)
}

var ErrProductBatchesNotFound *ProductBatchesError = &ProductBatchesError{
	Msg: "Not Fund error ",
}

var ErrProductBatchesAlredyExists *ProductBatchesError = &ProductBatchesError{
	Msg: "ProductBatches alredy exist",
}

var ErrProductBatchesAllSections *ProductBatchesError = &ProductBatchesError{
	Msg: "no hay sections",
}
