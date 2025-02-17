package models

type Locality struct {
	Id           int
	ZipCode      string
	Name         string
	ProvinceName string
	CountryName  string
}

type LocalityPatch struct {
	ZipCode    *string
	Name       *string
	ProvinceId *int
}

type LocalitySellers struct {
	LocalityId   *int
	ZipCode      *string
	Name         *string
	SellersCount *int
}

type LocalityCarriers struct {
	LocalityId    *int
	ZipCode       *string
	Name          *string
	CarriersCount *int
}
