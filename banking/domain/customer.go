package domain

import (
	"banking-lib/errs"
	"banking/dto"
)

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	ZipCode string
	Birth   string `db:"date_of_birth"`
	Status  string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:      c.Id,
		Name:    c.Name,
		City:    c.City,
		ZipCode: c.ZipCode,
		Birth:   c.Birth,
		Status:  c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
