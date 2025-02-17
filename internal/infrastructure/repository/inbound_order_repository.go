package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type InboundOrderRepository struct {
	db *sql.DB
}

func NewInboundOrderRepository(db *sql.DB) *InboundOrderRepository {
	return &InboundOrderRepository{db: db}
}

func (repository *InboundOrderRepository) CreateInboundOrder(inboundOrder *models.InboundOrder) error {

	result, err := repository.db.Exec(
		"INSERT INTO inbound_orders (`order_date`, `order_number`, `employe_id`, `product_batch_id`, `wareHouse_id`) VALUES (?, ?, ?, ?, ?)",
		(*inboundOrder).OrderDate, (*inboundOrder).OrderNumber, (*inboundOrder).EmployeeId, (*inboundOrder).ProductBatchId, (*inboundOrder).WarehouseId,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "for key 'inbound_orders.order_number'") {
			return errors.New("ONAE_DB")
		}
		if strings.Contains(err.Error(), "foreign key constraint fails") {
			if strings.Contains(err.Error(), "FOREIGN KEY (`employe_id`)") {
				return errors.New("EIDNF_DB")
			}
			if strings.Contains(err.Error(), "FOREIGN KEY (`product_batch_id`)") {
				return errors.New("PBIDNF_DB")
			}
			if strings.Contains(err.Error(), "FOREIGN KEY (`wareHouse_id`)") {
				return errors.New("WHIDNF_DB")
			}
		}
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	(*inboundOrder).Id = int(lastInsertId)
	return nil
}
