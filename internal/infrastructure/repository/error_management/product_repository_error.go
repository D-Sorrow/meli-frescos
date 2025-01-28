package error_management

const CodeProductIsExist = "Code Product Is Exist"
const ProductNotExist = "Product Not Exist"
const AttributeIsNotValid = "Attribute Is Not Valid"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}
