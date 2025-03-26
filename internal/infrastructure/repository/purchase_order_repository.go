package repository

import (
	"database/sql"
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type PurchaseOrderRepository struct {
	db *sql.DB
}

func NewPurchaseOrderRepository(db *sql.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}

func (b *PurchaseOrderRepository) GetById(
	id int,
) (purchaseOrder entities.PurchaseOrderEntity, err error) {
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
			err = error_management.HandleRepositoryError(
				repository.ErrPurchaseOrderNotFoundWithID,
				err,
			)
			return
		}

		err = error_management.HandleRepositoryError(
			repository.ErrPurchaseOrderUnexpectedError,
			err,
		)
	}

	return
}

func (b *PurchaseOrderRepository) Create(
	purchaseOrder entities.PurchaseOrderEntity,
) (newPurchaseOrder entities.PurchaseOrderEntity, err error) {
	query, args := purchaseOrder.GetCreateQuery()
	result, err := b.db.Exec(query, args...)

	if err != nil {
		err = error_management.HandleRepositoryError(
			error_management.HandlePurchaseOrderRepositoryError(err),
			err,
		)
		return
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		err = error_management.HandleRepositoryError(
			repository.ErrPurchaseOrderUnexpectedError,
			err,
		)
		return
	}

	purchaseOrder.ID = lastId
	newPurchaseOrder = purchaseOrder
	return
}
