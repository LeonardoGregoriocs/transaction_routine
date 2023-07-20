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
	client, _ := s.ClientRepository.GetBy(newTransaction.AccountID)

	if client.ID == newTransaction.AccountID {
		transactions, err := NewTransaction(newTransaction.AccountID, newTransaction.OperationTypeID, newTransaction.Amout)
		if err != nil {
			return 1, err
		}

		if newTransaction.OperationTypeID == 4 {
			newCredit := client.AvailableCreditLimit + newTransaction.Amout
			s.ClientRepository.UpdateCredit(newCredit, newTransaction.AccountID)
		}

		if newTransaction.OperationTypeID != 4 {
			newCredit, err := s.ValidateCredit(newTransaction.Amout, client.AvailableCreditLimit)
			if err != nil {
				return 1, err
			}
			s.ClientRepository.UpdateCredit(newCredit, newTransaction.AccountID)
		}

		err = s.Repository.Save(transactions)
		if err != nil {
			return 1, internalerrors.ErrInternal
		}

		return transactions.ID, nil
	}

	return 1, errors.New("Client not exist")
}

func (s *Service) ValidateCredit(amout float64, credit float64) (float64, error) {
	newCreditLitit := (credit + amout)

	if newCreditLitit < 0 {
		return credit, errors.New("Unavailable balance")
	} else {
		return newCreditLitit, nil
	}
}
