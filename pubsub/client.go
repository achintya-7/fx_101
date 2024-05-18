package pubsub

import (
	"fx_101/config"
	"log"
)

type Client struct {
	Topic string
}

func NewPubSub(config *config.Config) *Client {
	return &Client{
		Topic: config.Topic,
	}
}

func (p *Client) Publish(message string) error {
	log.Printf("publishing message: %s", message)
	return nil
}
