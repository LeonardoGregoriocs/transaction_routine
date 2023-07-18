package client

import (
	"testing"
	"time"

	"github.com/leonardogregoriocs/internal/contract"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(client *Client) error {
	args := r.Called(client)
	return args.Error(0)
}

func (r *repositoryMock) GetBy(id int) (*Client, error) {
	args := r.Called(id)
	return &Client{
		DocumentNumber:   "12365497899",
		DateCreateUpdate: time.Now(),
	}, args.Error(0)
}

func (r *repositoryMock) GetClientByDocumentNumber(documentNumber string) (string, error) {
	args := r.Called(documentNumber)
	return "", args.Error(0)
}

func TestCreateAndSaveClient(t *testing.T) {
	repository := new(repositoryMock)
	service := Service{Repository: repository}

	newClient := contract.NewClient{
		DocumentNumber: "12365497899",
	}

	repository.On("Save", mock.Anything).Return(nil)
	repository.On("GetClientByDocumentNumber", mock.Anything).Return(nil)

	service.CreateClient(newClient)

	repository.AssertExpectations(t)
}
