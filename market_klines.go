package bybit_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Klines Market Kline (GET /v5/market/kline)
type Klines struct {
	c         *Client
	klineType string
	category  string
	symbol    string
	interval  string
	limit     *int
	start     *uint64
	end       *uint64
}

// Limit set limit
func (s *Klines) Limit(limit int) *Klines {
	s.limit = &limit
	return s
}

// Start set startTime
func (s *Klines) Start(startTime uint64) *Klines {
	s.start = &startTime
	return s
}

// End set endTime
func (s *Klines) End(endTime uint64) *Klines {
	s.end = &endTime
	return s
}

func (s *Klines) Do(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/" + s.klineType,
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.start != nil {
		r.setParam("start", *s.start)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
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
