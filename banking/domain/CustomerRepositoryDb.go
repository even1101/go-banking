package domain

import (
	"banking-lib/errs"
	"banking-lib/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found.")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error.")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	var err error
	customers := make([]Customer, 0)
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	err = d.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, err
	}

	return customers, nil
}

func NewCustomerRepositoryDb(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client}
}
