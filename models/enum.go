package models

type Category string

const (
	CategorySpot    Category = "spot"
	CategoryLinear  Category = "linear"  //Unified Account: USDT perpetual, and USDC contract, including USDC perp, USDC futures; Classic Account: USDT perp
	CategoryInverse Category = "inverse" // Inverse contract, including Inverse perp, Inverse futures
	CategoryOption  Category = "option"
)

type OrderStatus string

const (
	OrderStatusCreated                 OrderStatus = "Created" // order has been accepted by the system but not yet put through the matching engine
	OrderStatusNew                     OrderStatus = "New"     // order has been placed successfully
	OrderStatusRejected                OrderStatus = "Rejected"
	OrderStatusPartiallyFilled         OrderStatus = "PartiallyFilled"
	OrderStatusPartiallyFilledCanceled OrderStatus = "PartiallyFilledCanceled" //Only spot has this order status
	OrderStatusFilled                  OrderStatus = "Filled"
	OrderStatusCancelled               OrderStatus = "Cancelled" // In derivatives, orders with this status may have an executed qty
	OrderStatusUntriggered             OrderStatus = "Untriggered"
	OrderStatusTriggered               OrderStatus = "Triggered"
	OrderStatusDeactivated             OrderStatus = "Deactivated" // UTA: Spot tp/sl order, conditional order, OCO order are cancelled before they are triggered
)

type TimeInForce string

const (
	TimeInForceGTC      TimeInForce = "GTC" //GoodTillCancel
	TimeInForceIOC      TimeInForce = "IOC" //ImmediateOrCancel
	TimeInForceFOK      TimeInForce = "FOK" //FillOrKill
	TimeInForcePostOnly TimeInForce = "PostOnly"
)

type OrderType string

const (
	OrderTypeMarket OrderType = "Market"
	Limit           OrderType = "Limit"
	UNKNOWN         OrderType = "UNKNOWN" // Is not a valid request parameter value. Is only used in some responses. Mainly, it is used when execType is Funding.
)

type SymbolStatus string

const (
	SymbolStatusPreLaunch  SymbolStatus = "PreLaunch"
	SymbolStatusTrading    SymbolStatus = "Trading"
	SymbolStatusDelivering SymbolStatus = "Delivering"
	SymbolStatusClosed     SymbolStatus = "Closed"
)
