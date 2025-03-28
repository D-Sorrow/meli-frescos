package loader

import (
	"encoding/json"
	"os"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/loader/entity"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/loader/mappers"
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
