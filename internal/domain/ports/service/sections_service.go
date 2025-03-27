package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type SectionsService interface {
	GetSections() map[int]models.Sections
	GetSectionsById(id int) (models.Sections, error)
	SaveSections(product models.Sections) error
	DeleteSections(id int) error
	//UpdateSections(id int, sectionsPatch models.SectionsPatch) (models.Sections, error)
}
