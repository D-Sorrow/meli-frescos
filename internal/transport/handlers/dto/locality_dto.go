package dto

type LocalityDto struct {
	Id           int    `json:"id"`
	ZipCode      string `json:"zip_code" validate:"required"`
	LocalityName string `json:"locality_name" validate:"required"`
	ProvinceId   int    `json:"province_id" validate:"required"`
}
