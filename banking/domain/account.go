package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}
