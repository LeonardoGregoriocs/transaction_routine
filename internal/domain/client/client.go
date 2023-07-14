package client

import (
	"errors"
	"regexp"
	"time"
)

type Client struct {
	ID               int
	DocumentNumber   string
	DateCreateUpdate time.Time
}

func NewClient(documentNumber string) (*Client, error) {
	documentNumberObj, err := regexp.Compile(`[-.,-]`)
	documentNumber = documentNumberObj.ReplaceAllString(documentNumber, "")

	if err != nil {
		return nil, errors.New("Error formated Document Number")
	}

	if len(documentNumber) != 11 {
		return nil, errors.New("Documento Number invalid")
	}

	return &Client{
		DocumentNumber:   documentNumber,
		DateCreateUpdate: time.Now(),
	}, nil
}
