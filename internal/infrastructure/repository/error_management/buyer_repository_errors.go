package error_management

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/go-sql-driver/mysql"
)

func HandleBuyerRepositoryError(err error) error {
	var mySqlErr *mysql.MySQLError
	if errors.As(err, &mySqlErr) {
		switch mySqlErr.Number {
		case 1062:
			return repository.ErrBuyerDuplicateCardNumberID
		case
			1451:
			return repository.ErrBuyerCannotDeleteBuyerWithOrders
		default:
			return repository.ErrBuyerUnexpectedError
		}
	}
	return repository.ErrBuyerUnexpectedError
}
