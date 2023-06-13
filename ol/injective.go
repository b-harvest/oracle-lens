package ol

import (
	"context"
	"sync"

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

func (injective *Injective) Check(ctx context.Context) (int, int, uint64, uint64) {
	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)
	go injective.queryLastPendingBatch(ctx, &wgCheck)
 	go injective.queryLastPendingValset(ctx, &wgCheck)
 	go injective.queryLastEventNonce(ctx, &wgCheck)
	go injective.queryEthLastEventNonce(ctx, &wgCheck)

	var r1, r2 int
	var r3, r4 uint64
	for {
		select {
		case result := <-injective.batch:
			r1 = result
		case result := <-injective.valset:
			r2 = result
		case result := <-injective.nonce:
			r3 = result
		case result := <- injective.eth_nonce:
			r4 = result
		case <- ctx.Done():
			injective.conn.Close()
			close(injective.batch)
			close(injective.valset)
			close(injective.nonce)
			close(injective.eth_nonce)
			return r1, r2, r3, r4
		}
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
		Error(err)
		return
	}

	select {
	case injective.batch <- resp.Size():
		return
	case <- ctx.Done():
		Error(err)
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
		Error(err)
		return
	}

	select {
	case injective.valset <- resp.Size():
		return
	case <- ctx.Done():
		Error("Time out")
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
		Error(err)
		return
	}

	select {
	case injective.nonce <- resp.GetLastClaimEvent().GetEthereumEventNonce():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (injective *Injective) queryEthLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
		Error(err)
		return
    }

	contract_address := common.HexToAddress("0xF955C57f9EA9Dc8781965FEaE0b6A2acE2BAD6f3")
	instance, err := contract.NewContract(contract_address, client)
	if err != nil {
		Error(err)
		return
	}

	last_event_nonce, err := instance.StateLastEventNonce(&bind.CallOpts{})
	if err != nil {
		Error(err)
		return
	}

	select {
	case injective.eth_nonce <- last_event_nonce.Uint64():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}
