package dto

import "errors"

type WarehouseDto struct {
	Id                 int    `json:"id"`
	WarehouseCode      string `json:"warehouse_code"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimumCapacity    int    `json:"minimun_capacity"`
	MinimumTemperature int    `json:"minimun_temperature"`
	LocalityId         int    `json:"locality_id"`
}

func (item *WarehouseDto) Validate() error {
	switch {
	case item.Address == "":
		return errors.New("address is required")
	case item.Telephone == "":
		return errors.New("phone number is required")
	case item.WarehouseCode == "":
		return errors.New("warehouse code is required")
	case item.MinimumCapacity < 0:
		return errors.New("capacity can not be less than cero")
	case item.MinimumTemperature < -18:
		return errors.New("temperature can not be less than -18 degrees")
	case item.MinimumTemperature > 15:
		return errors.New("temperature can not be greater than 15 degrees")
	case item.LocalityId <= 0:
		return errors.New("enter a valid locality id")
	default:
		return nil
	}
}
