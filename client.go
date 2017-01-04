package mcmarket

import "github.com/ngmachado/mkm"

type Client struct {
	Account *Account
	Message *Message
}

func NewClient(keys *mkm.Keys, endpoint mkm.Endpoint) *Client {
	cl := mkm.NewClient(keys, endpoint, mkm.V2, mkm.JSON)
	client := &Client{
		Account: &Account{cl},
		Message: &Message{cl},
	}
	return client
}
