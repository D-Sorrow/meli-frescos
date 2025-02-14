package error_management

import (
	"fmt"
)

type LocalityErrors struct {
	Msg string
}

func (se *LocalityErrors) Error() string {
	return fmt.Sprintf("%s", se.Msg)
}

var ErrLocalityNotFound *LocalityErrors = &LocalityErrors{
	Msg: "not found error",
}

var ErrLocalityAlreadyExists *LocalityErrors = &LocalityErrors{
	Msg: "locality already exists",
}
