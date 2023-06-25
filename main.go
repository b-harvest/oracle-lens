package main

import (
	"context"
	"runtime"
	"sync"
	"time"

	"bharvest.io/oracle-lens/ol"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for {
		switch {
		select <- time.Tick(10 * time.Second):
			ctx, cancel := context.WithTimeout(context.Background(), 21 * time.Second)
			wg := sync.WaitGroup{}
			wg.Add(4)

			go ol.CheckUmee(ctx, &wg)
			go ol.CheckUmeeOracle(ctx, &wg)
			go ol.CheckInjective(ctx, &wg)
			go ol.CheckGbridge(ctx, &wg)

			wg.Wait()
			cancel() 
		}
	}
}
