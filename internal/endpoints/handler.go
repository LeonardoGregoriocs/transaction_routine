package endpoints

import (
	"github.com/leonardogregoriocs/internal/domain/client"
	"github.com/leonardogregoriocs/internal/domain/transactions"
)

type Handler struct {
	ClientService       client.Service
	TransactionsService transactions.Service
}
