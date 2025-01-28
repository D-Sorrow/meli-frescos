package error_management

const IdInvalid = "ID must be positive"
const DescriptionInvalid = "Description must not be empty"
const ExpirationInvalid = "Rates must not be negative"
const FreezingInvalid = "FreezingRate must not be negative"
const DimensionsInvalid = "Invalid dimensions"
const NetWeightInvalid = "Net weight must be positive"
const ProductCodeInvalid = "ProductCode must not be empty"
const ProductTypeInvalid = "ProductTypeId must not be negative"
const SellerIdInvalid = "SellerId must not be negative"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}
