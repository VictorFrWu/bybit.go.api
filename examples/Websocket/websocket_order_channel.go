package main

import (
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	_ = ws.Connect([]string{"order"})
	select {}
}
