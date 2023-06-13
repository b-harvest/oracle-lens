package ol

import (
	"context"
	"sync"

	"bharvest.io/oracle-lens/types/gbridge"
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

func (gravity *Gravity) Check(ctx context.Context) (int, int, uint64, uint64) {
	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)
	go gravity.queryLastPendingBatch(ctx, &wgCheck)
 	go gravity.queryLastPendingValset(ctx, &wgCheck)
 	go gravity.queryLastEventNonce(ctx, &wgCheck)
	go gravity.queryEthLastEventNonce(ctx, &wgCheck)

	var r1, r2 int
	var r3, r4 uint64
	for {
		select {
		case result := <-gravity.batch:
			r1 = result
		case result := <-gravity.valset:
			r2 = result
		case result := <-gravity.nonce:
			r3 = result
		case result := <- gravity.eth_nonce:
			r4 = result
		case <- ctx.Done():
			gravity.conn.Close()
			close(gravity.batch)
			close(gravity.valset)
			close(gravity.nonce)
			close(gravity.eth_nonce)
			return r1, r2, r3, r4
		}
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
