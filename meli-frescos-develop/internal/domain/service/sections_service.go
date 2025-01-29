package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
)

type SectionsService struct {
	repo repository.SectionsRepository
}

func NewSectionsService(repo repository.SectionsRepository) *SectionsService {
	return &SectionsService{repo: repo}
}

func (p SectionsService) GetSections() map[int]models.Sections {
	return p.repo.GetSections()
}
