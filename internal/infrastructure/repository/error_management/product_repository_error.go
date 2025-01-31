package error_management

import (
	"errors"
)

const codeProductIsExist = "Code Product Is Exist"
const productNotExist = "Product Not Exist"
const attributeIsNotValid = "Attribute Is Not Valid"

var CodeProductIsExistErr error = errors.New(codeProductIsExist)
var ProductNotExistErr error = errors.New(productNotExist)
var AttributeIsNotValidErr error = errors.New(attributeIsNotValid)
