package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	serviceErrors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	repoErros "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
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
		switch {
		case errors.Is(err, repoErros.ErrIdNotFound):
			errNotFound := serviceErrors.ErrIdNotFound{Id: id}
			return models.Warehouse{}, errNotFound.Error()
		default:
			return models.Warehouse{}, serviceErrors.InternalServerErr()
		}
	}

	return warehouse, nil
}

func (ws *WarehouseService) DeleteWarehouse(id int) error {
	err := ws.repository.DeleteWarehouse(id)
	if err != nil {
		switch {
		case errors.Is(err, repoErros.ErrIdNotFound):
			errNotFound := serviceErrors.ErrIdNotFound{Id: id}
			return errNotFound.Error()
		default:
			return serviceErrors.InternalServerErr()
		}
	}

	return nil
}
