package client

import (
	"github.com/leonardogregoriocs/internal/contract"
	internalerrors "github.com/leonardogregoriocs/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) CreateClient(newClient contract.NewClient) (int, error) {

	client, err := NewClient(newClient.DocumentNumber)
	if err != nil {
		return 1, err
	}

	err = s.Repository.Save(client)
	if err != nil {
		return 1, internalerrors.ErrInternal
	}

	s.Repository.Save(client)

	return client.ID, nil
}

func (s *Service) GetBy(id int) (*contract.ClientResponse, error) {

	client, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	return &contract.ClientResponse{
		ID:             client.ID,
		DocumentNumber: client.DocumentNumber,
	}, nil
}
