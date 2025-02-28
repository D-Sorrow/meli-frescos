package error_management

import (
	"errors"
	"strings"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/go-sql-driver/mysql"
)

func HandlePurchaseOrderRepositoryError(err error) error {
	var mySqlErr *mysql.MySQLError
	if errors.As(err, &mySqlErr) {
		switch mySqlErr.Number {
		case 1452:
			foreignKeyErrors := map[string]func() error{
				"FOREIGN KEY (`wareHouse_id`)":    func() error { return repository.ErrPurchaseOrderFKWareHouseIdNotValid },
				"FOREIGN KEY (`buyer_id`)":        func() error { return repository.ErrPurchaseOrderFKBuyerIdNotValid },
				"FOREIGN KEY (`order_status_id`)": func() error { return repository.ErrPurchaseOrderFKOrderStatusIdNotValid },
				"FOREIGN KEY (`carrier_id`)":      func() error { return repository.ErrPurchaseOrderFKCarrierIdNotValid },
			}

			for foreignKey, errorFunc := range foreignKeyErrors {
				if strings.Contains(err.Error(), foreignKey) {
					return errorFunc()
				}
			}

			return repository.ErrPurchaseOrderUnexpectedError
		default:
			return repository.ErrPurchaseOrderUnexpectedError
		}
	}
	return repository.ErrPurchaseOrderUnexpectedError
}
