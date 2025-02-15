package loader

import (
	"encoding/json"
	"os"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/entity"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader/mappers"
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
