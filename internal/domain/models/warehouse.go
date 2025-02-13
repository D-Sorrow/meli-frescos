package models

type Warehouse struct {
	Id                 int
	WarehouseCode      string
	Address            string
	Telephone          string
	MinimunCapacity    int
	MinimunTemperature int
	LocalityId         int
}
