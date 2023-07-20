package client

import (
	"errors"
	"time"
)

type Client struct {
	ID                   int
	DocumentNumber       string
	DateCreateUpdate     time.Time
	AvailableCreditLimit float64
}

func NewClient(documentNumber string) (*Client, error) {
	if len(documentNumber) != 11 {
		return nil, errors.New("Documento Number invalid")
	}

	return &Client{
		DocumentNumber:   documentNumber,
		DateCreateUpdate: time.Now(),
	}, nil
}
