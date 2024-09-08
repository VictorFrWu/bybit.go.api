package bybit_connector

import (
	"context"
	"github.com/wuhewuhe/bybit.go.api/handlers"
	"net/http"
)

func (s *BybitClientRequest) GetInsLoanInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/product-infos",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetInsMarginCoinInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/ensure-tokens-convert",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetInsLoanOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/loan-order",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetInsRepayOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/repaid-history",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetInsLoanToValue(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/ins-loan/ltv-convert",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) AssociateInsLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/ins-loan/association-uid",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetC2cLendingCoinInfo is deprecated.
func (s *BybitClientRequest) GetC2cLendingCoinInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/lending/info",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetC2cLendingOrders is deprecated.
func (s *BybitClientRequest) GetC2cLendingOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/lending/history-order",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetC2cLendingAccountInfo is deprecated.
func (s *BybitClientRequest) GetC2cLendingAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/lending/account",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: C2cDepositFunds is deprecated.
func (s *BybitClientRequest) C2cDepositFunds(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/lending/purchase",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: C2cRedeemFunds is deprecated.
func (s *BybitClientRequest) C2cRedeemFunds(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/lending/redeem",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: C2cCancelRedeemFunds is deprecated.
func (s *BybitClientRequest) C2cCancelRedeemFunds(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/lending/redeem-cancel",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
