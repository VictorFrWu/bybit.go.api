package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	InstrumentInfo()
}

func InstrumentInfo() {

	client := bybit.NewBybitHttpClient("", "")

	// NewServerTimeService
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "status": "Trading"}
	marketKline, err := client.NewMarketInfoService(params).GetInstrumentInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(marketKline))
}
