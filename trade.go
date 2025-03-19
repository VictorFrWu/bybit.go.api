package bybit_connector

import (
	"context"
	"net/http"

	"github.com/mudrex/bybit.go.api/handlers"
	"github.com/mudrex/bybit.go.api/models"
)

func (s *BybitClientRequest) PlaceOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/create",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) AmendOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/amend",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) CancelOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/cancel",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetOpenOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/order/realtime",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetOrderHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/order/history",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpotBorrowQuota(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/order/spot-borrow-check",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) CancelAllOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/cancel-all",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) SetDisconnectCancelAll(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/disconnected-cancel-all",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) PlaceBatchOrder(ctx context.Context, opts ...RequestOption) (res *models.BatchOrderServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/create-batch",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetBatchOrderServerResponse(err, data)
}

func (s *BybitClientRequest) AmendBatchOrder(ctx context.Context, opts ...RequestOption) (res *models.BatchOrderServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/amend-batch",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetBatchOrderServerResponse(err, data)
}

func (s *BybitClientRequest) CancelBatchOrder(ctx context.Context, opts ...RequestOption) (res *models.BatchOrderServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/cancel-batch",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetBatchOrderServerResponse(err, data)
}

func (s *BybitClientRequest) GetTradeHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/execution/list",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) RequestTestFund(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/demo-apply-money",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
