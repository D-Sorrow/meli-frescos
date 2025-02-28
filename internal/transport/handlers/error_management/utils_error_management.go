package error_management

import (
	"log"
)

func HandleHandlerError(handlerError error) error {
	log.Printf("ERR_HLR: %v", handlerError.Error())
	return handlerError
}
