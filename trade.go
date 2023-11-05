package bybit_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

type OrderResult struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type Order struct {
	c                *Client
	category         string
	symbol           string
	isLeverage       *int
	side             string
	orderType        string
	qty              string
	price            *string
	triggerDirection *int
	orderFilter      *string
	triggerPrice     *string
	triggerBy        *string
	orderIv          *string
	timeInForce      *string
	positionIdx      *int
	orderLinkId      *string
	takeProfit       *string
	stopLoss         *string
	tpTriggerBy      *string
	slTriggerBy      *string
	reduceOnly       *bool
	closeOnTrigger   *bool
	smpType          *string
	mmp              *bool
	tpslMode         *string
	tpLimitPrice     *string
	slLimitPrice     *string
	tpOrderType      *string
	slOrderType      *string
}

func (o *Order) TimeInForce(tif string) *Order {
	o.timeInForce = &tif
	return o
}

func (s *Order) IsLeverage(isLeverage int) *Order {
	s.isLeverage = &isLeverage
	return s
}

func (s *Order) TriggerPrice(triggerPrice string) *Order {
	s.triggerPrice = &triggerPrice
	return s
}

func (s *Order) OrderLinkId(orderLinkId string) *Order {
	s.orderLinkId = &orderLinkId
	return s
}

func (o *Order) Price(price string) *Order {
	o.price = &price
	return o
}

func (o *Order) TriggerDirection(direction int) *Order {
	o.triggerDirection = &direction
	return o
}

func (o *Order) OrderFilter(filter string) *Order {
	o.orderFilter = &filter
	return o
}

func (o *Order) TriggerBy(triggerBy string) *Order {
	o.triggerBy = &triggerBy
	return o
}

func (o *Order) OrderIv(iv string) *Order {
	o.orderIv = &iv
	return o
}

func (o *Order) PositionIdx(idx int) *Order {
	o.positionIdx = &idx
	return o
}

func (o *Order) TakeProfit(profit string) *Order {
	o.takeProfit = &profit
	return o
}

func (o *Order) StopLoss(loss string) *Order {
	o.stopLoss = &loss
	return o
}

func (o *Order) TpTriggerBy(triggerBy string) *Order {
	o.tpTriggerBy = &triggerBy
	return o
}

func (o *Order) SlTriggerBy(triggerBy string) *Order {
	o.slTriggerBy = &triggerBy
	return o
}

func (o *Order) ReduceOnly(reduce bool) *Order {
	o.reduceOnly = &reduce
	return o
}

func (o *Order) CloseOnTrigger(close bool) *Order {
	o.closeOnTrigger = &close
	return o
}

func (o *Order) SmpType(smp string) *Order {
	o.smpType = &smp
	return o
}

func (o *Order) Mmp(mmp bool) *Order {
	o.mmp = &mmp
	return o
}

func (o *Order) TpslMode(mode string) *Order {
	o.tpslMode = &mode
	return o
}

func (o *Order) TpLimitPrice(price string) *Order {
	o.tpLimitPrice = &price
	return o
}

func (o *Order) SlLimitPrice(price string) *Order {
	o.slLimitPrice = &price
	return o
}

func (o *Order) TpOrderType(orderType string) *Order {
	o.tpOrderType = &orderType
	return o
}

func (o *Order) SlOrderType(orderType string) *Order {
	o.slOrderType = &orderType
	return o
}

func (s *Order) Do(ctx context.Context, opts ...RequestOption) (*ServerResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/create",
		secType:  secTypeSigned,
	}
	m := params{
		"category":  s.category,
		"symbol":    s.symbol,
		"side":      s.side,
		"orderType": s.orderType,
		"qty":       s.qty,
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.triggerDirection != nil {
		m["triggerDirection"] = *s.triggerDirection
	}
	if s.orderFilter != nil {
		m["orderFilter"] = *s.orderFilter
	}
	if s.triggerPrice != nil {
		m["triggerPrice"] = *s.triggerPrice
	}
	if s.triggerBy != nil {
		m["triggerBy"] = *s.triggerBy
	}
	if s.orderIv != nil {
		m["orderIv"] = *s.orderIv
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.positionIdx != nil {
		m["positionIdx"] = *s.positionIdx
	}
	if s.orderLinkId != nil {
		m["orderLinkId"] = *s.orderLinkId
	}
	if s.takeProfit != nil {
		m["takeProfit"] = *s.takeProfit
	}
	if s.stopLoss != nil {
		m["stopLoss"] = *s.stopLoss
	}
	if s.tpTriggerBy != nil {
		m["tpTriggerBy"] = *s.tpTriggerBy
	}
	if s.slTriggerBy != nil {
		m["slTriggerBy"] = *s.slTriggerBy
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.closeOnTrigger != nil {
		m["closeOnTrigger"] = *s.closeOnTrigger
	}
	if s.smpType != nil {
		m["smpType"] = *s.smpType
	}
	if s.mmp != nil {
		m["mmp"] = *s.mmp
	}
	if s.tpslMode != nil {
		m["tpslMode"] = *s.tpslMode
	}
	if s.tpLimitPrice != nil {
		m["tpLimitPrice"] = *s.tpLimitPrice
	}
	if s.slLimitPrice != nil {
		m["slLimitPrice"] = *s.slLimitPrice
	}
	if s.tpOrderType != nil {
		m["tpOrderType"] = *s.tpOrderType
	}
	if s.slOrderType != nil {
		m["slOrderType"] = *s.slOrderType
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
