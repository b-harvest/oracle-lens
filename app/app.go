package app

import (
	"context"
	"sync"
	"time"

	"bharvest.io/oracle-lens/log"
)

func Run(ctx context.Context, cfg *Config) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Check Injective
	go func() {
		defer wg.Done()
		if !cfg.Injective.Enable {
			return
		}

		err := cfg.CheckInjective(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	// Check Umee Oracle
	go func() {
		defer wg.Done()
		if !cfg.UmeeOracle.Enable {
			return
		}

		err := cfg.CheckUmeeOracle(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()

	return
}
