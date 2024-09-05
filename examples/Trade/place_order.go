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
	client := bybit.NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", bybit.WithBaseURL(bybit.TESTNET), bybit.WithDebug(true))
	orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "15").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
