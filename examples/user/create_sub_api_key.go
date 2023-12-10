package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	CreateSubApiKey()
}

func CreateSubApiKey() {
	client := bybit.NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"subuid": "100366706", "readOnly": 0,
		"permissions": []map[string]interface{}{
			{
				"Wallet": []string{"AccountTransfer", "SubMemberTransferList"},
			},
		},
	}
	userResult, err := client.NewUserService(params).CreateSubApiKey(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(userResult))
}
