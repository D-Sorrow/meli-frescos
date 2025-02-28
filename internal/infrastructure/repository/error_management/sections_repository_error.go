package error_management

import (
	"errors"
)

const codeSectionsIsExist = "Code Product Is Exist"

var CodeSectionsNotUpdate = errors.New("sections not update")

var CodeSectionsIsExistErr error = errors.New(codeSectionsIsExist)
