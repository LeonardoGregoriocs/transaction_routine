package client

// import (
// 	"testing"

// 	"github.com/leonardogregoriocs/internal/contract"
// 	"github.com/stretchr/testify/mock"
// )

// type repositoryMock struct {
// 	mock.Mock
// }

// func (r *repositoryMock) Save(client *Client) error {
// 	args := r.Called(client)
// 	return args.Error(0)
// }

// var (
// 	newClient = contract.NewClient{
// 		DocumentNumber: "123456789",
// 	}
// 	repository = new(repositoryMock)
// 	service    = Service{Repository: repository}
// )

// func TestCreateAndSaveClient(t *testing.T) {
// 	newClient := contract.NewClient{
// 		DocumentNumber: "123456789",
// 	}

// 	repository.On("Save", mock.Anything).Return(nil)

// 	service.CreateClient(newClient)

// 	repository.AssertExpectations(t)
// }
