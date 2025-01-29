package service

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type SectionsService interface {
	GetSections() map[int]models.Sections
}
