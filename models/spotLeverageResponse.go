package models

import "time"

// SpotLeverageTokenInfo holds the information for leveraged tokens.
type SpotLeverageTokenInfo struct {
	List []struct {
		LtCoin           string `json:"ltCoin"`
		LtName           string `json:"ltName"`
		MaxPurchase      string `json:"maxPurchase"`
		MinPurchase      string `json:"minPurchase"`
		MaxPurchaseDaily string `json:"maxPurchaseDaily"`
		MaxRedeem        string `json:"maxRedeem"`
		MinRedeem        string `json:"minRedeem"`
		MaxRedeemDaily   string `json:"maxRedeemDaily"`
		PurchaseFeeRate  string `json:"purchaseFeeRate"`
		RedeemFeeRate    string `json:"redeemFeeRate"`
		LtStatus         string `json:"ltStatus"`
		FundFee          string `json:"fundFee"`
		FundFeeTime      string `json:"fundFeeTime"` // Should parse to time.Time if needed
		ManageFeeRate    string `json:"manageFeeRate"`
		ManageFeeTime    string `json:"manageFeeTime"` // Should parse to time.Time if needed
		Value            string `json:"value"`
		NetValue         string `json:"netValue"`
		Total            string `json:"total"`
	}
}

// SpotLeverageTokenMarket holds the market information for a leveraged token.
type SpotLeverageTokenMarket struct {
	LtCoin      string    `json:"ltCoin"`
	Nav         string    `json:"nav"`
	NavTime     time.Time `json:"navTime"` // Assuming time is in milliseconds, need to convert
	Circulation string    `json:"circulation"`
	Basket      string    `json:"basket"`
	Leverage    string    `json:"leverage"`
}

// SpotLeverageOrdersRecords holds the orders records for leveraged tokens.
type SpotLeverageOrdersRecords struct {
	List []struct {
		LtCoin        string `json:"ltCoin"`
		OrderId       string `json:"orderId"`
		LtOrderType   int    `json:"ltOrderType"` // Use int for integer type
		OrderTime     int64  `json:"orderTime"`   // Assuming this is a timestamp, int64 is used
		UpdateTime    int64  `json:"updateTime"`  // Assuming this is a timestamp, int64 is used
		LtOrderStatus string `json:"ltOrderStatus"`
		Fee           string `json:"fee"`
		Amount        string `json:"amount"`
		Value         string `json:"value"`
		ValueCoin     string `json:"valueCoin"`
		SerialNo      string `json:"serialNo"`
	}
}

// LeverageTokenPurchaseResult represents the result of a leverage token purchase.
type LeverageTokenPurchaseResult struct {
	LtCoin        string `json:"ltCoin"`
	LtOrderStatus string `json:"ltOrderStatus"`
	ExecQty       string `json:"execQty"`
	ExecAmt       string `json:"execAmt"`
	Amount        string `json:"amount"`
	PurchaseId    string `json:"purchaseId"`
	SerialNo      string `json:"serialNo"`
	ValueCoin     string `json:"valueCoin"`
}

// LeverageTokenRedeem represents the redeem information for a leveraged token.
type LeverageTokenRedeem struct {
	LtCoin        string `json:"ltCoin"`
	LtOrderStatus string `json:"ltOrderStatus"`
	Quantity      string `json:"quantity"`
	ExecQty       string `json:"execQty"`
	ExecAmt       string `json:"execAmt"`
	RedeemId      string `json:"redeemId"`
	SerialNo      string `json:"serialNo"`
	ValueCoin     string `json:"valueCoin"`
}
