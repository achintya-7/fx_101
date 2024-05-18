package postgres

import "log"

type Client struct {
	Url string
}

func NewPostgres(url string) *Client {
	return &Client{
		Url: url,
	}
}

func (p *Client) Ping() error {
	log.Println("pinging postgres")
	return nil
}
