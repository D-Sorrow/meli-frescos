package error_management

import (
	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
)

func HandleSellerRepositoryError(err error) error {
	switch e := err.(type) {
	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			return repository.ErrSellerAlreadyExists
		default:
			return repository.ErrSellerRepositoryGeneric
		}
	default:
		return repository.ErrSellerRepositoryGeneric
	}
}
