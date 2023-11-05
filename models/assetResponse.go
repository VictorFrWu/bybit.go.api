package models

type DeliveryRecordInfo struct {
	Category       string                `json:"category"`
	List           []DeliveryRecordEntry `json:"list"`
	NextPageCursor string                `json:"nextPageCursor"`
}

type DeliveryRecordEntry struct {
	DeliveryTime  int64  `json:"deliveryTime"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Position      string `json:"position"`
	DeliveryPrice string `json:"deliveryPrice"`
	Strike        string `json:"strike"`
	Fee           string `json:"fee"`
	DeliveryPnl   string `json:"deliveryRpl"`
}

type USDCSettlementInfo struct {
	Category       string            `json:"category"`
	List           []SettlementEntry `json:"list"`
	NextPageCursor string            `json:"nextPageCursor"`
}

type SettlementEntry struct {
	Symbol          string `json:"symbol"`
	Side            string `json:"side"`
	Size            string `json:"size"`
	SessionAvgPrice string `json:"sessionAvgPrice"`
	MarkPrice       string `json:"markPrice"`
	RealisedPnl     string `json:"realisedPnl"`
	CreatedTime     string `json:"createdTime"`
}
