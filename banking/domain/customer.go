package domain

import "banking/errs"

type Customer struct {
	Id      string `db:"customer_id" json:"id" xml:"id"`
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip"`
	Birth   string `db:"date_of_birth" json:"birth" xml:"birth"`
	Status  string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
