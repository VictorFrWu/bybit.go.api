package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	ServerTime()
}

func ServerTime() {

	client := bybit.NewBybitHttpClient("", "")

	// NewServerTimeService
	serverTime, err := client.NewMarketInfoServiceNoParams().GetServerTime(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverTime))
}
