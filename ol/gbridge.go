package ol

import (
	"context"
	"errors"
	"sync"

	types "bharvest.io/oracle-lens/types/gbridge"
	"google.golang.org/grpc"
)

type Gravity struct {
	orchestrator string
	conn *grpc.ClientConn
	batch chan int
	valset chan int
	nonce chan uint64
	eth_nonce chan uint64
}

func NewGBridge(orchestrator string, grpc_url string) *Gravity {
	conn, err := grpc.Dial(grpc_url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &Gravity{
		orchestrator,
		conn,
		make(chan int),
		make(chan int),
		make(chan uint64),
		make(chan uint64),
	}
}

func (gravity *Gravity) Check(ctx context.Context) (int, int, uint64, uint64, error) {
	defer gravity.conn.Close()
	defer close(gravity.batch)
	defer close(gravity.valset)
	defer close(gravity.nonce)
	defer close(gravity.eth_nonce)

	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)
	go gravity.queryLastPendingBatch(ctx, &wgCheck)
 	go gravity.queryLastPendingValset(ctx, &wgCheck)
 	go gravity.queryLastEventNonce(ctx, &wgCheck)
	go gravity.queryEthLastEventNonce(ctx, &wgCheck)

	var batch, valset int
	var nonce, eth_nonce uint64
	select {
	case result := <-gravity.batch:
		batch = result
		valset = <- gravity.valset
		nonce = <- gravity.nonce
		eth_nonce = <- gravity.eth_nonce
		return batch, valset, nonce, eth_nonce, nil
	case <- ctx.Done():
		return 0, 0, 0, 0, errors.New("Time out")
	}

}

func (gravity *Gravity) queryLastPendingBatch(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(gravity.conn)
	resp, err := client.LastPendingBatchRequestByAddr(
		ctx,
		&types.QueryLastPendingBatchRequestByAddrRequest{Address: gravity.orchestrator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case gravity.batch <- resp.Size():
		return
	case <- ctx.Done():
		Error(err)
		return
	}
}

func (gravity *Gravity) queryLastPendingValset(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(gravity.conn)
	resp, err := client.LastPendingValsetRequestByAddr(
		ctx,
		&types.QueryLastPendingValsetRequestByAddrRequest{Address: gravity.orchestrator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case gravity.valset <- resp.Size():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (gravity *Gravity) queryLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(gravity.conn)
	resp, err := client.LastEventNonceByAddr(
		ctx,
		&types.QueryLastEventNonceByAddrRequest{Address: gravity.orchestrator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case gravity.nonce <- resp.GetEventNonce():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (gravity *Gravity) queryEthLastEventNonce(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := types.NewQueryClient(gravity.conn)
	resp, err := client.GetLastObservedEthNonce(
		ctx,
		&types.QueryLastObservedEthNonceRequest{},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case gravity.eth_nonce <- resp.GetNonce():
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}
