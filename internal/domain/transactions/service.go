package transactions

import (
	"github.com/leonardogregoriocs/internal/contract"
	internalerrors "github.com/leonardogregoriocs/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) CreateNewTransactions(newTransaction contract.NewTransaction) (int, error) {

	transactions, err := NewTransaction(newTransaction.AccountID, newTransaction.OperationTypeID, newTransaction.Amout)
	if err != nil {
		return 1, err
	}

	err = s.Repository.Save(transactions)
	if err != nil {
		return 1, internalerrors.ErrInternal
	}

	return transactions.TransactionsID, nil
}
