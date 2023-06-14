package ol

import (
	"context"
	"fmt"
	"sync"
)

func CheckUmee(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	chain := NewUmee("umee1m2wqppm7pujtusd6u6zt0k92rd50e80uk7esz8", "umee-grpc.polkachu.com:13690")
	batch, valset, nonce, eth_nonce, err := chain.Check(ctx)
	if err != nil {
		Error(err)
		return
	}

	fmt.Println("========================= Umee =========================")
	diff := eth_nonce - nonce
	if diff == 0 && batch == 0 && valset == 0 {
		fmt.Println("Status:🟢")
	} else  if diff <= 5 {
		fmt.Println("Status: 🟠")
	} else {
		fmt.Println("Status: 🛑")
	}
	fmt.Println("pending batch:", batch)
	fmt.Println("pending valset:", valset)
	fmt.Println("nonce:", nonce)
	fmt.Println("eth nonce:", eth_nonce)
	fmt.Println("========================================================")

	return
}

func CheckUmeeOracle(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	chain := NewUmeeOracle("umeevaloper1g7ddsx97qj4tfvm7u22xe5h6y3efdup6spsxnw", "umee-grpc.polkachu.com:13690")
	status, uptime, min_uptime, accept_list, voted_list, err := chain.Check(ctx)
	if err != nil {
		Error(err)
		return
	}

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

	return
}

func CheckInjective(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

 	chain := NewInjective("inj18uvllr7jvqgre49vtx7w4gs73v5d7dunen2mz5", "injective-grpc.polkachu.com:14390")
 	batch, valset, nonce,eth_nonce, err := chain.Check(ctx)
	if err != nil {
		Error(err)
		return
	}

	fmt.Println("========================= Injective =========================")
	diff := eth_nonce - nonce
	if diff == 0 && batch == 0 && valset == 0 {
		fmt.Println("Status:🟢")
	} else  if diff <= 5 {
		fmt.Println("Status: 🟠")
	} else {
		fmt.Println("Status: 🛑")
	}
	fmt.Println("pending batch:", batch)
	fmt.Println("pending valset:", valset)
	fmt.Println("nonce:", nonce)
	fmt.Println("eth nonce:", eth_nonce)
	fmt.Println("========================================================")

	return
}

func CheckGbridge(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

 	chain := NewGBridge("gravity1uhuwpa0dssdrx95a8hajc4tcaknxl4d8p8vrr4", "gravity-grpc.polkachu.com:14290")
 	batch, valset, nonce,eth_nonce, err := chain.Check(ctx)
	if err != nil {
		Error(err)
		return
	}

 	fmt.Println("========================= Gravity Bridge =========================")
	diff := eth_nonce - nonce
	if diff == 0 && batch == 0 && valset == 0 {
		fmt.Println("Status:🟢")
	} else  if diff <= 5 {
		fmt.Println("Status: 🟠")
	} else {
		fmt.Println("Status: 🛑")
	}
	fmt.Println("pending batch:", batch)
	fmt.Println("pending valset:", valset)
	fmt.Println("nonce:", nonce)
	fmt.Println("eth nonce:", eth_nonce)
	fmt.Println("========================================================")

	return
}
