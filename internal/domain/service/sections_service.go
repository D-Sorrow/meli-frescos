package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"

	er "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
)

type SectionsService struct {
	repo repository.SectionsRepository
}

func NewSectionsService(repo repository.SectionsRepository) *SectionsService {
	return &SectionsService{repo: repo}
}

func (s SectionsService) GetSections() map[int]models.Sections {
	return s.repo.GetSections()
}

func (s SectionsService) GetSectionsById(id int) (models.Sections, error) {
	sections, errGet := s.repo.GetSectionsById(id)
	if errGet != nil {
		return models.Sections{}, er.ErrSectionsNotFound
	}
	return sections, nil
}

func (s SectionsService) SaveSections(sectionstSave models.Sections) error {
	errSave := s.repo.SaveSections(sectionstSave)
	if errSave != nil {
		return er.Error{Code: er.CodeNotFound, Message: er.SectionsIsAlreadyExist}
	}
	return nil
}

func (s SectionsService) DeleteSections(id int) error {
	errDeleate := s.repo.DeleteSections(id)
	if errDeleate != nil {
		return er.Error{Code: er.CodeNotFound, Message: er.SectionsNotFound}
	}
	return nil
}

/*func (s SectionsService) UpdateSections(id int, sectionsPatch models.SectionsPatch) (models.Sections, error) {

	value, err := s.repo.UpdateSections(id, sectionsPatch)
	if err != nil {
		return models.Sections{}, er.ErrSectionsNotFound
	}
	return value, nil

}
*/
