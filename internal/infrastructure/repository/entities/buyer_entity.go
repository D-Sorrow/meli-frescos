package entities

type BuyerEntity struct {
	ID           int
	CardNumberID *string
	FirstName    *string
	LastName     *string
}

type ReportPurchaseOrdersEntity struct {
	ID                  int
	CardNumberID        string
	FirstName           string
	LastName            string
	PurchaseOrdersCount int
}

func (p *BuyerEntity) GetAllQuery() (query string, args []interface{}) {
	query = `SELECT  id,
		card_number_id,
		first_name,
		last_name
	FROM buyers;`
	args = []interface{}{}

	return
}

func (p *BuyerEntity) GetByIdQuery(id int) (query string, args []interface{}) {
	query = `SELECT id,
        card_number_id,
        first_name,
        last_name
	FROM buyers
	WHERE id =?
	LIMIT 1;`
	args = []interface{}{id}

	return
}

func (p *BuyerEntity) GetCreateQuery() (query string, args []interface{}) {
	query = `INSERT INTO buyers (
		card_number_id,
		first_name,
		last_name)
	VALUES (?,?,?);`
	args = []interface{}{&p.CardNumberID, &p.FirstName, &p.LastName}

	return
}

func (p *BuyerEntity) GetUpdateQuery() (query string, args []interface{}) {
	query = `UPDATE buyers SET
		card_number_id = ?,
		first_name =?,
        last_name =?
	WHERE id =?;`
	args = []interface{}{&p.CardNumberID, &p.FirstName, &p.LastName, &p.ID}

	return
}

func (p *BuyerEntity) GetDeleteQuery(id int) (query string, args []interface{}) {
	query = `DELETE FROM buyers WHERE id =?;`
	args = []interface{}{id}

	return
}

func (p *ReportPurchaseOrdersEntity) GetReportPurchaseOrdersQuery(buyerID *int) (query string, args []interface{}) {
	query = `SELECT
		b.id,
		b.card_number_id,
		b.first_name,
		b.last_name,
		COUNT(po.id) AS purchase_orders_count
	FROM buyers b
	LEFT JOIN purchase_orders po ON b.id = po.buyer_id
	WHERE (COALESCE(?, b.id) = b.id)
	GROUP BY b.id;`

	if buyerID == nil {
		args = []interface{}{nil}
	} else {
		args = []interface{}{&buyerID}
	}

	return
}
