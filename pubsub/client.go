package pubsub

import "log"

type Client struct {
	Topic string
}

func NewPubSub(topic string) *Client {
	return &Client{
		Topic: topic,
	}
}

func (p *Client) Publish(message string) error {
	log.Printf("publishing message: %s", message)
	return nil
}
