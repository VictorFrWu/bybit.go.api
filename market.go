package bybit

import (
	"context"
	"net/http"
)

// Binance Check Server Time endpoint (GET /v5/market/time)
type ServerTime struct {
	c *Client
}

// Send the request
func (s *ServerTime) Do(ctx context.Context, opts ...RequestOption) (res []byte) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/time",
		secType:  secTypeNone,
	}
	data, _ := s.c.callAPI(ctx, r, opts...)
	res = data
	return res
}
