package bybit_connector

import (
	"context"
	"net/http"

	"github.com/mudrex/bybit.go.api/handlers"
)

func (s *BybitClientRequest) GetBrokerEarning(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/broker/earnings-info",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetBrokerAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/broker/account-info",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetAllSubMembersDepositRecords(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/broker/asset/query-sub-member-deposit-record",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
