package error_management

import "fmt"

const SectionsNotFound = "sections not found "
const SectionsIsAlreadyExist = "sections is already exist ASDASFSAF"
const CodeNotFound = 404

type Error struct {
	Message string
	Code    int
}

type ErrorSections struct {
	Msg string
}

func (sec *ErrorSections) Error() string {
	return fmt.Sprintf("%s", sec.Msg)
}

func (e Error) Error() string {
	return e.Message
}

var ErrSectionsNotFound *ErrorSections = &ErrorSections{
	Msg: "not found error",
}
