package error_management

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/go-sql-driver/mysql"
)

func HandleLocalityRepositoryError(err error) error {
	switch e := err.(type) {
	case *mysql.MySQLError:
		switch e.Number {
		case 1062:
			return repository.ErrLocalityAlreadyExists
		default:
			return repository.ErrLocalityRepositoryGeneric
		}
	default:
		return repository.ErrLocalityRepositoryGeneric
	}
}
