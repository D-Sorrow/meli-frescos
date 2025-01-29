package service

import (
	"fmt"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
)

type WarehouseService struct {
	repository repository.WarehouseRepositoryInterface
}

func NewWarehouseService(repository repository.WarehouseRepositoryInterface) *WarehouseService {
	return &WarehouseService{repository: repository}
}

func (ws *WarehouseService) GetWarehouses() map[int]models.Warehouse {
	return ws.repository.GetWarehouses()
}

func (ws *WarehouseService) GetWarehouseById(id int) (models.Warehouse, error) {
	warehouse, err := ws.repository.GetWarehouseById(id)
	if err != nil {
		return models.Warehouse{}, fmt.Errorf("id %d does not exist", id)
	}

	return warehouse, nil
}
