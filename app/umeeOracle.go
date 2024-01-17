package app

import (
	"context"
	"fmt"
	"sync"

	"bharvest.io/oracle-lens/client/grpc"
	"bharvest.io/oracle-lens/log"
	"bharvest.io/oracle-lens/server"
	"bharvest.io/oracle-lens/tg"
)

func (cfg *Config) CheckUmeeOracle(ctx context.Context) error {
	grpcQuery := grpc.NewUmee(cfg.UmeeOracle.GRPC)
	err := grpcQuery.Connect(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	defer func() {
		err := grpcQuery.Terminate(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(4)

	// For get Window, Minimum Uptime, Accept List
	var params *grpc.UmeeOracleParams
	go func() {
		defer wg.Done()

		params, err = grpcQuery.GetOracleParams(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	// For get current window progress
	var windowProgress uint64
	go func() {
		defer wg.Done()

		windowProgress, err = grpcQuery.GetWindowProgress(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	// For get miss count
	var missCnt uint64
	go func() {
		defer wg.Done()

		missCnt, err = grpcQuery.GetMissCount(ctx, cfg.UmeeOracle.Wallet.PrintValoper())
		if err != nil {
			log.Error(err)
			return
		}
	}()

	// For get vote list
	var voteList []string
	go func() {
		defer wg.Done()

		voteList, err = grpcQuery.GetVoteList(ctx, cfg.UmeeOracle.Wallet.PrintValoper())
		if err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()

	for _, voteItem := range voteList {
		params.AcceptList[voteItem] = true
	}
	uptime := calcUptime(windowProgress, missCnt)
	server.GlobalState.UmeeOracle.AcceptList = params.AcceptList
	server.GlobalState.UmeeOracle.Window = fmt.Sprintf("%d / %d", windowProgress, params.SlashWindow)
	server.GlobalState.UmeeOracle.Uptime = fmt.Sprintf("%f / %f", uptime, params.MinUptime)

	check := true

	for k, v := range server.GlobalState.UmeeOracle.AcceptList {
		if !v {
			check = false
			log.Info(fmt.Sprint(k, ":", v))
			break
		}
	}
	if uptime < cfg.UmeeOracle.MinUptime {
		check = false
	}

	server.GlobalState.UmeeOracle.Status = check
	m := "Umee Oracle status: "
	if check {
		m += "🟢"
		log.Info(m)
	} else {
		m += "🛑"
		log.Info(m)
		tg.SendMsg(m)
	}

	return nil
}

func calcUptime(window_progress uint64, miss_cnt uint64) float64 {
	return (float64(window_progress-miss_cnt) / float64(window_progress)) * 100
}
