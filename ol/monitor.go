package ol

import (
	"context"
	"fmt"
	"sync"
)

func CheckUmee(ctx context.Context, wg *sync.WaitGroup) {
	chain := NewUmee("umee1m2wqppm7pujtusd6u6zt0k92rd50e80uk7esz8", "umee-grpc.polkachu.com:13690")
	r1, r2, r3,r4 := chain.Check(ctx)
	select {
	case <- ctx.Done():
		fmt.Println("========================= Umee =========================")
		diff := r4 - r3
		if diff == 0 && r1 == 0 && r2 == 0 {
			fmt.Println("Status:🟢")
		} else  if diff <= 5 {
			fmt.Println("Status: 🟠")
		} else {
			fmt.Println("Status: 🛑")
		}
		fmt.Println("pending batch:", r1)
		fmt.Println("pending valset:", r2)
		fmt.Println("nonce:", r3)
		fmt.Println("eth nonce:", r4)
		fmt.Println("========================================================")
		wg.Done()
		return
	}
}

func CheckUmeeOracle(ctx context.Context, wg *sync.WaitGroup) {
	chain := NewUmeeOracle("umeevaloper1g7ddsx97qj4tfvm7u22xe5h6y3efdup6spsxnw", "umee-grpc.polkachu.com:13690")
	status, uptime, min_uptime, accept_list, voted_list := chain.Check(ctx)
	select {
	case <- ctx.Done():
		fmt.Println("========================= Umee Oracle =========================")
		len_acc := len(accept_list)
		len_vot := len(voted_list)
		if len_acc != len_vot {
		}
		if len_acc == len_vot && uptime > min_uptime*10 {
			fmt.Println("Status:🟢")
		} else {
			fmt.Println("Status: 🛑")
		}
		fmt.Println("Window/Current\tUptime\t  Missed")
		fmt.Println(status)
		fmt.Println("Accept List:", len_acc, accept_list)
		fmt.Println("Voted List:", len_vot, voted_list)
		fmt.Println("========================================================")
		wg.Done()
		return
	}
}

func CheckInjective(ctx context.Context, wg *sync.WaitGroup) {
 	chain := NewInjective("inj18uvllr7jvqgre49vtx7w4gs73v5d7dunen2mz5", "injective-grpc.polkachu.com:14390")
 	r1, r2, r3,r4 := chain.Check(ctx)

 	select {
 	case <- ctx.Done():
 		fmt.Println("========================= Injective =========================")
 		diff := r4 - r3
 		if diff == 0 && r1 == 0 && r2 == 0 {
 			fmt.Println("Status:🟢")
 		} else  if diff <= 5 {
 			fmt.Println("Status: 🟠")
 		} else {
 			fmt.Println("Status: 🛑")
 		}
 		fmt.Println("pending batch:", r1)
 		fmt.Println("pending valset:", r2)
 		fmt.Println("nonce:", r3)
 		fmt.Println("eth nonce:", r4)
 		fmt.Println("========================================================")

 		wg.Done()
 		return
 	}
}

func CheckGbridge(ctx context.Context, wg *sync.WaitGroup) {
 	chain := NewGBridge("gravity1uhuwpa0dssdrx95a8hajc4tcaknxl4d8p8vrr4", "gravity-grpc.polkachu.com:14290")
 	r1, r2, r3,r4 := chain.Check(ctx)

 	select {
 	case <- ctx.Done():
 		fmt.Println("========================= Gravity Bridge =========================")
 		diff := r4 - r3
 		if diff == 0 && r1 == 0 && r2 == 0 {
 			fmt.Println("Status:🟢")
 		} else  if diff <= 5 {
 			fmt.Println("Status: 🟠")
 		} else {
 			fmt.Println("Status: 🛑")
 		}
 		fmt.Println("pending batch:", r1)
 		fmt.Println("pending valset:", r2)
 		fmt.Println("nonce:", r3)
 		fmt.Println("eth nonce:", r4)
 		fmt.Println("========================================================")

 		wg.Done()
 		return
 	}
}
