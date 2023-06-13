package ol

import (
	"context"
	"sync"

	contract "bharvest.io/oracle-lens/ol/gravity-contract"
	types "bharvest.io/oracle-lens/types/gbridge"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type Umee struct {
	orchestrator string
	conn *grpc.ClientConn
	batch chan int
	valset chan int
	nonce chan uint64
	eth_nonce chan uint64
}

func NewUmee(orchestrator string, grpc_url string) *Umee {
	conn, err := grpc.Dial(grpc_url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &Umee{
		orchestrator,
		conn,
		make(chan int),
		make(chan int),
		make(chan uint64),
		make(chan uint64),
	}
}

func (umee *Umee) Check(ctx context.Context) (int, int, uint64, uint64) {
	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)

	go umee.queryLastPendingBatch(ctx, &wgCheck)
	go umee.queryLastPendingValset(ctx, &wgCheck)
	go umee.queryLastEventNonce(ctx, &wgCheck)
	go umee.queryEthLastEventNonce(ctx, &wgCheck)

	var r1, r2 int
	var r3, r4 uint64
	for {
		select {
		case result := <-umee.batch:
			r1 = result
		case result := <-umee.valset:
			r2 = result
		case result := <-umee.nonce:
			r3 = result
		case result := <- umee.eth_nonce:
			r4 = result
		case <- ctx.Done():
			umee.conn.Close()
			close(umee.batch)
			close(umee.valset)
			close(umee.nonce)
			close(umee.eth_nonce)
			return r1, r2, r3, r4
		}
	}
}

func (umee *Umee) queryLastPendingBatch(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(umee.conn)
	resp, err := client.LastPendingBatchRequestByAddr(
		ctx,
		&types.QueryLastPendingBatchRequestByAddrRequest{Address: umee.orchestrator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case umee.batch <- resp.Size():
		return
	case <- ctx.Done():
		Error(err)
		return
	}
}

func (umee *Umee) queryLastPendingValset(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(umee.conn)
	resp, err := client.LastPendingValsetRequestByAddr(
		ctx,
		&types.QueryLastPendingValsetRequestByAddrRequest{Address: umee.orchestrator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case umee.valset <- resp.Size():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (umee *Umee) queryLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(umee.conn)
	resp, err := client.LastEventNonceByAddr(
		ctx,
		&types.QueryLastEventNonceByAddrRequest{Address: umee.orchestrator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case umee.nonce <- resp.GetEventNonce():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (umee *Umee) queryEthLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
		Error(err)
		return
    }

	contract_address := common.HexToAddress("0xb564ac229E9D6040a9f1298B7211b9e79eE05a2c")
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
	case umee.eth_nonce <- last_event_nonce.Uint64():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}
