package error_management

import (
	"fmt"
	"log"
)

func HandleRepositoryError(repositoryError error, err error) error {
	log.Printf("%v: %v", repositoryError.Error(), err.Error())
	return fmt.Errorf("%w: %v", repositoryError, err)
}
