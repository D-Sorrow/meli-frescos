package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"

type CarryService struct {
	repository repository.CarryRepositoryInterface
}

func NewCarryService(repository repository.CarryRepositoryInterface) *CarryService {
	return &CarryService{repository: repository}
}
