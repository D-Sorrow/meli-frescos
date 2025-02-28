package dto

import "errors"

type CarrierDto struct {
	Id          int    `json:"id"`
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

func (item *CarrierDto) Validate() error {
	switch {
	case item.Cid == "":
		return errors.New("carrier cid is required")
	case item.CompanyName == "":
		return errors.New("company name is required")
	case item.Address == "":
		return errors.New("address is required")
	case item.Telephone == "":
		return errors.New("phone number is required")
	case item.LocalityId <= 0:
		return errors.New("enter a valid locality id")
	default:
		return nil
	}
}
