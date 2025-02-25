package error_management

import "log"

func HandleRepositoryError(repositoryError error, err error) error {
	log.Printf("%v - error: %v", repositoryError.Error(), err.Error())
	return repositoryError
}
