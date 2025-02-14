package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	serviceErrors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	repoErros "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type CarrierService struct {
	repository repository.CarrierRepositoryInterface
}

func NewCarryService(repository repository.CarrierRepositoryInterface) *CarrierService {
	return &CarrierService{repository: repository}
}

func (cs *CarrierService) CreateCarrier(carrier models.Carrier) (models.Carrier, error) {
	newCarrier, err := cs.repository.CreateCarrier(carrier)
	if err != nil {
		switch {
		case errors.Is(err, repoErros.ErrIdDuplicate):
			err = serviceErrors.ErrIdDuplicate()
			return models.Carrier{}, err
		case errors.Is(err, repoErros.ErrCarrierCidDuplicate):
			err = serviceErrors.ErrCarrierCidDuplicate()
			return models.Carrier{}, err
		case errors.Is(err, repoErros.ErrLocalityId):
			err = serviceErrors.ErrEntityId()
			return models.Carrier{}, err
		default:
			err = serviceErrors.InternalServerErr()
			return models.Carrier{}, err
		}
	}
	return newCarrier, nil
}
