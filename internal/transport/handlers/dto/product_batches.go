package dto

import "time"

type ProductBatchesDto struct {
	Id                 int       `json:"id"`
	BatchNumber        string    `json:"batch_number" validate:"required"`
	CurrentQuantity    int       `json:"current_quantity" validate:"required"`
	CurrentTemperature float64   `json:"current_temperature" validate:"required"`
	DueDate            time.Time `json:"due_date" validate:"required"`
	InitialQuantity    int       `json:"initial_quantity" validate:"required"`
	ManufacturingDate  time.Time `json:"manufacturing_date" validate:"required"`
	ManufacturingHour  time.Time `json:"manufacturing_hour" validate:"required"`
	MinumumTemperature float64   `json:"minimum_temperature" validate:"required"`
	ProductId          int       `json:"product_id"`
	SectionId          int       `json:"section_id"`
}

type SectionsProductsDto struct {
	Id                  int     `json:"id"`
	Section_number      string  `json:"section_number"`
	Current_temperature float64 `json:"current_temperature"`
	Minimum_temperature float64 `json:"minimum_temperature"`
	Current_capacity    int     `json:"current_capacity"`
	Minimum_capacity    int     `json:"minimum_capacity"`
	Maximum_capacity    int     `json:"maximum_capacity"`
	Warehouse_id        int     `json:"warehouse_id"`
	Product_type_id     int     `json:"product_type"`
	ProductCount        int     `json:"product_count"`
}

type ProductBatchesDtoReport struct {
	ProductId int `json:"product_id"`
	SectionId int `json:"section_id"`
}
