package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	repository_errors "github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
	"github.com/go-sql-driver/mysql"
)

type SellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) *SellerRepository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) GetSellers() (v map[int]models.Seller, err error) {
	v = make(map[int]models.Seller)

	rows, err := r.db.Query("SELECT id, cid, company_name, address, telephone, locality_id from sellers")
	if err != nil {
		log.Print(err)
		return nil, repository_errors.ErrSellerNotFound
	}

	for rows.Next() {
		var seller models.Seller
		err := rows.Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityId)
		if err != nil {
			log.Print(err)
			return nil, repository_errors.ErrSellerNotFound
		}
		v[seller.Id] = seller
	}

	if len(v) <= 0 {
		return nil, repository_errors.ErrSellerNotFound
	}
	return
}

func (r *SellerRepository) GetSellerById(id int) (models.Seller, error) {
	var seller models.Seller

	row := r.db.QueryRow("SELECT id, cid, company_name, address, telephone, locality_id from sellers where id = ?", id)
	if err := row.Err(); err != nil {
		log.Print(err)
		return models.Seller{}, repository_errors.ErrSellerNotFound
	}

	err := row.Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityId)
	if errors.Is(err, sql.ErrNoRows) {
		log.Print(err)
		return models.Seller{}, repository_errors.ErrSellerNotFound
	}

	return seller, nil
}

func (r *SellerRepository) CreateSeller(seller models.Seller) (models.Seller, error) {

	result, err := r.db.Exec("INSERT INTO sellers (cid,company_name,address,telephone,locality_id) values (?,?,?,?,?)", seller.Cid, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityId)

	if err != nil {
		log.Print(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return models.Seller{}, repository_errors.ErrSellerAlreadyExists
		}
		return models.Seller{}, repository_errors.ErrSellerCannotBeCreated
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return models.Seller{}, repository_errors.ErrSellerCannotBeCreated
	}

	seller.Id = int(lastInsertId)
	return seller, nil

}

func (r *SellerRepository) UpdateSeller(id int, seller models.SellerPatch) (models.Seller, error) {
	value, err := r.GetSellerById(id)
	if err != nil {
		return models.Seller{}, err
	}

	if seller.Cid != nil {
		value.Cid = *seller.Cid
	}
	if seller.Address != nil {
		value.Address = *seller.Address
	}
	if seller.CompanyName != nil {
		value.CompanyName = *seller.CompanyName
	}
	if seller.Telephone != nil {
		value.Telephone = *seller.Telephone
	}
	if seller.LocalityId != nil {
		value.LocalityId = *seller.LocalityId
	}

	_, err = r.db.Exec("UPDATE sellers SET cid=?, address=?, company_name=?, telephone=?, locality_id=? where id = ?", value.Cid, value.Address, value.CompanyName, value.Telephone, value.LocalityId, value.Id)
	if err != nil {
		log.Print(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return models.Seller{}, repository_errors.ErrSellerAlreadyExists
		}
		return models.Seller{}, repository_errors.ErrSellerCannotBeUpdated
	}

	return value, nil

}

func (r *SellerRepository) DeleteSeller(id int) error {
	_, err := r.GetSellerById(id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM sellers WHERE id = ?", id)

	if err != nil {
		log.Print(err)
		return repository_errors.ErrSellerCannotBeDeleted
	}

	return nil
}
