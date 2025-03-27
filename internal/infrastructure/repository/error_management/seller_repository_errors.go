package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/go-sql-driver/mysql"
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
