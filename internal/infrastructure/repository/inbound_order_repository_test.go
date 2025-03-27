package repository

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCreateInboundOrder_Success(t *testing.T) {
	// Prepare the test database and mock
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewInboundOrderRepository(db)

	inboundOrder := &models.InboundOrder{
		OrderDate:      "2023-03-18",
		OrderNumber:    "ORD12345",
		EmployeeId:     1,
		ProductBatchId: 2,
		WarehouseId:    3,
	}

	// Mock the expected behavior
	mock.ExpectExec("INSERT INTO inbound_orders").
		WithArgs(inboundOrder.OrderDate, inboundOrder.OrderNumber, inboundOrder.EmployeeId, inboundOrder.ProductBatchId, inboundOrder.WarehouseId).
		WillReturnResult(sqlmock.NewResult(1, 1)) // 1 is the last insert ID

	// Call the actual method
	err = repo.CreateInboundOrder(inboundOrder)

	require.NoError(t, err)
	require.Equal(t, 1, inboundOrder.Id) // Check if the inbound order ID is set correctly

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
