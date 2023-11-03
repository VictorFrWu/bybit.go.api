package main

import (
	bybit "bybit.go.api"
	"fmt"
)

func main() {
	ws := bybit.NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	// Connect and subscribe to the desired topic
	_ = ws.Connect([]string{"order"})
	select {}
}
