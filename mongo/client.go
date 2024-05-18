package mongo

import (
	"fx_101/config"
	"log"
)

type Client struct {
	Url string
}

func NewMongo(config *config.Config) *Client {
	return &Client{
		Url: config.MongoUrl,
	}
}

func (m *Client) Ping() error {
	log.Println("pinging mongo")
	return nil
}
