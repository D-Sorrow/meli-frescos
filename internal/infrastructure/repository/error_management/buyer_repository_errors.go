package error_management

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
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
