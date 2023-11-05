package main

import (
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	// Create a public WebSocket client
	ws := bybit.NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	// Connect and subscribe to the desired topic
	_ = ws.Connect([]string{"orderbook.1.BTCUSDT"})
	select {}
}
