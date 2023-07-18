package transactions

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	AccountID       = 1
	OperationTypeID = 3
	Amout           = -100.0
)

func TestNewTransaction(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(reflect.TypeOf(AccountID).Kind(), reflect.Int)
	assert.Equal(reflect.TypeOf(OperationTypeID).Kind(), reflect.Int)
	assert.Equal(reflect.TypeOf(Amout).Kind(), reflect.Float64)
}

func TestDateTransactionsMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	newTransaction, _ := NewTransaction(AccountID, OperationTypeID, Amout)

	assert.Greater(newTransaction.EventDate, now)
}

func TestNewTransactionWhenOperationTypeIsOneOrTwo(t *testing.T) {
	assert := assert.New(t)

	newTransaction1 := &Transactions{
		AccountID:       1,
		OperationTypeID: 1,
		Amout:           -100,
	}

	newTransaction2 := &Transactions{
		AccountID:       1,
		OperationTypeID: 2,
		Amout:           -100,
	}

	newTransaction1, _ = NewTransaction(newTransaction1.AccountID, newTransaction1.OperationTypeID, newTransaction1.Amout)
	newTransaction2, _ = NewTransaction(newTransaction2.AccountID, newTransaction2.OperationTypeID, newTransaction2.Amout)

	assert.NotNil(newTransaction1)
	assert.NotNil(newTransaction2)
}

func TestNewTransactionWhenOperationTypeIsFour(t *testing.T) {
	assert := assert.New(t)

	newTransaction1 := &Transactions{
		AccountID:       1,
		OperationTypeID: 4,
		Amout:           100,
	}

	newTransaction1, _ = NewTransaction(newTransaction1.AccountID, newTransaction1.OperationTypeID, newTransaction1.Amout)

	assert.NotNil(newTransaction1)
}
