package models

type Sections struct {
	Id                  int
	Warehouse_id        int
	Product_type_id     int
	Section_number      string
	Current_temperature float64
	Minimum_temperature float64
	Current_capacity    int
	Minimum_capacity    int
	Maximum_capacity    int
}

type SectionsPatch struct {
	Id                  *int
	Warehouse_id        *int
	Product_type_id     *int
	Section_number      *string
	Current_temperature *float64
	Minimum_temperature *float64
	Current_capacity    *int
	Minimum_capacity    *int
	Maximum_capacity    *int
}

type ReportProductBatch struct {
	Id                  int
	Warehouse_id        int
	Product_type_id     int
	Section_number      string
	Current_temperature float64
	Minimum_temperature float64
	Current_capacity    int
	Minimum_capacity    int
	Maximum_capacity    int
	ProductBatchesCount int
}

/*type SectionsAtributePatch struct {
	Section_number      *string
	Current_temperature *float64
	Minimum_temperature *float64
	Current_capacity    *int
	Minimum_capacity    *int
	Maximum_capacity    *int
}*/
