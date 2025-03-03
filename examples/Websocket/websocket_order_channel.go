package main

import (
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "X6wmWloIPvaLXAKqv2", "rY1CWGYLHy0AUjdNZqqspvd3Krhp79fHp1sP", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	}, bybit.WithPingInterval(2))
	_, _ = ws.Connect().SendSubscription([]string{"order", "position", "wallet"})
	select {}
}
