package dto

type SellerDto struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required"`
}

type SellerUpdateDto struct {
	Cid         *int    `json:"cid" validate:"omitempty"`
	CompanyName *string `json:"company_name" validate:"omitempty"`
	Address     *string `json:"address" validate:"omitempty"`
	Telephone   *string `json:"telephone" validate:"omitempty"`
}
