package transactions

import (
	"errors"
	"time"
)

type Transactions struct {
	ID              int
	AccountID       int
	OperationTypeID int
	Amout           float64
	EventDate       time.Time
	Balance         float64
}

func NewTransaction(accountID int, operationTypeID int, amout float64) (*Transactions, error) {
	if accountID == 0 || operationTypeID == 0 {
		return nil, errors.New("Incomplete data")
	}

	if operationTypeID == 4 && amout < 0.01 || operationTypeID < 4 && amout > 0 {
		return nil, errors.New("Operation Type invalid")
	}

	return &Transactions{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amout:           amout,
		EventDate:       time.Now(),
	}, nil
}
