package main

import (
	bybit "bybit.go.api"
	"context"
	"fmt"
)

func main() {
	ServerTime()
}

func ServerTime() {

	client := bybit.NewClient("", "")

	// set to debug mode
	client.Debug = true

	// NewServerTimeService
	serverTime := client.NewServerTimeService().Do(context.Background())
	fmt.Println(bybit.PrettyPrint(serverTime))
}
