package mongo

import "log"

type Client struct {
	Url string
}

func NewMongo(url string) *Client {
	return &Client{
		Url: url,
	}
}

func (m *Client) Ping() error {
	log.Println("pinging mongo")
	return nil
}
