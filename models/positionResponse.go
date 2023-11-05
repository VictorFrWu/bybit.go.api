package models

type PositionListInfo struct {
	Category string `json:"category"`
	List     []struct {
		PositionIdx            int    `json:"positionIdx"`
		RiskId                 int    `json:"riskId"`
		RiskLimitValue         string `json:"riskLimitValue"`
		Symbol                 string `json:"symbol"`
		Side                   string `json:"side"`
		Size                   string `json:"size"`
		AvgPrice               string `json:"avgPrice"`
		PositionValue          string `json:"positionValue"`
		TradeMode              int    `json:"tradeMode"`
		AutoAddMargin          int    `json:"autoAddMargin"`
		PositionStatus         string `json:"positionStatus"`
		Leverage               string `json:"leverage"`
		MarkPrice              string `json:"markPrice"`
		LiqPrice               string `json:"liqPrice"`
		BustPrice              string `json:"bustPrice"`
		PositionIM             string `json:"positionIM"`
		PositionMM             string `json:"positionMM"`
		PositionBalance        string `json:"positionBalance"`
		TpslMode               string `json:"tpslMode"`
		TakeProfit             string `json:"takeProfit"`
		StopLoss               string `json:"stopLoss"`
		TrailingStop           string `json:"trailingStop"`
		UnrealisedPnl          string `json:"unrealisedPnl"`
		CumRealisedPnl         string `json:"cumRealisedPnl"`
		AdlRankIndicator       int    `json:"adlRankIndicator"`
		IsReduceOnly           bool   `json:"isReduceOnly"`
		MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
		LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
		CreatedTime            string `json:"createdTime"`
		UpdatedTime            string `json:"updatedTime"`
		Seq                    int64  `json:"seq"`
	} `json:"list"`
	NextPageCursor string `json:"nextPageCursor"`
}

type PositionTpslMode struct {
	TpSlMode string `json:"tpSlMode"`
}

type PositionRiskInfo struct {
	Category       string `json:"category"`
	RiskId         int64  `json:"riskId"`
	RiskLimitValue string `json:"riskLimitValue"`
}

type PositionUpdateMargin struct {
	Category       string `json:"category"`
	Symbol         string `json:"symbol"`
	PositionIdx    int    `json:"positionIdx"`
	RiskId         int    `json:"riskId"`
	RiskLimitValue string `json:"riskLimitValue"`
	Size           string `json:"size"`
	AvgPrice       string `json:"avgPrice"`
	LiqPrice       string `json:"liqPrice"`
	BustPrice      string `json:"bustPrice"`
	MarkPrice      string `json:"markPrice"`
	PositionValue  string `json:"positionValue"`
	Leverage       string `json:"leverage"`
	AutoAddMargin  int    `json:"autoAddMargin"`
	PositionStatus string `json:"positionStatus"`
	PositionIM     string `json:"positionIM"`
	PositionMM     string `json:"positionMM"`
	TakeProfit     string `json:"takeProfit"`
	StopLoss       string `json:"stopLoss"`
	TrailingStop   string `json:"trailingStop"`
	UnrealisedPnl  string `json:"unrealisedPnl"`
	CumRealisedPnl string `json:"cumRealisedPnl"`
	CreatedTime    string `json:"createdTime"`
	UpdatedTime    string `json:"updatedTime"`
}

type PositionExecutionInfo struct {
	Category string                   `json:"category"`
	List     []PositionExecutionEntry `json:"list"`
}

type PositionExecutionEntry struct {
	Symbol          string `json:"symbol"`
	OrderID         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	Side            string `json:"side"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	ExecFee         string `json:"execFee"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	IsMaker         bool   `json:"isMaker"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	BlockTradeId    string `json:"blockTradeId"`
	ClosedSize      string `json:"closedSize"`
	Seq             int64  `json:"seq"`
	NextPageCursor  string `json:"nextPageCursor"`
}

type PositionClosedPnlInfo struct {
	Category       string               `json:"category"`
	List           []ClosedPnlInfoEntry `json:"list"`
	NextPageCursor string               `json:"nextPageCursor"`
}

type ClosedPnlInfoEntry struct {
	Symbol        string `json:"symbol"`
	OrderID       string `json:"orderId"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	OrderPrice    string `json:"orderPrice"`
	OrderType     string `json:"orderType"`
	ExecType      string `json:"execType"`
	ClosedSize    string `json:"closedSize"`
	CumEntryValue string `json:"cumEntryValue"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	CumExitValue  string `json:"cumExitValue"`
	AvgExitPrice  string `json:"avgExitPrice"`
	ClosedPnl     string `json:"closedPnl"`
	FillCount     string `json:"fillCount"`
	Leverage      string `json:"leverage"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}
