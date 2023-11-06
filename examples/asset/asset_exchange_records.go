package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	AssetInfo()
}

func AssetInfo() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"accountType": "spot"}
	assetResult, err := client.NewAssetService(params).GetAssetInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(assetResult))
}
