package models

type GetServerTimeResponse struct {
	RetCode    int              `json:"retCode"`
	RetMsg     string           `json:"retMsg"`
	Result     ServerTimeResult `json:"result"`
	RetExtInfo struct{}         `json:"retExtInfo"`
	Time       int64            `json:"time"`
}

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
	Category Category             `json:"category"`
	Symbol   string               `json:"symbol"`
	List     []*MarketKlineCandle `json:"list"`
}

type MarketMarkPriceKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
}

type MarketMarkPriceKlineResponse struct {
	Category Category                      `json:"category"`
	Symbol   string                        `json:"symbol"`
	List     []*MarketMarkPriceKlineCandle `json:"list"`
}

type MarketIndexPriceKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
}

type MarketIndexPriceKlineResponse struct {
	Category Category                       `json:"category"`
	Symbol   string                         `json:"symbol"`
	List     []*MarketIndexPriceKlineCandle `json:"list"`
}

type MarketPremiumIndexPriceKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
}

type MarketPremiumIndexPriceKlineResponse struct {
	Category Category                              `json:"category"`
	Symbol   string                                `json:"symbol"`
	List     []*MarketPremiumIndexPriceKlineCandle `json:"list"`
}

type InstrumentInfoResponse struct {
	Category       Category      `json:"category"`
	NextPageCursor string        `json:"nextPageCursor"`
	List           []interface{} `json:"list"`
}

type SpotInstrument struct {
	Symbol             string         `json:"symbol"`
	ContractType       string         `json:"contractType"`
	OptionType         string         `json:"optionType"`
	Innovation         string         `json:"innovation"`
	Status             SymbolStatus   `json:"status"`
	BaseCoin           string         `json:"baseCoin"`
	QuoteCoin          string         `json:"quoteCoin"`
	LaunchTime         string         `json:"launchTime"`
	DeliveryTime       string         `json:"deliveryTime"`
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`
	PriceScale         string         `json:"priceScale"`
	MarginTrading      string         `json:"marginTrading"`
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
	BasePrecision       string `json:"basePrecision"`
	QuotePrecision      string `json:"quotePrecision"`
	MaxOrderAmt         string `json:"maxOrderAmt"`
	MinOrderAmt         string `jsoN:"minOrderAmt"`
}

type MarketOrderBookResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     OrderBookInfo `json:"result"`
	RetExtInfo struct{}      `json:"retExtInfo"`
	Time       int64         `json:"time"`
}

// type OrderBookEntry struct {
// Price string `json:"0"`
// Size  string `json:"1"`
// }

type OrderBookEntry []string

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
	Ask1Iv                 string `json:"ask1Iv"`
	Bid1Iv                 string `json:"bid1Iv"`
	MarkIv                 string `json:"markIv"`
	UnderlyingPrice        string `json:"underlyingPrice"`
	TotalVolume            string `json:"totalVolume"`
	TotalTurnover          string `json:"totalTurnover"`
	Change24h              string `json:"change24h"`
	UsdIndexPrice          string `json:"usdIndexPrice"`
}

type MarketTickersResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     MarketTickers `json:"result"`
	RetExtInfo struct{}      `json:"retExtInfo"`
	Time       int64         `json:"time"`
}

type MarketTickers struct {
	Category string        `json:"category"`
	List     []*TickerInfo `json:"list"`
}

type MarketFundingRatesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     FundingRate `json:"result"`
	RetExtInfo struct{}    `json:"retExtInfo"`
	Time       int64       `json:"time"`
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

type GetPublicRecentTradesResponse struct {
	RetCode    int                      `json:"retCode"`
	RetMsg     string                   `json:"retMsg"`
	Result     PublicRecentTradeHistory `json:"result"`
	RetExtInfo struct{}                 `json:"retExtInfo"`
	Time       int64                    `json:"time"`
}

type GetOpenInterestsResponse struct {
	RetCode    int              `json:"retCode"`
	RetMsg     string           `json:"retMsg"`
	Result     OpenInterestInfo `json:"result"`
	RetExtInfo struct{}         `json:"retExtInfo"`
	Time       int64            `json:"time"`
}

type OpenInterestInfo struct {
	Category       string         `json:"category"`
	Symbol         string         `json:"symbol"`
	List           []OpenInterest `json:"list"`
	NextPageCursor string         `json:"nextPageCursor"`
}

type OpenInterest struct {
	OpenInterest string `json:"openInterest"`
	Timestamp    string `json:"timeStamp"`
}

type VolatilityData struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

type HistoricalVolatilityInfo struct {
	Category string           `json:"category"`
	List     []VolatilityData `json:"result"`
}

type GetInsuranceInfoResponse struct {
	RetCode    int                 `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	Result     MarketInsuranceInfo `json:"result"`
	RetExtInfo struct{}            `json:"retExtInfo"`
	Time       int64               `json:"time"`
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

type GetRiskLimitResponse struct {
	RetCode    int                 `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	Result     MarketRiskLimitInfo `json:"result"`
	RetExtInfo struct{}            `json:"retExtInfo"`
	Time       int64               `json:"time"`
}

type RiskLimitData struct {
	Id                int    `json:"id"`
	Symbol            string `json:"symbol"`
	RiskLimitValue    string `json:"riskLimitValue"`
	MaintenanceMargin string `json:"maintenanceMargin"`
	InitialMargin     string `json:"initialMargin"`
	IsLowestRisk      int    `json:"isLowestRisk"`
	MaxLeverage       string `json:"maxLeverage"`
}

type MarketRiskLimitInfo struct {
	Category string          `json:"category"`
	List     []RiskLimitData `json:"list"`
}

type GetDeliveryPriceResponse struct {
	RetCode    int               `json:"retCode"`
	RetMsg     string            `json:"retMsg"`
	Result     DeliveryPriceInfo `json:"result"`
	RetExtInfo struct{}          `json:"retExtInfo"`
	Time       int64             `json:"time"`
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

type GetMarketLSRatioResponse struct {
	RetCode    int                      `json:"retCode"`
	RetMsg     string                   `json:"retMsg"`
	Result     MarketLongShortRatioInfo `json:"result"`
	RetExtInfo struct{}                 `json:"retExtInfo"`
	Time       int64                    `json:"time"`
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
