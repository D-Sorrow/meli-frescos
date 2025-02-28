package repository

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type InboundOrderRepository struct {
	db *sql.DB
}

func NewInboundOrderRepository(db *sql.DB) *InboundOrderRepository {
	return &InboundOrderRepository{db: db}
}

func (_repository *InboundOrderRepository) CreateInboundOrder(inboundOrder *models.InboundOrder) error {

	result, err := _repository.db.Exec(
		"INSERT INTO inbound_orders (`order_date`, `order_number`, `employe_id`, `product_batch_id`, `wareHouse_id`) VALUES (?, ?, ?, ?, ?)",
		(*inboundOrder).OrderDate, (*inboundOrder).OrderNumber, (*inboundOrder).EmployeeId, (*inboundOrder).ProductBatchId, (*inboundOrder).WarehouseId,
	)
	if err != nil {
		return error_management.HandleRepositoryError(error_management.HandleInboundOrderRepositoryError(err), err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return error_management.HandleRepositoryError(repository.ErrInboundOrderLastInsertId, err)
	}

	(*inboundOrder).Id = int(lastInsertId)
	return nil
}
