package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	PlaceTrade()
}

func PlaceTrade() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"coin": "USDT"}
	orderResult, err := client.NewLendingService(params).GetC2cLendingCoinInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
