package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/entity"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/mappers"
)

type ProductJSONFile struct {
	path string
}

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{path: path}
}

func (p *ProductJSONFile) Load() (products map[int]models.Product, err error) {
	file, err := os.Open(p.path)
	if err != nil {
		fmt.Println("hola desde loader")
		return
	}
	defer file.Close()

	var productsJSON []entity.ProductJSON
	err = json.NewDecoder(file).Decode(&productsJSON)
	if err != nil {
		return
	}

	products = mappers.JsonToProductModel(productsJSON)

	return products, nil
}
