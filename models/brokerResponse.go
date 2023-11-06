package models

// BrokerEarningInfo represents an individual record of broker earnings.
type BrokerEarningInfo struct {
	UserId   string `json:"userId"`
	BizType  string `json:"bizType"`
	Symbol   string `json:"symbol"`
	Coin     string `json:"coin"`
	Earning  string `json:"earning"`
	OrderId  string `json:"orderId"`
	ExecTime string `json:"execTime"`
}

// BrokerEarningResult represents the paginated result of broker earnings.
type BrokerEarningResult struct {
	List           []BrokerEarningInfo `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}
