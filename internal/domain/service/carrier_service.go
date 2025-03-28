package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	serviceErrors "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
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
