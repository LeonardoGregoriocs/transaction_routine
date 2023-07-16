package database

import (
	"github.com/leonardogregoriocs/internal/domain/transactions"
	"gorm.io/gorm"
)

type TransactionsRepository struct {
	Db *gorm.DB
}

func (t *TransactionsRepository) Save(transactions *transactions.Transactions) error {
	result := t.Db.Create(transactions)
	return result.Error
}
