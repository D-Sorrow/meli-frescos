package mappers

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func MapperToSectionDto(sections models.Sections) dto.SectionsDto {

	return dto.SectionsDto{
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

func MapperToSectionsDto(sections map[int]models.Sections) []dto.SectionsDto {

	var sectionstSliceDto []dto.SectionsDto
	for _, p := range sections {
		sectionsDto := MapperToSectionDto(p)
		sectionstSliceDto = append(sectionstSliceDto, sectionsDto)
	}
	return sectionstSliceDto
}

func MapperToSectionsModel(sections dto.SectionsDto) models.Sections {
	return models.Sections{
		Id:                  sections.Id,
		Warehouse_id:        sections.Warehouse_id,
		Product_type_id:     sections.Product_type_id,
		Section_number:      sections.Section_number,
		Current_temperature: sections.Current_temperature,
		Minimum_temperature: sections.Minimum_temperature,
		Current_capacity:    sections.Current_capacity,
		Minimum_capacity:    sections.Minimum_capacity,
		Maximum_capacity:    sections.Maximum_capacity,
	}
}
