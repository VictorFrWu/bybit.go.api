package models

type TransactionLogInfo struct {
	List           []TransactionLogEntry `json:"list"`
	NextPageCursor string                `json:"nextPageCursor"`
}

type TransactionLogEntry struct {
	Symbol          string `json:"symbol"`
	Category        string `json:"category"`
	Side            string `json:"side"`
	TransactionTime string `json:"transactionTime"`
	Type            string `json:"type"`
	Qty             string `json:"qty"`
	Size            string `json:"size"`
	Currency        string `json:"currency"`
	TradePrice      string `json:"tradePrice"`
	Funding         string `json:"funding"`
	Fee             string `json:"fee"`
	CashFlow        string `json:"cashFlow"`
	Change          string `json:"change"`
	CashBalance     string `json:"cashBalance"`
	FeeRate         string `json:"feeRate"`
	BonusChange     string `json:"bonusChange"`
	TradeID         string `json:"tradeId"`
	OrderID         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
}
