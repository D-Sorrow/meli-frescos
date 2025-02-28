package error_management

import (
	"fmt"
	"log"
)

func HandleServiceError(serviceError error, err error) error {
	log.Printf("%v: %v", serviceError.Error(), err.Error())
	return fmt.Errorf("%w: %v", serviceError, err)
}
