package ol

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	oracleTypes "bharvest.io/oracle-lens/types/oracle"
	"google.golang.org/grpc"
)

type UmeeOracle struct {
	validator string
	conn *grpc.ClientConn
	accept_list chan oracleTypes.DenomList
	aggregate_vote chan oracleTypes.ExchangeRateTuples
	window chan uint64
	min_uptime chan float64
	window_progress chan uint64
	miss_count chan uint64
}

func NewUmeeOracle(validator string, grpc_url string) *UmeeOracle {
	conn, err := grpc.Dial(grpc_url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &UmeeOracle{
		validator,
		conn,
		make(chan oracleTypes.DenomList),
		make(chan oracleTypes.ExchangeRateTuples),
		make(chan uint64),
		make(chan float64),
		make(chan uint64),
		make(chan uint64),
	}
}

func (umee *UmeeOracle) Check(ctx context.Context) (string, float64, float64, []string, []string) {
	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)

	go umee.queryOracleParams(ctx, &wgCheck)
	go umee.queryWindow(ctx, &wgCheck)
	go umee.queryVote(ctx, &wgCheck)
	go umee.queryMissCnt(ctx, &wgCheck)

	var r1 oracleTypes.DenomList
	var r2 oracleTypes.ExchangeRateTuples
	var r4 float64
	var r3, r5, r6 uint64
	for {
		select {
		case result := <-umee.accept_list:
			r1 = result
		case result := <-umee.aggregate_vote:
			r2 = result
		case result := <-umee.window:
			r3 = result
		case result := <- umee.min_uptime:
			r4 = result * 100
		case result := <- umee.window_progress:
			r5 = result
		case result := <- umee.miss_count:
			r6 = result
		case <- ctx.Done():
			umee.conn.Close()
			close(umee.accept_list)
			close(umee.aggregate_vote)
			close(umee.window)
			close(umee.min_uptime)
			close(umee.window_progress)
			close(umee.miss_count)

			uptime := (float64(r5-r6) / float64(r5)) * 100
			status := fmt.Sprintf("%d/%d\t%.2f%%\t  %d", r3, r5, uptime, r6)

			l := len(r1)
			accept_list := make([]string, l, l)
			for i, item := range r1 {
				accept_list[i] = strings.ToUpper(item.SymbolDenom)
			}
			sort.Strings(accept_list)

			l = len(r2)
			voted_list := make([]string, l, l)
			for i, item := range r2 {
				voted_list[i] = item.Denom
			}
			sort.Strings(voted_list)

			return status, uptime, r4, accept_list, voted_list
		}
	}
}

func (umee *UmeeOracle) queryOracleParams(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := oracleTypes.NewQueryClient(umee.conn)
	resp, err := client.Params(
		ctx,
		&oracleTypes.QueryParams{},
	)
	if err != nil {
		Error(err)
		return
	}

	
	min_uptime, err := resp.Params.MinValidPerWindow.Float64()
	if err != nil {
		Error(err)
		return
	}
	select {
	case umee.accept_list <- resp.Params.AcceptList:
		umee.window <-resp.Params.SlashWindow 
		umee.min_uptime <- min_uptime
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (umee *UmeeOracle) queryWindow(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := oracleTypes.NewQueryClient(umee.conn)
	resp, err := client.SlashWindow(
		ctx,
		&oracleTypes.QuerySlashWindow{},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case umee.window_progress <- resp.WindowProgress:
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (umee *UmeeOracle) queryMissCnt(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := oracleTypes.NewQueryClient(umee.conn)
	resp, err := client.MissCounter(
		ctx,
		&oracleTypes.QueryMissCounter{ValidatorAddr: umee.validator},
	)
	if err != nil {
		Error(err)
		return
	}

	select {
	case umee.miss_count <- resp.MissCounter:
		return
	case <- ctx.Done():
		Error("Time out")
		return
	}
}

func (umee *UmeeOracle) queryVote(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := oracleTypes.NewQueryClient(umee.conn)

	for i:=0; i<5; i++ {
		resp, err := client.AggregateVote(
			ctx,
			&oracleTypes.QueryAggregateVote{ValidatorAddr: umee.validator},
		)

		if resp != nil {
			umee.aggregate_vote <- resp.AggregateVote.ExchangeRateTuples
			return
		}
		if err != nil {
			<- time.After(time.Second*2)
			Info(err)
		}
	}

	select {
	case <- ctx.Done():
		Error("Time out")
		return
	}
}
