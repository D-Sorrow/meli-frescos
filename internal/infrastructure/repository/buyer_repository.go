package repository

import (
	"database/sql"
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
	"github.com/go-sql-driver/mysql"
)

type BuyerRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}

func (b *BuyerRepository) GetAll() (buyers []entities.BuyerEntity, err error) {
	buyers = make([]entities.BuyerEntity, 0)

	query, args := (&entities.BuyerEntity{}).GetAllQuery()

	rows, err := b.db.Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var buyer entities.BuyerEntity
		err = rows.Scan(&buyer.ID,
			&buyer.CardNumberID,
			&buyer.FirstName,
			&buyer.LastName,
		)

		if err != nil {
			return
		}

		buyers = append(buyers, buyer)
	}

	if len(buyers) == 0 {
		err = repository.ErrNoRegisteredBuyersYet
		return
	}

	return
}

func (b *BuyerRepository) GetById(id int) (buyer entities.BuyerEntity, err error) {
	query, args := buyer.GetByIdQuery(id)

	err = b.db.QueryRow(query, args...).Scan(&buyer.ID,
		&buyer.CardNumberID,
		&buyer.FirstName,
		&buyer.LastName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repository.ErrBuyerNotFoundWithID
			return
		}
	}

	return
}

func (b *BuyerRepository) Create(buyer entities.BuyerEntity) (newBuyer entities.BuyerEntity, err error) {
	query, args := buyer.GetCreateQuery()

	result, err := b.db.Exec(query, args...)

	if err != nil {
		var mySqlErr *mysql.MySQLError
		if errors.As(err, &mySqlErr) && mySqlErr.Number == 1062 {
			err = repository.ErrDuplicateCardNumberID
			return
		}

		return
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return
	}

	newBuyer, err = b.GetById(int(lastId))

	return
}

func (b *BuyerRepository) Patch(id int, buyerToPatch entities.BuyerEntity) (updatedBuyer entities.BuyerEntity, err error) {
	updatedBuyer, err = b.GetById(id)

	if err != nil {
		return
	}

	if buyerToPatch.CardNumberID != nil {
		updatedBuyer.CardNumberID = buyerToPatch.CardNumberID
	}

	if buyerToPatch.FirstName != nil {
		updatedBuyer.FirstName = buyerToPatch.FirstName
	}

	if buyerToPatch.LastName != nil {
		updatedBuyer.LastName = buyerToPatch.LastName
	}

	query, args := updatedBuyer.GetUpdateQuery()

	_, err = b.db.Exec(query, args...)

	if err != nil {
		var mySqlErr *mysql.MySQLError
		if errors.As(err, &mySqlErr) && mySqlErr.Number == 1062 {
			err = repository.ErrDuplicateCardNumberID
			return
		}

		return
	}

	return
}

func (b *BuyerRepository) Delete(id int) (err error) {
	query, args := (&entities.BuyerEntity{}).GetDeleteQuery(id)

	result, err := b.db.Exec(query, args...)

	if err != nil {
		var mySqlErr *mysql.MySQLError
		if errors.As(err, &mySqlErr) && mySqlErr.Number == 1451 {
			err = repository.ErrCannotDeleteBuyerWithOrders
			return
		}

		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		err = repository.ErrBuyerNotFoundWithID
		return
	}

	return
}

func (b *BuyerRepository) GetReportPurchaseOrders(buyerID *int) (report []entities.ReportPurchaseOrdersEntity, err error) {
	report = make([]entities.ReportPurchaseOrdersEntity, 0)

	query, args := (&entities.ReportPurchaseOrdersEntity{}).GetReportPurchaseOrdersQuery(buyerID)

	rows, err := b.db.Query(query, args...)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var _report entities.ReportPurchaseOrdersEntity
		err = rows.Scan(
			&_report.ID,
			&_report.CardNumberID,
			&_report.FirstName,
			&_report.LastName,
			&_report.PurchaseOrdersCount,
		)
		if err != nil {
			return
		}
		report = append(report, _report)
	}

	if len(report) == 0 {
		err = repository.ErrBuyerNotFoundOrBuyerHasNoOrders
		return
	}

	return
}
