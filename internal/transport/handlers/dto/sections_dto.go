package dto

const IdInvalidSections = "ID must be positive"
const DescriptionInvalidSections = "Description must not be empty."
const ExpirationInvalidSections = "Rates must not be negative"

type SectionsDto struct {
	Id                  int     `json:"id"`
	Section_number      string  `json:"section_number"`
	Current_temperature float64 `json:"current_temperature"`
	Minimum_temperature float64 `json:"minimum_temperature"`
	Current_capacity    int     `json:"current_capacity"`
	Minimum_capacity    int     `json:"minimum_capacity"`
	Maximum_capacity    int     `json:"maximum_capacity"`
	Warehouse_id        int     `json:"warehouse_id"`
	Product_type_id     int     `json:"product_type"`
}

type SectionsUpdateDto struct {
	Id                  *int     `json:"id"`
	Section_number      *string  `json:"section_number"`
	Current_temperature *float64 `json:"current_temperature"`
	Minimum_temperature *float64 `json:"minimum_temperature"`
	Current_capacity    *int     `json:"current_capacity"`
	Minimum_capacity    *int     `json:"minimum_capacity"`
	Maximum_capacity    *int     `json:"maximum_capacity"`
	Warehouse_id        *int     `json:"warehouse_id"`
	Product_type_id     *int     `json:"product_type"`
}

func (p *SectionsDto) Validate() error {

	return nil
}
