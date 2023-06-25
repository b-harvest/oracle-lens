package ol

import (
	"context"
	"errors"
	"sync"

	"bharvest.io/oracle-lens/log"
	contract "bharvest.io/oracle-lens/ol/gravity-contract"
	types "bharvest.io/oracle-lens/types/peggy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type Injective struct {
	orchestrator string
	conn *grpc.ClientConn
	batch chan int
	valset chan int
	nonce chan uint64
	eth_nonce chan uint64
}

func NewInjective(orchestrator string, grpc_url string) *Injective {
	conn, err := grpc.Dial(grpc_url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &Injective{
		orchestrator,
		conn,
		make(chan int),
		make(chan int),
		make(chan uint64),
		make(chan uint64),
	}
}

func (injective *Injective) Check(ctx context.Context) (int, int, uint64, uint64, error) {
	defer injective.conn.Close()
	defer close(injective.batch)
	defer close(injective.valset)
	defer close(injective.nonce)
	defer close(injective.eth_nonce)


	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)
	go injective.queryLastPendingBatch(ctx, &wgCheck)
 	go injective.queryLastPendingValset(ctx, &wgCheck)
 	go injective.queryLastEventNonce(ctx, &wgCheck)
	go injective.queryEthLastEventNonce(ctx, &wgCheck)

	var batch, valset int
	var nonce, eth_nonce uint64
	select {
	case result := <-injective.batch:
		batch = result
		valset = <- injective.valset
		nonce = <- injective.nonce
		eth_nonce = <- injective.eth_nonce
		return batch, valset, nonce, eth_nonce, nil
	case <- ctx.Done():
		return 0, 0, 0, 0, errors.New("Time out")
	}

}

func (injective *Injective) queryLastPendingBatch(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(injective.conn)
	resp, err := client.LastPendingBatchRequestByAddr(
		ctx,
		&types.QueryLastPendingBatchRequestByAddrRequest{Address: injective.orchestrator},
	)
	if err != nil {
		log.Error(err, "injective.go", "queryLastPendingBatch()")
		return
	}

	select {
	case injective.batch <- resp.Size():
		return
	case <- ctx.Done():
		log.Error(err, "injective.go", "queryLastPendingBatch()")
		return
	}
}

func (injective *Injective) queryLastPendingValset(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(injective.conn)
	resp, err := client.LastPendingValsetRequestByAddr(
		ctx,
		&types.QueryLastPendingValsetRequestByAddrRequest{Address: injective.orchestrator},
	)
	if err != nil {
		log.Error(err, "injective.go", "queryLastPendingValset()")
		return
	}

	select {
	case injective.valset <- resp.Size():
		return
	case <- ctx.Done():
		log.Error("Time out", "injective.go", "queryLastPendingValset()")
		return
	}
}

func (injective *Injective) queryLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(injective.conn)
	resp, err := client.LastEventByAddr(
		ctx,
		&types.QueryLastEventByAddrRequest{Address: injective.orchestrator},
	)
	if err != nil {
		log.Error(err, "injective.go", "queryLastEventNonce()")
		return
	}

	select {
	case injective.nonce <- resp.GetLastClaimEvent().GetEthereumEventNonce():
		return
	case <- ctx.Done():
		log.Error("Time out", "injective.go", "queryLastEventNonce()")
		return
	}
}

func (injective *Injective) queryEthLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
		log.Error(err, "injective.go", "queryEthLastEventNonce()")
		return
    }

	contract_address := common.HexToAddress("0xF955C57f9EA9Dc8781965FEaE0b6A2acE2BAD6f3")
	instance, err := contract.NewContract(contract_address, client)
	if err != nil {
		log.Error(err, "injective.go", "queryEthLastEventNonce()")
		return
	}

	last_event_nonce, err := instance.StateLastEventNonce(&bind.CallOpts{})
	if err != nil {
		log.Error(err, "injective.go", "queryEthLastEventNonce()")
		return
	}

	select {
	case injective.eth_nonce <- last_event_nonce.Uint64():
		return
	case <- ctx.Done():
		log.Error("Time out", "injective.go", "queryEthLastEventNonce()")
		return
	}
}
