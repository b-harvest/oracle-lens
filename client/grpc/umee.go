package grpc

import (
	"context"
	"errors"
	"sort"
	"strings"
	"time"

	umeeOracleTypes "bharvest.io/oracle-lens/client/grpc/protobuf/umee-oracle"
	"bharvest.io/oracle-lens/log"
	"google.golang.org/grpc"
)

func NewUmee(host string) *Umee {
	return &Umee{
		host: host,
	}
}

func (client *Umee) Connect(ctx context.Context) error {
	conn, err := grpc.DialContext(
		ctx,
		client.host,
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	client.conn = conn
	client.oracleClient = umeeOracleTypes.NewQueryClient(conn)

	log.Info("Umee GRPC connected")

	return nil
}

func (client *Umee) Terminate(_ context.Context) error {
	err := client.conn.Close()
	log.Info("Umee GRPC connection terminated")

	return err
}

func (client *Umee) GetOracleParams(ctx context.Context) (*UmeeOracleParams, error) {
	resp, err := client.oracleClient.Params(
		ctx,
		&umeeOracleTypes.QueryParams{},
	)
	if err != nil {
		return nil, err
	}

	// Set-like data structure
	// For remove duplicated item
	acceptList := make(map[string]bool)
	for _, item := range resp.Params.AcceptList {
		symbol := strings.ToUpper(item.SymbolDenom)
		acceptList[symbol] = false
	}
	
	min_uptime, err := resp.Params.MinValidPerWindow.Float64()
	if err != nil {
		return nil, err
	}

	params := UmeeOracleParams{
		acceptList,
		resp.Params.SlashWindow,
		min_uptime*100,

	}

	return &params, nil
}

func (client *Umee) GetWindowProgress(ctx context.Context) (uint64, error) {
	resp, err := client.oracleClient.SlashWindow(
		ctx,
		&umeeOracleTypes.QuerySlashWindow{},
	)
	if err != nil {
		return 0, err
	}

	return resp.WindowProgress, nil
}

func (client *Umee) GetMissCount(ctx context.Context, validator_valoper string) (uint64, error) {
	resp, err := client.oracleClient.MissCounter(
		ctx,
		&umeeOracleTypes.QueryMissCounter{ValidatorAddr: validator_valoper},
	)
	if err != nil {
		return 0, err
	}

	return resp.MissCounter, nil
}

func (client *Umee) GetVoteList(ctx context.Context, validator_valoper string) ([]string, error) {
	for i:=0; i<10; i++ {
		resp, err := client.oracleClient.AggregateVote(
			ctx,
			&umeeOracleTypes.QueryAggregateVote{ValidatorAddr: validator_valoper},
		)
		if err != nil {
			<- time.After(2*time.Second)
			log.Info(err.Error())
			continue
		}

		l := len(resp.AggregateVote.ExchangeRateTuples)
		voteList := make([]string, l, l)
		for i, item := range resp.AggregateVote.ExchangeRateTuples {
			voteList[i] = strings.ToUpper(item.Denom)
		}
		sort.Strings(voteList)

		return voteList, nil
	}
	return nil, errors.New("Can't get a vote list on chain.")
}
