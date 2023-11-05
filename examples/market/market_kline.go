package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	MarketKline()
}

func MarketKline() {

	client := bybit.NewBybitHttpClient("", "")

	// NewServerTimeService
	marketKline, err := client.NewMarketKlineService("premium-index-price-kline", "linear", "BTCUSDT", "1").Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(marketKline))
}
