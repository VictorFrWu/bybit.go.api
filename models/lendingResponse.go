package models

// LendingCoin represents information about a specific coin available for lending.
type LendingCoin struct {
	Coin            string `json:"coin"`
	MaxRedeemQty    string `json:"maxRedeemQty"`
	MinPurchaseQty  string `json:"minPurchaseQty"`
	Precision       string `json:"precision"`
	Rate            string `json:"rate"`
	LoanToPoolRatio string `json:"loanToPoolRatio"`
	ActualApy       string `json:"actualApy"`
}

// LendingCoinInfoResult represents the list of coins available for lending.
type LendingCoinInfoResult struct {
	List []LendingCoin `json:"list"`
}

// DepositFund represents the status of a deposit into a lending product.
type DepositFund struct {
	Coin        string `json:"coin"`
	CreatedTime string `json:"createdTime"`
	OrderId     string `json:"orderId"`
	Quantity    string `json:"quantity"`
	SerialNo    string `json:"serialNo"`
	Status      string `json:"status"`
	UpdatedTime string `json:"updatedTime"`
}

// RedeemFund represents the status of a redemption request from a lending product.
type RedeemFund struct {
	Coin         string `json:"coin"`
	CreatedTime  string `json:"createdTime"`
	OrderId      string `json:"orderId"`
	PrincipalQty string `json:"principalQty"`
	SerialNo     string `json:"serialNo"`
	Status       string `json:"status"`
	UpdatedTime  string `json:"updatedTime"`
}

// CancelRedeemFund represents the status of a cancelled redemption request from a lending product.
type CancelRedeemFund struct {
	OrderId     string `json:"orderId"`
	SerialNo    string `json:"serialNo"`
	UpdatedTime string `json:"updatedTime"`
}

// LendingOrderRecord represents an individual lending order record.
type LendingOrderRecord struct {
	Coin        string `json:"coin"`
	CreatedTime string `json:"createdTime"`
	OrderId     string `json:"orderId"`
	Quantity    string `json:"quantity"`
	SerialNo    string `json:"serialNo"`
	Status      string `json:"status"`
	UpdatedTime string `json:"updatedTime"`
}

// LendingOrdersRecordsResult represents the result of a query for lending order records.
type LendingOrdersRecordsResult struct {
	List []LendingOrderRecord `json:"list"`
}
