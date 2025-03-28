package service

import "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"

type SectionsService interface {
	GetSections() map[int]models.Sections
	GetSectionsById(id int) (models.Sections, error)
	SaveSections(product models.Sections) error
	DeleteSections(id int) error
	//UpdateSections(id int, sectionsPatch models.SectionsPatch) (models.Sections, error)
}
