package error_management

const ProductNotFound = "Product not found"
const ProductIsAlreadyExist = "Product is already exist"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}
