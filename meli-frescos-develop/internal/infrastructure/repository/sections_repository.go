package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type SectionsRepository struct {
	sectionsMap map[int]models.Sections
}

func NewSectionsRepository(sectionsMap map[int]models.Sections) *SectionsRepository {
	return &SectionsRepository{sectionsMap: sectionsMap}
}

func (p SectionsRepository) GetSections() map[int]models.Sections {
	return p.sectionsMap
}
