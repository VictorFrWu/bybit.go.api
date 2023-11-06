package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	PlaceOrder()
}

func PlaceOrder() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "10").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
