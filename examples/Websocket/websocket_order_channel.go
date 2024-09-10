package main

import (
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	}, bybit.WithPingInterval(2))
	_, _ = ws.Connect().SendSubscription([]string{"order", "position", "wallet"})
	select {}
}
