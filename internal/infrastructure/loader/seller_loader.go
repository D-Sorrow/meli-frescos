package loader

import (
	"encoding/json"
	"os"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

type SellerJSONFile struct {
	path string
}

func NewSellerJSONFile(path string) *SellerJSONFile {
	return &SellerJSONFile{path: path}
}

func (l *SellerJSONFile) Load() (sellers map[int]models.Seller, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var SellerJSON []dto.SellerDto
	err = json.NewDecoder(file).Decode(&SellerJSON)
	if err != nil {
		return
	}

	sellers = make(map[int]models.Seller)
	for _, seller := range SellerJSON {
		sellers[seller.Id] = models.Seller{
			Id:          seller.Id,
			Cid:         seller.Cid,
			CompanyName: seller.CompanyName,
			Address:     seller.Address,
			Telephone:   seller.Telephone,
		}
	}

	return
}
