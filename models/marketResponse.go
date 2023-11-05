package models

type ServerTimeResult struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

type MarketKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
	Volume     string `json:"volume"`
	Turnover   string `json:"turnover"`
}

type MarketKlineResponse struct {
	Category string              `json:"category"`
	Symbol   string              `json:"symbol"`
	List     []MarketKlineCandle `json:"list"`
}

type InstrumentInfo struct {
	Category       string       `json:"category"`
	NextPageCursor string       `json:"nextPageCursor"`
	List           []Instrument `json:"list"`
}

type Instrument struct {
	Symbol             string         `json:"symbol"`
	ContractType       string         `json:"contractType"`
	Status             string         `json:"status"`
	BaseCoin           string         `json:"baseCoin"`
	QuoteCoin          string         `json:"quoteCoin"`
	LaunchTime         string         `json:"launchTime"`
	DeliveryTime       string         `json:"deliveryTime"`
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`
	PriceScale         string         `json:"priceScale"`
	LeverageFilter     LeverageFilter `json:"leverageFilter"`
	PriceFilter        PriceFilter    `json:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `json:"lotSizeFilter"`
	UnifiedMarginTrade bool           `json:"unifiedMarginTrade"`
	FundingInterval    int            `json:"fundingInterval"`
	SettleCoin         string         `json:"settleCoin"`
	CopyTrading        string         `json:"copyTrading"`
}

type LeverageFilter struct {
	MinLeverage  string `json:"minLeverage"`
	MaxLeverage  string `json:"maxLeverage"`
	LeverageStep string `json:"leverageStep"`
}

type PriceFilter struct {
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`
}

type LotSizeFilter struct {
	MaxOrderQty         string `json:"maxOrderQty"`
	MinOrderQty         string `json:"minOrderQty"`
	QtyStep             string `json:"qtyStep"`
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
}

type OrderBookEntry struct {
	Price string `json:"0"`
	Size  string `json:"1"`
}

type OrderBookInfo struct {
	Symbol    string           `json:"s"`
	Bids      []OrderBookEntry `json:"b"`
	Asks      []OrderBookEntry `json:"a"`
	Timestamp int64            `json:"ts"`
	UpdateID  int64            `json:"u"`
}

type TickerInfo struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice"`
	MarkPrice              string `json:"markPrice"`
	PrevPrice24h           string `json:"prevPrice24h"`
	Price24hPcnt           string `json:"price24hPcnt"`
	HighPrice24h           string `json:"highPrice24h"`
	LowPrice24h            string `json:"lowPrice24h"`
	PrevPrice1h            string `json:"prevPrice1h"`
	OpenInterest           string `json:"openInterest"`
	OpenInterestValue      string `json:"openInterestValue"`
	Turnover24h            string `json:"turnover24h"`
	Volume24h              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate"`
	NextFundingTime        string `json:"nextFundingTime"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
	BasisRate              string `json:"basisRate"`
	Basis                  string `json:"basis"`
	DeliveryFeeRate        string `json:"deliveryFeeRate"`
	DeliveryTime           string `json:"deliveryTime"`
	Ask1Size               string `json:"ask1Size"`
	Bid1Price              string `json:"bid1Price"`
	Ask1Price              string `json:"ask1Price"`
	Bid1Size               string `json:"bid1Size"`
}

type MarketTickers struct {
	Category string       `json:"category"`
	List     []TickerInfo `json:"list"`
}

type FundingRateInfo struct {
	Symbol               string `json:"symbol"`
	FundingRate          string `json:"fundingRate"`
	FundingRateTimestamp string `json:"fundingRateTimestamp"`
}

type FundingRate struct {
	Category string            `json:"category"`
	List     []FundingRateInfo `json:"list"`
}

type TradeInfo struct {
	ExecId       string `json:"execId"`
	Symbol       string `json:"symbol"`
	Price        string `json:"price"`
	Size         string `json:"size"`
	Side         string `json:"side"`
	Time         string `json:"time"`
	IsBlockTrade bool   `json:"isBlockTrade"`
}

type PublicRecentTradeHistory struct {
	Category string      `json:"category"`
	List     []TradeInfo `json:"list"`
}

type VolatilityData struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

type HistoricalVolatilityInfo struct {
	Category string           `json:"category"`
	List     []VolatilityData `json:"list"`
}

type InsuranceData struct {
	Coin    string `json:"coin"`
	Balance string `json:"balance"`
	Value   string `json:"value"`
}

type MarketInsuranceInfo struct {
	UpdatedTime string          `json:"updatedTime"`
	List        []InsuranceData `json:"list"`
}

type RiskLimitData struct {
	Id                int     `json:"id"`
	Symbol            string  `json:"symbol"`
	RiskLimitValue    string  `json:"riskLimitValue"`
	MaintenanceMargin float64 `json:"maintenanceMargin"`
	InitialMargin     float64 `json:"initialMargin"`
	IsLowestRisk      int     `json:"isLowestRisk"`
	MaxLeverage       string  `json:"maxLeverage"`
}

type MarketRiskLimitInfo struct {
	Category string          `json:"category"`
	List     []RiskLimitData `json:"list"`
}

type DeliveryPriceData struct {
	Symbol        string `json:"symbol"`
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryTime  string `json:"deliveryTime"`
}

type DeliveryPriceInfo struct {
	Category       string              `json:"category"`
	List           []DeliveryPriceData `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type LongShortRatioData struct {
	Symbol    string `json:"symbol"`
	BuyRatio  string `json:"buyRatio"`
	SellRatio string `json:"sellRatio"`
	Timestamp string `json:"timestamp"`
}

type MarketLongShortRatioInfo struct {
	List []LongShortRatioData `json:"list"`
}
