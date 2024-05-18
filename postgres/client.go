package postgres

import (
	"fx_101/config"
	"log"
)

type Client struct {
	Url string
}

func NewPostgres(config *config.Config) *Client {
	return &Client{
		Url: config.PostgresUrl,
	}
}

func (p *Client) Ping() error {
	log.Println("pinging postgres")
	return nil
}
