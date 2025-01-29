package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/entity"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/mappers"
)

type SectionsJSONFile struct {
	path string
}

func NewSectionsJSONFile(path string) *SectionsJSONFile {
	return &SectionsJSONFile{path: path}
}

func (p *SectionsJSONFile) Load() (sections map[int]models.Sections, err error) {

	fmt.Println("welcome")
	file, err := os.Open(p.path)
	if err != nil {
		fmt.Println("hola desde sections loader")
		return
	}
	defer file.Close()

	var sectionsJSON []entity.SectionsJSON
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return
	}

	sections = mappers.JsonSectionsModel(sectionsJSON)

	return sections, nil
}
