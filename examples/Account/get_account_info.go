package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	GetAccountInfo()
}

func GetAccountInfo() {
	client := bybit.NewBybitHttpClient("X6wmWloIPvaLXAKqv2", "rY1CWGYLHy0AUjdNZqqspvd3Krhp79fHp1sP", bybit.WithBaseURL(bybit.TESTNET))
	accountResult, err := client.NewUtaBybitServiceNoParams().GetAccountInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
