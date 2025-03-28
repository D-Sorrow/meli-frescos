package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	serviceErrors "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
)

type WarehouseService struct {
	repository repository.WarehouseRepositoryInterface
}

func NewWarehouseService(repository repository.WarehouseRepositoryInterface) *WarehouseService {
	return &WarehouseService{repository: repository}
}

func (ws *WarehouseService) GetWarehouses() (map[int]models.Warehouse, error) {
	warehouses, err := ws.repository.GetWarehouses()
	if err != nil {
		return nil, serviceErrors.HandleErrorWarehouseService(err)
	}

	return warehouses, nil
}

func (ws *WarehouseService) GetWarehouseById(id int) (warehouse models.Warehouse, err error) {
	warehouse, err = ws.repository.GetWarehouseById(id)
	if err != nil {
		err = serviceErrors.HandleErrorWarehouseService(err)
	}

	return
}

func (ws *WarehouseService) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	newWarehouse, err := ws.repository.CreateWarehouse(warehouse)
	if err != nil {
		return models.Warehouse{}, serviceErrors.HandleErrorWarehouseService(err)
	}

	return newWarehouse, nil
}

func (ws *WarehouseService) PatchWarehouse(
	id int,
	data map[string]interface{},
) (models.Warehouse, error) {
	warehouse, err := ws.repository.PatchWarehouse(id, data)
	if err != nil {
		return models.Warehouse{}, serviceErrors.HandleErrorWarehouseService(err)
	}

	return warehouse, nil
}

func (ws *WarehouseService) DeleteWarehouse(id int) error {
	err := ws.repository.DeleteWarehouse(id)
	if err != nil {
		return serviceErrors.HandleErrorWarehouseService(err)
	}

	return nil
}
