package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	SwitchPositionMode()
}

func SwitchPositionMode() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "coin": "USDT", "mode": 3}
	orderResult, err := client.NewPositionService(params).SwitchPositionMode(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
