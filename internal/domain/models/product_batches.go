package models

import "time"

type ProductBatches struct {
	Id                 int
	BatchNumber        string
	CurrentQuantity    int
	CurrentTemperature float64
	DueDate            time.Time
	InitialQuantity    int
	ManufacturingDate  time.Time
	ManufacturingHour  time.Time
	MinumumTemperature float64
	ProductId          int
	SectionId          int
}

type ProductBatches2FKs struct {
	ProductId int
	SectionId int
}

type ProductBatches2Attributes struct {
	Id                 int
	BatchNumber        string
	CurrentQuantity    int
	CurrentTemperature float64
	DueDate            time.Time
	InitialQuantity    int
	ManufacturingDate  time.Time
	ManufacturingHour  time.Time
	MinumumTemperature float64
}

type ProductBatches2AttributesFks struct {
	ProductBatches2Attributes
	ProductBatches2FKs
}

type ProductBatches3 struct {
	ID int
	ProductBatches2Attributes
	ProductBatches2FKs
}

type SectionsProductsBatchs struct {
	Id                  *int
	Warehouse_id        *int
	Product_type_id     *int
	Section_number      *string
	Current_temperature *float64
	Minimum_temperature *float64
	Current_capacity    *int
	Minimum_capacity    *int
	Maximum_capacity    *int
	ProductCount        *int
}
