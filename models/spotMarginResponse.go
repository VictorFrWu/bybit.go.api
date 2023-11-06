package models

// SpotMarginDataResult holds data for spot margin including VIP level specific coin list.
type SpotMarginDataResult struct {
	VipCoinList []struct {
		List []struct {
			Borrowable         bool   `json:"borrowable"`
			CollateralRatio    string `json:"collateralRatio"`
			Currency           string `json:"currency"`
			HourlyBorrowRate   string `json:"hourlyBorrowRate"`
			LiquidationOrder   string `json:"liquidationOrder"`
			MarginCollateral   bool   `json:"marginCollateral"`
			MaxBorrowingAmount string `json:"maxBorrowingAmount"`
		} `json:"list"`
		VipLevel string `json:"vipLevel"`
	} `json:"vipCoinList"`
}

// ClassicalSpotMarginCoinResult holds information about coins in classical spot margin.
type ClassicalSpotMarginCoinResult struct {
	List []struct {
		Coin             string `json:"coin"`
		ConversionRate   string `json:"conversionRate"`
		LiquidationOrder int    `json:"liquidationOrder"`
	} `json:"list"`
}

// ClassicalSpotMarginBorrowCoinResult holds the borrowing precision information for coins in classical spot margin.
type ClassicalSpotMarginBorrowCoinResult struct {
	List []struct {
		Coin               string `json:"coin"`
		BorrowingPrecision int    `json:"borrowingPrecision"`
		RepaymentPrecision int    `json:"repaymentPrecision"`
	} `json:"list"`
}

// ClassicalSpotMarginInterestResult contains information about the interest rate on spot margin.
type ClassicalSpotMarginInterestResult struct {
	Coin           string `json:"coin"`
	InterestRate   string `json:"interestRate"`
	LoanAbleAmount string `json:"loanAbleAmount"`
	MaxLoanAmount  string `json:"maxLoanAmount"`
}

// ClassicalSpotMarginLoanResult holds the loan account information in the classical spot margin.
type ClassicalSpotMarginLoanResult struct {
	AcctBalanceSum  string `json:"acctBalanceSum"`
	DebtBalanceSum  string `json:"debtBalanceSum"`
	LoanAccountList []struct {
		Free         string `json:"free"`
		Interest     string `json:"interest"`
		Loan         string `json:"loan"`
		RemainAmount string `json:"remainAmount"`
		Locked       string `json:"locked"`
		TokenId      string `json:"tokenId"`
		Total        string `json:"total"`
	} `json:"loanAccountList"`
	RiskRate     string `json:"riskRate"`
	Status       int    `json:"status"`       // Use int for integer type
	SwitchStatus int    `json:"switchStatus"` // Use int for integer type
}

type SpotMarginBorrowOrders struct {
	List []struct {
		AccountId       string `json:"accountId"`       // Account ID
		Coin            string `json:"coin"`            // Coin name
		CreatedTime     int64  `json:"createdTime"`     // Borrow order created timestamp (ms)
		Id              string `json:"id"`              // Borrow order ID
		InterestAmount  string `json:"interestAmount"`  // Total interest
		InterestBalance string `json:"interestBalance"` // Outstanding interest
		LoanAmount      string `json:"loanAmount"`      // Principal amount
		LoanBalance     string `json:"loanBalance"`     // Outstanding principal
		RemainAmount    string `json:"remainAmount"`    // Remaining debt = interestBalance + loanBalance
		Status          int    `json:"status"`          // Status 1: uncleared, 2: cleared
		Type            int    `json:"type"`            // Order Type 1: manual loan, 2: auto loan
	} `json:"list"`
}

type SpotBorrowOrderResult struct {
	TransactId string `json:"transactId"`
}

type SpotRepayOrderResult struct {
	RepayId string `json:"repayId"`
}

type SpotToggleMarginResult struct {
	SwitchStatus string `json:"switchStatus"`
}

type SpotMarginLeverageResult struct {
	Leverage string `json:"leverage"`
}

type SpotMarginStateResult struct {
	SpotLeverage   string `json:"spotLeverage"`
	SpotMarginMode string `json:"spotMarginMode"`
}
