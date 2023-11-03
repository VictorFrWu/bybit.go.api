package bybit_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

type ServerTimeResult struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

// ServerTime Binance Check Server Time endpoint (GET /v5/market/time)
type ServerTime struct {
	c *Client
}

// Do Send the request
func (s *ServerTime) Do(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/time",
		secType:  secTypeNone,
	}
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
