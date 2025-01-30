package loader

import (
	"encoding/json"
	"os"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
)

func NewBuyerJSONFile(path string) repository.BuyerLoader {
	return &BuyerJSONFile{
		path: path,
	}
}

type BuyerJSONFile struct {
	path string
}

func (l *BuyerJSONFile) Load() (buyers map[int]models.Buyer, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var BuyersJSON []dto.BuyerDTO
	err = json.NewDecoder(file).Decode(&BuyersJSON)
	if err != nil {
		return
	}

	buyers = make(map[int]models.Buyer)
	for _, buyer := range BuyersJSON {
		buyers[buyer.ID] = models.Buyer{
			ID: buyer.ID,
			BuyerAttributes: models.BuyerAttributes{
				CardNumberID: buyer.CardNumberID,
				FirstName:    buyer.FirstName,
				LastName:     buyer.LastName,
			},
		}
	}

	return
}
