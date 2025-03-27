package entities

import "time"

type ProductBatchesEntity struct {
	Id                 int
	BatchNumber        string
	CurrentQuantity    int
	CurrentTemperature float64
	DueDate            time.Time
	InitialQuantity    int
	ManufacturingDate  time.Time
	ManufacturingHour  time.Time
	MinumumTemperature float64
	ProductId          int
	SectionId          int
}

func (p *ProductBatchesEntity) GetByIdQuery(id int) (query string, args []interface{}) {
	query = `SELECT id,
		batch_number,
		current_quantity,
		current_temperature,
		due_date,
		initial_quantity,
		manufacturing_date,
		manufacturing_hour,
		minimum_temperature,
		product_id,
		section_id
	FROM product_batches
	WHERE id =?
	LIMIT 1;`
	args = []interface{}{id}

	return
}

func (p *ProductBatchesEntity) GetCreateQuery() (query string, args []interface{}) {
	query = `INSERT INTO product_batches (
		batch_number,
		current_quantity,
		current_temperature,
		due_date,
		initial_quantity,
		manufacturing_date,
		manufacturing_hour,
		minimum_temperature,
		product_id,
		section_id
	) VALUES (?,?,?,?,?,?,?,?,?,?);`
	args = []interface{}{p.BatchNumber, p.CurrentQuantity, p.CurrentTemperature, p.DueDate, p.InitialQuantity, p.ManufacturingDate, p.ManufacturingHour, p.ManufacturingHour, p.MinumumTemperature, p.ProductId, p.SectionId}

	return
}
