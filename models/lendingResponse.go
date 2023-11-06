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

type MarginProductInfo struct {
	ProductId              string         `json:"productId"`
	Leverage               string         `json:"leverage"`
	SupportSpot            int            `json:"supportSpot"`
	SupportContract        int            `json:"supportContract"`
	SupportMarginTrading   int            `json:"supportMarginTrading"`
	WithdrawLine           string         `json:"withdrawLine"`
	TransferLine           string         `json:"transferLine"`
	SpotBuyLine            string         `json:"spotBuyLine"`
	SpotSellLine           string         `json:"spotSellLine"`
	ContractOpenLine       string         `json:"contractOpenLine"`
	LiquidationLine        string         `json:"liquidationLine"`
	StopLiquidationLine    string         `json:"stopLiquidationLine"`
	ContractLeverage       string         `json:"contractLeverage"`
	TransferRatio          string         `json:"transferRatio"`
	SpotSymbols            []string       `json:"spotSymbols"`
	ContractSymbols        []string       `json:"contractSymbols"`
	SupportUSDCContract    int            `json:"supportUSDCContract"`
	SupportUSDCOptions     int            `json:"supportUSDCOptions"`
	USDTPerpetualOpenLine  string         `json:"USDTPerpetualOpenLine"`
	USDCContractOpenLine   string         `json:"USDCContractOpenLine"`
	USDCOptionsOpenLine    string         `json:"USDCOptionsOpenLine"`
	USDTPerpetualCloseLine string         `json:"USDTPerpetualCloseLine"`
	USDCContractCloseLine  string         `json:"USDCContractCloseLine"`
	USDCOptionsCloseLine   string         `json:"USDCOptionsCloseLine"`
	USDCContractSymbols    []string       `json:"USDCContractSymbols"`
	USDCOptionsSymbols     []string       `json:"USDCOptionsSymbols"`
	MarginLeverage         string         `json:"marginLeverage"`
	USDTPerpetualLeverage  []LeverageInfo `json:"USDTPerpetualLeverage"`
	USDCContractLeverage   []LeverageInfo `json:"USDCContractLeverage"`
}

type LeverageInfo struct {
	Symbol   string `json:"symbol"`
	Leverage string `json:"leverage"`
}

type TokenInfo struct {
	Token            string         `json:"token"`
	ConvertRatioList []ConvertRatio `json:"convertRatioList"`
}

type ConvertRatio struct {
	Ladder       string `json:"ladder"`
	ConvertRatio string `json:"convertRatio"`
}

type MarginToken struct {
	ProductId string      `json:"productId"`
	TokenInfo []TokenInfo `json:"tokenInfo"`
}

// LoanInfo To do repay & loand to value
type LoanInfo struct {
	OrderId               string   `json:"orderId"`
	OrderProductId        string   `json:"orderProductId"`
	ParentUid             string   `json:"parentUid"`
	LoanTime              string   `json:"loanTime"`
	LoanCoin              string   `json:"loanCoin"`
	LoanAmount            string   `json:"loanAmount"`
	UnpaidAmount          string   `json:"unpaidAmount"`
	UnpaidInterest        string   `json:"unpaidInterest"`
	RepaidAmount          string   `json:"repaidAmount"`
	RepaidInterest        string   `json:"repaidInterest"`
	InterestRate          string   `json:"interestRate"`
	Status                string   `json:"status"`
	Leverage              string   `json:"leverage"`
	SupportSpot           string   `json:"supportSpot"`
	SupportContract       string   `json:"supportContract"`
	WithdrawLine          string   `json:"withdrawLine"`
	TransferLine          string   `json:"transferLine"`
	SpotBuyLine           string   `json:"spotBuyLine"`
	SpotSellLine          string   `json:"spotSellLine"`
	ContractOpenLine      string   `json:"contractOpenLine"`
	LiquidationLine       string   `json:"liquidationLine"`
	StopLiquidationLine   string   `json:"stopLiquidationLine"`
	ContractLeverage      string   `json:"contractLeverage"`
	TransferRatio         string   `json:"transferRatio"`
	SpotSymbols           []string `json:"spotSymbols"`
	ContractSymbols       []string `json:"contractSymbols"`
	SupportUSDCContract   string   `json:"supportUSDCContract"`
	SupportUSDCOptions    string   `json:"supportUSDCOptions"`
	SupportMarginTrading  string   `json:"supportMarginTrading"`
	USDTPerpetualOpenLine string   `json:"USDTPerpetualOpenLine"`
	USDCContractOpenLine  string   `json:"USDCContractOpenLine"`
}
