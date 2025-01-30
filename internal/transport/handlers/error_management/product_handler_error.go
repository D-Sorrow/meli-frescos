package error_management

type ProductError struct {
	Message string
	Code    int
}

func (e ProductError) Error() string {
	return e.Message
}
