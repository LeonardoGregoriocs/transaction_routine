package client

type Repository interface {
	Save(client *Client) error
	GetBy(id int) (*Client, error)
}
