package transactions

import (
	"errors"

	"github.com/leonardogregoriocs/internal/contract"
	"github.com/leonardogregoriocs/internal/domain/client"
	internalerrors "github.com/leonardogregoriocs/internal/internalErrors"
)

type Service struct {
	Repository       Repository
	ClientRepository client.Repository
}

func (s *Service) CreateNewTransactions(newTransaction contract.NewTransaction) (int, error) {
	clientIsValid, _ := s.ClientRepository.GetBy(newTransaction.AccountID)

	if clientIsValid.ID == newTransaction.AccountID {
		transactions, err := NewTransaction(newTransaction.AccountID, newTransaction.OperationTypeID, newTransaction.Amout)
		if err != nil {
			return 1, err
		}

		err = s.Repository.Save(transactions)
		if err != nil {
			return 1, internalerrors.ErrInternal
		}

		return transactions.ID, nil
	}

	return 1, errors.New("Client not exist")
}
