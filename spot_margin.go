package bybit_connector

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wuhewuhe/bybit.go.api/handlers"
	"net/http"
)

type SpotMarginClient struct {
	c      *Client
	isUta  bool
	params map[string]interface{}
}

func (s *SpotMarginClient) GetSpotMarginData(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	var endpoint string
	if !s.isUta {
		endpoint = "/v5/spot-margin-trade/data"
	} else {
		endpoint = "/v5/spot-cross-margin-trade/data"
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SpotMarginClient) GetSpotMarginCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/pledge-token",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SpotMarginClient) GetSpotMarginBorrowCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/borrow-token",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SpotMarginClient) GetSpotMarginInterests(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/loan-info",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SpotMarginClient) GetSpotMarginLoanAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/account",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
