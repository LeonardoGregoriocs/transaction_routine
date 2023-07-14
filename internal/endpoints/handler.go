package endpoints

import "github.com/leonardogregoriocs/internal/domain/client"

type Handler struct {
	ClientService client.Service
}
