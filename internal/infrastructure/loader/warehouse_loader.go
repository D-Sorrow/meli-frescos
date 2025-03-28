package loader

import (
	"encoding/json"
	"os"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/loader/entity"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/loader/mappers"
)

type WarehouseJSONFile struct {
	path string
}

func NewWarehouseJSONFile(path string) *WarehouseJSONFile {
	return &WarehouseJSONFile{path: path}
}

func (w *WarehouseJSONFile) Load() (warehouses map[int]models.Warehouse, err error) {
	file, err := os.Open(w.path)
	if err != nil {
		return
	}
	defer file.Close()

	var warehousesJSON []entity.WarehouseJSON
	err = json.NewDecoder(file).Decode(&warehousesJSON)
	if err != nil {
		return
	}

	warehouses = mappers.JsonToWarehouseModel(warehousesJSON)

	return warehouses, nil
}
