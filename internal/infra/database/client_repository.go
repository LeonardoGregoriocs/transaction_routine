package database

import (
	"github.com/leonardogregoriocs/internal/domain/client"
	"gorm.io/gorm"
)

type ClientRepository struct {
	Db *gorm.DB
}

func (c *ClientRepository) Save(client *client.Client) error {
	result := c.Db.Create(client)
	return result.Error
}

func (c *ClientRepository) GetBy(id int) (*client.Client, error) {
	var client client.Client
	result := c.Db.First(&client, id)
	return &client, result.Error
}

func (c *ClientRepository) GetClientByDocumentNumber(documentNumber string) (string, error) {
	var client client.Client
	result := c.Db.Find(&client, "document_number = ?", documentNumber)
	return client.DocumentNumber, result.Error
}
