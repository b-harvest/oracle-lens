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
