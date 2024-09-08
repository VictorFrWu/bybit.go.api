package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	GetCoinInfo()
}

func GetCoinInfo() {
	client := bybit.NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"coin": "USDT"}
	assetResult, err := client.NewUtaBybitServiceWithParams(params).GetCoinInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(assetResult))
}
