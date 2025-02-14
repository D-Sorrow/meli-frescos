package dto

type LocalityDto struct {
	Id           int    `json:"id"`
	ZipCode      string `json:"zip_code" validate:"required"`
	LocalityName string `json:"locality_name" validate:"required"`
	ProvinceId   int    `json:"province_id" validate:"required"`
}

type LocalitySellersDto struct {
	LocalityId   int    `json:"locality_id"`
	ZipCode      string `json:"zip_code"`
	Name         string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}
