package api

import (
	"context"

	contract "bharvest.io/oracle-lens/client/eth-api/gravity-contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) GetETHLastEventNonce(ctx context.Context, contract_hex_address string) (uint64, error) {
	contract_address := common.HexToAddress(contract_hex_address)
	instance, err := contract.NewContract(contract_address, c.client)
	if err != nil {
		return 0, err
	}

	lastEventNonce, err := instance.StateLastEventNonce(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return lastEventNonce.Uint64(), nil
}
