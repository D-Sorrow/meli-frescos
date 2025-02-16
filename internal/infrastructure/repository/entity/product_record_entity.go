package entity

import "fmt"

type ProductRecordEntity struct {
	ProductId      int
	SalePrice      float64
	PurchasePrice  float64
	LastUpdateTime string
}

var (
	partGroupBy = "GROUP BY rec.product_id"
	partWhere   = "WHERE rec.product_id = "
)

func (pre ProductRecordEntity) SaveProductRecord() string {

	return "INSERT INTO melifresh.product_records" +
		"(last_update_date, purchase_price, sale_price, product_id)" +
		"VALUES (?, ?, ?, ?)"
}

func (pre ProductRecordEntity) GetRecord(productId int) string {
	part := fmt.Sprint("SELECT pro.id, pro.description, COUNT(*)\n" +
		"FROM product_records rec\n" +
		"INNER JOIN products  pro\n" +
		"ON rec.product_id = pro.id\n")
	if productId == 0 {
		return fmt.Sprint(part, partGroupBy)
	}
	return fmt.Sprint(part, partWhere, productId)
}
