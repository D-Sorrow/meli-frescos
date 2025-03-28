package service

import (
	"testing"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	repoMock "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/mocks/internal_/infrastructure/repository"
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

	t.Run("case 2: create fails", func(t *testing.T) {
		mockRepo := new(repoMock.MockInboundOrderRepository)

		inboundOrder := &models.InboundOrder{
			OrderDate:      "2023-03",
			OrderNumber:    "ORD12345",
			EmployeeId:     1,
			ProductBatchId: 2,
			WarehouseId:    3,
		}

		mockRepo.On("CreateInboundOrder", inboundOrder).
			Return(repository.ErrInboundOrderDateInvalid)

		serv := NewInboundOrderService(mockRepo)

		err := serv.CreateInboundOrder(inboundOrder)

		require.Error(t, err)
		require.Equal(t, service.ErrInboundOrderDateInvalid, err)
	})

}
