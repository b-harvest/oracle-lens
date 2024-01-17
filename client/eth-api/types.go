package api

import "github.com/ethereum/go-ethereum/ethclient"

type Client struct {
	host string
	client *ethclient.Client
}
