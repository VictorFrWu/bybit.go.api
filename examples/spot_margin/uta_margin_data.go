package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	SpotMarginData()
}

func SpotMarginData() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"ltCoin": "BTC3L"}
	leverageTokenResult, err := client.NewSpotMarginDataService(params, true).GetSpotMarginData(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(leverageTokenResult))
}
