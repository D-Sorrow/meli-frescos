package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type SectionsRepository interface {
	GetSections() map[int]models.Sections
}
