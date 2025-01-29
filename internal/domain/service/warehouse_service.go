package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	errormanagement "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
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

func (ws *WarehouseService) GetWarehouseById(id int) (warehouse models.Warehouse, err error) {
	warehouse, err = ws.repository.GetWarehouseById(id)
	if err != nil {
		notFound := errormanagement.ErrIdNotFound{Id: id}
		return models.Warehouse{}, notFound.Error()
	}

	return warehouse, nil
}
