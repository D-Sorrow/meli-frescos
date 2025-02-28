package entities

type SectionsEntity struct {
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

type ReportProductBatchEntity struct {
	Id                  int
	Section_number      string
	Current_temperature float64
	Minimum_temperature float64
	Current_capacity    int
	Minimum_capacity    int
	Maximum_capacity    int
	Warehouse_id        int
	Product_type_id     int
	ProductBatchesCount int
}
