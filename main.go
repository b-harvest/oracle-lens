package main

import (
	"bharvest.io/oracle-lens/metrics"
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"bharvest.io/oracle-lens/app"
	"bharvest.io/oracle-lens/client/grpc"
	"bharvest.io/oracle-lens/log"
	"bharvest.io/oracle-lens/server"
	"bharvest.io/oracle-lens/tg"
	"bharvest.io/oracle-lens/wallet"
	"github.com/BurntSushi/toml"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx := context.Background()

	cfgPath := flag.String("config", "", "Config file")
	flag.Parse()
	if *cfgPath == "" {
		panic("Error: Please input config file path with -config flag.")
	}

	f, err := os.ReadFile(*cfgPath)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	cfg := app.Config{}
	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if cfg.Injective.Enable {
		err = PrePrepareForInjective(ctx, &cfg)
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}

	if cfg.UmeeOracle.Enable {
		err = PrePrepareForUmee(ctx, &cfg)
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}

	tgTitle := fmt.Sprintf("🤖 Oracle Lens for %s 🤖", cfg.General.Network)
	tg.SetTg(cfg.Tg.Enable, tgTitle, cfg.Tg.Token, cfg.Tg.ChatID)

	metrics.Initialize(cfg.General.Network, cfg.Injective.Wallet.PrintAcc())

	go server.Run(cfg.General.ListenPort)
	for {
		app.Run(ctx, &cfg)
		time.Sleep(time.Duration(cfg.General.Period) * time.Minute)
	}
}

func PrePrepareForInjective(ctx context.Context, cfg *app.Config) error {
	grpcQuery := grpc.NewInjective(cfg.Injective.GRPC)
	err := grpcQuery.Connect(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	defer func() {
		err = grpcQuery.Terminate(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	tmpWallet, err := wallet.NewWallet(ctx, cfg.Injective.ValidatorAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	// Get Orchestrator Address from Validator Address
	OrchAcc, err := grpcQuery.GetDelegateKeyByValidator(ctx, tmpWallet.PrintValoper())
	if err != nil {
		log.Error(err)
		return err
	}
	cfg.Injective.Wallet, err = wallet.NewWallet(ctx, OrchAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func PrePrepareForUmee(ctx context.Context, cfg *app.Config) error {
	var err error
	cfg.UmeeOracle.Wallet, err = wallet.NewWallet(ctx, cfg.UmeeOracle.ValidatorAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
