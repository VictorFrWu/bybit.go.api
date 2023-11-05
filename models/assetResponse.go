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

type AssetInfo struct {
	Spot SpotInfo `json:"spot"`
}

type SpotInfo struct {
	Status string      `json:"status"`
	Assets []AssetItem `json:"assets"`
}

type AssetItem struct {
	Coin     string `json:"coin"`
	Frozen   string `json:"frozen"`
	Free     string `json:"free"`
	Withdraw string `json:"withdraw"`
}

type AllCoinsBalance struct {
	AccountType string          `json:"accountType"`
	MemberId    string          `json:"memberId"`
	Balance     []BalanceDetail `json:"balance"`
}

type BalanceDetail struct {
	Coin            string `json:"coin"`
	WalletBalance   string `json:"walletBalance"`
	TransferBalance string `json:"transferBalance"`
	Bonus           string `json:"bonus"`
}

type SingleCoinBalance struct {
	AccountType string      `json:"accountType"`
	BizType     int         `json:"bizType"`
	AccountId   string      `json:"accountId"`
	MemberId    string      `json:"memberId"`
	Balance     CoinBalance `json:"balance"`
}

type CoinBalance struct {
	Coin                  string `json:"coin"`
	WalletBalance         string `json:"walletBalance"`
	TransferBalance       string `json:"transferBalance"`
	Bonus                 string `json:"bonus"`
	TransferSafeAmount    string `json:"transferSafeAmount"`
	LtvTransferSafeAmount string `json:"ltvTransferSafeAmount"`
}

type TransferResult struct {
	TransferId string `json:"transferId"`
}

type InternalTransferInfo struct {
	List           []TransferDetail `json:"list"`
	NextPageCursor string           `json:"nextPageCursor"`
}

type TransferDetail struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type SubUidsInfo struct {
	SubMemberIds             []string `json:"subMemberIds"`
	TransferableSubMemberIds []string `json:"transferableSubMemberIds"`
}

type UniversalTransferInfo struct {
	List           []UniversalTransferDetail `json:"list"`
	NextPageCursor string                    `json:"nextPageCursor"`
}

type UniversalTransferDetail struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromMemberId    string `json:"fromMemberId"`
	ToMemberId      string `json:"toMemberId"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type WithdrawAssetResult struct {
	Id string `json:"id"`
}

type CancelWithdrawAssetResult struct {
	Status string `json:"status"`
}

type SetDepositAccountResult struct {
	Status string `json:"status"`
}

type AllowDepositCoinInfo struct {
	ConfigList     []DepositCoinConfig `json:"configList"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type DepositCoinConfig struct {
	Coin               string `json:"coin"`
	Chain              string `json:"chain"`
	CoinShowName       string `json:"coinShowName"`
	ChainType          string `json:"chainType"`
	BlockConfirmNumber int    `json:"blockConfirmNumber"`
	MinDepositAmount   string `json:"minDepositAmount"`
}

type DepositRecords struct {
	Rows           []DepositRecord `json:"rows"`
	NextPageCursor string          `json:"nextPageCursor"`
}

type DepositRecord struct {
	Coin              string `json:"coin"`
	Chain             string `json:"chain"`
	Amount            string `json:"amount"`
	TxID              string `json:"txID"`
	Status            int    `json:"status"`
	ToAddress         string `json:"toAddress"`
	Tag               string `json:"tag"`
	DepositFee        string `json:"depositFee"`
	SuccessAt         string `json:"successAt"`
	Confirmations     string `json:"confirmations"`
	TxIndex           string `json:"txIndex"`
	BlockHash         string `json:"blockHash"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
	DepositType       int    `json:"depositType"`
}

type SubDepositResult struct {
	Coin   string             `json:"coin"`
	Chains []DepositChainInfo `json:"chains"`
}

// MasterDepositResult represents the structure for master deposit results.
type MasterDepositResult struct {
	Coin   string             `json:"coin"`
	Chains []DepositChainInfo `json:"chains"`
}

// DepositChainInfo represents the shared structure for deposit chain information.
type DepositChainInfo struct {
	ChainType         string `json:"chainType"`
	AddressDeposit    string `json:"addressDeposit"`
	TagDeposit        string `json:"tagDeposit"`
	Chain             string `json:"chain"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
}

// CoinInfoResult represents the structure for coin info results.
type CoinInfoResult struct {
	Rows []CoinInfoRow `json:"rows"`
}

// CoinInfoRow represents the structure for each row of coin information.
type CoinInfoRow struct {
	Name         string          `json:"name"`
	Coin         string          `json:"coin"`
	RemainAmount string          `json:"remainAmount"`
	Chains       []CoinChainInfo `json:"chains"`
}

// CoinChainInfo represents the structure for each chain's information for a coin.
type CoinChainInfo struct {
	Chain                 string `json:"chain"`
	ChainType             string `json:"chainType"`
	Confirmation          string `json:"confirmation"`
	WithdrawFee           string `json:"withdrawFee"`
	DepositMin            string `json:"depositMin"`
	WithdrawMin           string `json:"withdrawMin"`
	MinAccuracy           string `json:"minAccuracy"`
	ChainDeposit          string `json:"chainDeposit"`
	ChainWithdraw         string `json:"chainWithdraw"`
	WithdrawPercentageFee string `json:"withdrawPercentageFee"`
}

// WithdrawRecords represents the structure for withdrawal records.
type WithdrawRecords struct {
	Rows           []WithdrawRecord `json:"rows"`
	NextPageCursor string           `json:"nextPageCursor"`
}

// WithdrawRecord represents the structure for each withdrawal record.
type WithdrawRecord struct {
	WithdrawID   string `json:"withdrawId"`
	TxID         string `json:"txID"`
	WithdrawType string `json:"withdrawType"`
	Coin         string `json:"coin"`
	Chain        string `json:"chain"`
	Amount       string `json:"amount"`
	WithdrawFee  string `json:"withdrawFee"`
	Status       string `json:"status"`
	ToAddress    string `json:"toAddress"`
	Tag          string `json:"tag"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

// WithdrawableAmount represents the structure for information about withdrawable amounts.
type WithdrawableAmount struct {
	LimitAmountUsd      string              `json:"limitAmountUsd"`
	WithdrawableAmounts []WithdrawableAsset `json:"withdrawableAmount"`
}

// WithdrawableAsset represents the structure for withdrawable amounts for each wallet type.
type WithdrawableAsset struct {
	SPOT *WalletInfo `json:"SPOT,omitempty"` // Omitted if empty
	FUND *WalletInfo `json:"FUND,omitempty"` // Omitted if empty
}

// WalletInfo represents the wallet information for withdrawable assets.
type WalletInfo struct {
	Coin               string `json:"coin"`
	WithdrawableAmount string `json:"withdrawableAmount"`
	AvailableBalance   string `json:"availableBalance"`
}
