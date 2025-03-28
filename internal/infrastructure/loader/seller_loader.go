package loader

import (
	"encoding/json"
	"os"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
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
