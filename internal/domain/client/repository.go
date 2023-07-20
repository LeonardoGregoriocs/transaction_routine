package client

type Repository interface {
	Save(client *Client) error
	GetBy(id int) (*Client, error)
	GetClientByDocumentNumber(documentNumber string) (string, error)
	UpdateCredit(amount float64, id int) error
}
