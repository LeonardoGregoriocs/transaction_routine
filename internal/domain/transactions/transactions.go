package transactions

import (
	"time"
)

type Transactions struct {
	TransactionsID  int
	AccountID       int
	OperationTypeID int
	Amout           float64
	EventDate       time.Time
}

func NewTransaction(accountID int, operationTypeID int, amout float64) (*Transactions, error) {
	return &Transactions{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amout:           amout,
		EventDate:       time.Now(),
	}, nil
}
