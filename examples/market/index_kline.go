package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	IndexKline()
}

func IndexKline() {

	client := bybit.NewBybitHttpClient("", "")

	// NewServerTimeService
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "interval": "1", "Limit": 2}
	marketKline, err := client.NewMarketKLinesService("index-price-kline", params).GetMarketKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(marketKline))
}
