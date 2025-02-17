package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	entity "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
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
