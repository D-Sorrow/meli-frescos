package mappers

import (
	"fmt"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/entity"
)

func JsonSectionsModel(sectionsJSON []entity.SectionsJSON) map[int]models.Sections {
	sections := make(map[int]models.Sections)
	fmt.Println("dolor desde sections mapper")
	for _, sectionsJSON := range sectionsJSON {
		sections[sectionsJSON.Id] = models.Sections{
			Id:                  sectionsJSON.Id,
			Section_number:      sectionsJSON.Section_number,
			Current_temperature: sectionsJSON.Current_temperature,
			Minimum_temperature: sectionsJSON.Minimum_temperature,
			Current_capacity:    sectionsJSON.Current_capacity,
			Minimum_capacity:    sectionsJSON.Minimum_capacity,
			Maximum_capacity:    sectionsJSON.Maximum_capacity,
			Warehouse_id:        sectionsJSON.Warehouse_id,
			Product_type_id:     sectionsJSON.Product_type_id,
		}
	}

	return sections
}
