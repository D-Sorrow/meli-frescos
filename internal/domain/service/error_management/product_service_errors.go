package error_management

const CodeUseCaseError = "004"

type ErrService struct {
	Code string
}

func (e *ErrService) Error() string {
	return e.Code
}
