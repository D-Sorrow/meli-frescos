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
			errNotFound := serviceErrors.ErrIdNotFound()
			return models.Warehouse{}, errNotFound
		default:
			return models.Warehouse{}, serviceErrors.InternalServerErr()
		}
	}

	return warehouse, nil
}

func (ws *WarehouseService) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	newWarehouse, err := ws.repository.CreateWarehouse(warehouse)
	if err != nil {
		switch {
		case errors.Is(err, repoErros.ErrIdDuplicate):
			err = serviceErrors.ErrIdDuplicate()
			return models.Warehouse{}, err
		case errors.Is(err, repoErros.ErrWarehouseCodeDuplicate):
			err = serviceErrors.ErrWarehouseCodeDuplicate()
			return models.Warehouse{}, err
		default:
			err = serviceErrors.InternalServerErr()
			return models.Warehouse{}, err
		}
	}

	return newWarehouse, nil
}

func (ws *WarehouseService) PatchWarehouse(id int, data map[string]interface{}) (models.Warehouse, error) {
	warehouse, err := ws.repository.PatchWarehouse(id, data)
	if err != nil {
		switch {
		case errors.Is(err, repoErros.ErrIdNotFound):
			errNotFound := serviceErrors.ErrIdNotFound()
			return models.Warehouse{}, errNotFound
		case errors.Is(err, repoErros.ErrWarehouseCodeDuplicate):
			return models.Warehouse{}, serviceErrors.ErrWarehouseCodeDuplicate()
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
			errNotFound := serviceErrors.ErrIdNotFound()
			return errNotFound
		default:
			return serviceErrors.InternalServerErr()
		}
	}

	return nil
}
