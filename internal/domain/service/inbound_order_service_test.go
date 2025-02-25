package service

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	repoMock "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func TestCreateInboundOrder(t *testing.T) {
	t.Run("case 1: create success", func(t *testing.T) {
		mockRepo := new(repoMock.MockInboundOrderRepository)

		inboundOrder := &models.InboundOrder{
			OrderDate:      "2023-03-18",
			OrderNumber:    "ORD12345",
			EmployeeId:     1,
			ProductBatchId: 2,
			WarehouseId:    3,
		}

		mockRepo.On("CreateInboundOrder", inboundOrder).Return(nil)

		serv := NewInboundOrderService(mockRepo)

		err := serv.CreateInboundOrder(inboundOrder)

		require.NoError(t, err)
	})

	t.Run("case 1: create success", func(t *testing.T) {
		mockRepo := new(repoMock.MockInboundOrderRepository)

		inboundOrder := &models.InboundOrder{
			OrderDate:      "2023-03",
			OrderNumber:    "ORD12345",
			EmployeeId:     1,
			ProductBatchId: 2,
			WarehouseId:    3,
		}

		mockRepo.On("CreateInboundOrder", inboundOrder).Return(repository.ErrInboundOrderDateInvalid)

		serv := NewInboundOrderService(mockRepo)

		err := serv.CreateInboundOrder(inboundOrder)

		require.Error(t, err)
		require.Equal(t, service.ErrInboundOrderDateInvalid, err)
	})

}
