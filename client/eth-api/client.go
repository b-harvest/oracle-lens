package api

import (
	"context"

	"bharvest.io/oracle-lens/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func New(host string) *Client {
	return &Client{
		host: host,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
		return err
    }

	c.client = client

	log.Info("ETH API connected")

	return nil
}

func (c *Client) Terminate(_ context.Context) {
	c.client.Close()
	log.Info("ETH API terminated")
}
