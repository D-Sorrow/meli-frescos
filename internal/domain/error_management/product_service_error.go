package error_management

const ProductNotFound = "Product not found"
const ProductIsAlreadyExist = "Product is already exist"
const CodeNotFound = 404

type ProductError struct {
	Message string
	Code    int
}

func (e ProductError) Error() string {
	return e.Message
}
