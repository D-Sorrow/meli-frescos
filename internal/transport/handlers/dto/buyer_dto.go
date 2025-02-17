package dto

type ReportPurchaseOrdersDTO struct {
	ID                  int    `json:"id"`
	CardNumberID        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

type BuyerDTO struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerCreateDTO struct {
	CardNumberID *string `json:"card_number_id" validate:"required,cardnumber"`
	FirstName    *string `json:"first_name" validate:"required,str_not_empty"`
	LastName     *string `json:"last_name" validate:"required,str_not_empty"`
}

type BuyerPatchDTO struct {
	CardNumberID *string `json:"card_number_id" validate:"omitempty,cardnumber"`
	FirstName    *string `json:"first_name" validate:"omitempty,str_not_empty"`
	LastName     *string `json:"last_name" validate:"omitempty,str_not_empty"`
}
