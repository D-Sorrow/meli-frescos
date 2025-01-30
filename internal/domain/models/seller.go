package models

type Seller struct {
	Id          int
	Cid         int
	CompanyName string
	Address     string
	Telephone   string
}


type SellerPatch struct {
	Cid         *int
	CompanyName *string
	Address     *string
	Telephone   *string
}
