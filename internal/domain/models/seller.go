package models

type Seller struct {
	Id          int
	Cid         int
	CompanyName string
	Address     string
	Telephone   string
	LocalityId  int
}

type SellerPatch struct {
	Cid         *int
	CompanyName *string
	Address     *string
	Telephone   *string
	LocalityId  *int
}
