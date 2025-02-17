package repository

import (
	"database/sql"
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	"github.com/go-sql-driver/mysql"
)

type PurchaseOrderRepository struct {
	db *sql.DB
}

func NewPurchaseOrderRepository(db *sql.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}

func (b *PurchaseOrderRepository) GetById(id int) (purchaseOrder entities.PurchaseOrderEntity, err error) {
	query, args := purchaseOrder.GetByIdQuery(id)

	err = b.db.QueryRow(query, args...).Scan(
		&purchaseOrder.ID,
		&purchaseOrder.OrderNumber,
		&purchaseOrder.OrderDate,
		&purchaseOrder.TrackingCode,
		&purchaseOrder.BuyerID,
		&purchaseOrder.CarrierID,
		&purchaseOrder.OrderStatusID,
		&purchaseOrder.WarehouseID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repository.ErrPurchaseOrderNotFoundWithID
			return
		}
	}

	return
}

func (b *PurchaseOrderRepository) Create(purchaseOrder entities.PurchaseOrderEntity) (newPurchaseOrder entities.PurchaseOrderEntity, err error) {
	query, args := purchaseOrder.GetCreateQuery()
	result, err := b.db.Exec(query, args...)

	if err != nil {
		var mySqlErr *mysql.MySQLError
		if errors.As(err, &mySqlErr) && mySqlErr.Number == 1452 {
			err = repository.ErrForeignKeysNotValid
			return
		}

		return
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return
	}

	newPurchaseOrder, err = b.GetById(int(lastId))
	return
}
