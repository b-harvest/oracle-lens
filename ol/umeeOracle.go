package ol

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"bharvest.io/oracle-lens/log"
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

func (umee *UmeeOracle) Check(ctx context.Context) (string, float64, float64, []string, []string, error) {
	defer umee.conn.Close()
	defer close(umee.accept_list)
	defer close(umee.aggregate_vote)
	defer close(umee.window)
	defer close(umee.min_uptime)
	defer close(umee.window_progress)
	defer close(umee.miss_count)

	wgCheck := sync.WaitGroup{}
	wgCheck.Add(4)

	go umee.queryOracleParams(ctx, &wgCheck)
	go umee.queryWindow(ctx, &wgCheck)
	go umee.queryVote(ctx, &wgCheck)
	go umee.queryMissCnt(ctx, &wgCheck)

	select {
	case result := <-umee.accept_list:
		accept_list_temp := result
		voted_list_temp :=  <- umee.aggregate_vote
		window := <- umee.window
		min_uptime := <- umee.min_uptime * 100
		window_progress := <- umee.window_progress
		miss_count := <- umee.miss_count

		uptime := (float64(window_progress-miss_count) / float64(window_progress)) * 100
		status := fmt.Sprintf("%d/%d\t%.2f%%\t  %d", window, window_progress, uptime, miss_count)

		l := len(accept_list_temp)
		accept_list := make([]string, l, l)
		for i, item := range accept_list_temp {
			accept_list[i] = strings.ToUpper(item.SymbolDenom)
		}
		sort.Strings(accept_list)

		l = len(voted_list_temp)
		voted_list := make([]string, l, l)
		for i, item := range voted_list_temp {
			voted_list[i] = item.Denom
		}
		sort.Strings(voted_list)

		return status, uptime, min_uptime, accept_list, voted_list, nil
	case <- ctx.Done():
		return "", 0, 0, nil, nil, errors.New("Time out")
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
		log.Error(err, "umeeOracle.go", "queryOracleParams()")
		return
	}

	
	min_uptime, err := resp.Params.MinValidPerWindow.Float64()
	if err != nil {
		log.Error(err, "umeeOracle.go", "queryOracleParams()")
		return
	}
	select {
	case umee.accept_list <- resp.Params.AcceptList:
		umee.window <-resp.Params.SlashWindow 
		umee.min_uptime <- min_uptime
		return
	case <- ctx.Done():
		log.Error("Time out", "umeeOracle.go", "queryOracleParams()")
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
		log.Error(err, "umeeOracle.go", "queryWindow()")
		return
	}

	select {
	case umee.window_progress <- resp.WindowProgress:
		return
	case <- ctx.Done():
		log.Error("Time out", "umeeOracle.go", "queryWindow()")
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
		log.Error(err, "umeeOracle.go", "queryMissCnt()")
		return
	}

	select {
	case umee.miss_count <- resp.MissCounter:
		return
	case <- ctx.Done():
		log.Error("Time out", "umeeOracle.go", "queryMissCnt()")
		return
	}
}

func (umee *UmeeOracle) queryVote(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := oracleTypes.NewQueryClient(umee.conn)

	for i:=0; i<10; i++ {
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
			log.Info(err, "umeeOracle.go", "queryVote()")
		}
	}

	select {
	case <- ctx.Done():
		log.Error("Time out", "umeeOracle.go", "queryVote()")
		return
	}
}
