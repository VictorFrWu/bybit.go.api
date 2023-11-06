package bybit_connector

import (
	"context"
	"encoding/json"
	"github.com/wuhewuhe/bybit.go.api/handlers"
	"net/http"
)

type BrokerServiceClient struct {
	c      *Client
	params map[string]interface{}
}

func (s *BrokerServiceClient) GetBrokerEarning(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/broker/earning-record",
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
