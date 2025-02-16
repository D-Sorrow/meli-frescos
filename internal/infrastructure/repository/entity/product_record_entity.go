package entity

type ProductRecordEntity struct {
	ProductId      int
	SalePrice      float64
	PurchasePrice  float64
	LastUpdateTime string
}

func (pre ProductRecordEntity) SaveProductRecord() string {

	return "INSERT INTO melifresh.product_records" +
		"(last_update_date, purchase_price, sale_price, product_id)" +
		"VALUES (?, ?, ?, ?)"
}
