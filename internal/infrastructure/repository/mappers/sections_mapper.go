package mappers

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	entity "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository/entities"
)

func ToSectionsEntity(sections models.Sections) entity.SectionsEntity {
	return entity.SectionsEntity{
		Id:                  sections.Id,
		Section_number:      sections.Section_number,
		Current_temperature: sections.Current_temperature,
		Minimum_temperature: sections.Minimum_temperature,
		Current_capacity:    sections.Current_capacity,
		Minimum_capacity:    sections.Minimum_capacity,
		Maximum_capacity:    sections.Maximum_capacity,
		Warehouse_id:        sections.Warehouse_id,
		Product_type_id:     sections.Product_type_id,
	}
}
