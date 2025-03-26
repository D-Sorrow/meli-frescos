package error_management

import (
	"errors"
)

const codeSectionsIsExist = "code Product Is Exist"

var ErrCodeSectionsNotUpdate = errors.New("sections not update")

var ErrCodeSectionsIsExistErr error = errors.New(codeSectionsIsExist)
