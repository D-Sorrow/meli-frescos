package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	serviceErrors "github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type CarrierService struct {
	repository repository.CarrierRepositoryInterface
}

func NewCarryService(repository repository.CarrierRepositoryInterface) *CarrierService {
	return &CarrierService{repository: repository}
}

func (cs *CarrierService) GetAllCarriers() ([]models.Carrier, error) {
	carriers, err := cs.repository.GetAllCarriers()
	if err != nil {
		return nil, serviceErrors.HandleErrorCarrierService(err)
	}

	return carriers, nil
}

func (cs *CarrierService) CreateCarrier(carrier models.Carrier) (models.Carrier, error) {
	newCarrier, err := cs.repository.CreateCarrier(carrier)
	if err != nil {
		return models.Carrier{}, serviceErrors.HandleErrorCarrierService(err)
	}
	return newCarrier, nil
}
