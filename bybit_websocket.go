package bybit_connector

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type MessageHandler func(message string) error

func (b *WebSocket) handleIncomingMessages() {
	for {
		_, message, err := b.conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading:", err)
			b.isConnected = false
			return
		}

		if b.onMessage != nil {
			err := b.onMessage(string(message))
			if err != nil {
				fmt.Println("Error handling message:", err)
				return
			}
		}
	}
}

func (b *WebSocket) monitorConnection(args []string) {
	ticker := time.NewTicker(time.Second * 5) // Check every 5 seconds
	defer ticker.Stop()

	for {
		<-ticker.C
		if !b.isConnected && b.ctx.Err() == nil { // Check if disconnected and context not done
			fmt.Println("Attempting to reconnect...")
			err := b.Connect(args) // Example, adjust parameters as needed
			if err != nil {
				fmt.Println("Reconnection failed:", err)
			} else {
				b.isConnected = true
				go b.handleIncomingMessages() // Restart message handling
			}
		}

		select {
		case <-b.ctx.Done():
			return // Stop the routine if context is done
		default:
		}
	}
}

func (b *WebSocket) SetMessageHandler(handler MessageHandler) {
	b.onMessage = handler
}

type WebSocket struct {
	conn         *websocket.Conn
	url          string
	apiKey       string
	apiSecret    string
	maxAliveTime string
	pingInterval int
	onMessage    MessageHandler
	ctx          context.Context
	cancel       context.CancelFunc
	isConnected  bool
}

type WebsocketOption func(*WebSocket)

func WithPingInterval(pingInterval int) WebsocketOption {
	return func(c *WebSocket) {
		c.pingInterval = pingInterval
	}
}

func WithMaxAliveTime(maxAliveTime string) WebsocketOption {
	return func(c *WebSocket) {
		c.maxAliveTime = maxAliveTime
	}
}

func NewBybitPrivateWebSocket(url, apiKey, apiSecret string, handler MessageHandler, options ...WebsocketOption) *WebSocket {
	c := &WebSocket{
		url:          url,
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		maxAliveTime: "",
		pingInterval: 20,
		onMessage:    handler,
	}

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func NewBybitPublicWebSocket(url string, handler MessageHandler) *WebSocket {
	c := &WebSocket{
		url:          url,
		pingInterval: 20, // default is 20 seconds
		onMessage:    handler,
	}

	return c
}

func (b *WebSocket) Connect(args []string) error {
	var err error
	wssUrl := b.url
	if b.maxAliveTime != "" {
		wssUrl += "?max_alive_time=" + b.maxAliveTime
	}
	b.conn, _, err = websocket.DefaultDialer.Dial(wssUrl, nil)
	if err != nil {
		return err
	}

	if b.requiresAuthentication() {
		if err = b.sendAuth(); err != nil {
			return err
		}
	}
	b.isConnected = true

	go b.handleIncomingMessages()
	go b.monitorConnection(args)

	b.ctx, b.cancel = context.WithCancel(context.Background())
	Ping(b)

	return b.sendSubscription(args)
}

func Ping(b *WebSocket) {
	if b.pingInterval <= 0 {
		fmt.Println("Ping interval is set to a non-positive value.")
		return
	}

	ticker := time.NewTicker(time.Duration(b.pingInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now().Unix()
			pingMessage := map[string]string{
				"op":     "ping",
				"req_id": fmt.Sprintf("%d", currentTime),
			}
			jsonPingMessage, err := json.Marshal(pingMessage)
			if err != nil {
				fmt.Println("Failed to marshal ping message:", err)
				continue
			}
			if err := b.conn.WriteMessage(websocket.TextMessage, jsonPingMessage); err != nil {
				fmt.Println("Failed to send ping:", err)
				return
			}
			fmt.Println("Ping sent with UTC time:", currentTime)

		case <-b.ctx.Done():
			fmt.Println("Ping context closed, stopping ping.")
			return
		}
	}
}

func (b *WebSocket) Disconnect() error {
	b.cancel()
	b.isConnected = false
	return b.conn.Close()
}

func (b *WebSocket) Send(message string) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (b *WebSocket) requiresAuthentication() bool {
	return b.url == WEBSOCKET_PRIVATE_MAINNET ||
		b.url == WEBSOCKET_PRIVATE_TESTNET || b.url == WEBSOCKET_TRADE_MAINNET || b.url == WEBSOCKET_TRADE_TESTNET || b.url == WEBSOCKET_TRADE_DEMO || b.url == WEBSOCKET_PRIVATE_DEMO
	// v3 offline
	/*
		b.url == V3_CONTRACT_PRIVATE ||
			b.url == V3_UNIFIED_PRIVATE ||
			b.url == V3_SPOT_PRIVATE
	*/
}

func (b *WebSocket) sendAuth() error {
	// Get current Unix time in milliseconds
	expires := time.Now().UnixNano()/1e6 + 10000
	val := fmt.Sprintf("GET/realtime%d", expires)

	h := hmac.New(sha256.New, []byte(b.apiSecret))
	h.Write([]byte(val))

	// Convert to hexadecimal instead of base64
	signature := hex.EncodeToString(h.Sum(nil))
	fmt.Println("signature generated : " + signature)

	authMessage := map[string]interface{}{
		"req_id": uuid.New(),
		"op":     "auth",
		"args":   []interface{}{b.apiKey, expires, signature},
	}
	fmt.Println("auth args:", fmt.Sprintf("%v", authMessage["args"]))
	return b.SendAsJson(authMessage)
}

func (b *WebSocket) sendSubscription(args []string) error {
	subMessage := map[string]interface{}{
		"req_id": uuid.New(),
		"op":     "subscribe",
		"args":   args,
	}
	fmt.Println("subscribe msg:", fmt.Sprintf("%v", subMessage["args"]))
	return b.SendAsJson(subMessage)
}

func (b *WebSocket) SendAsJson(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return b.Send(string(data))
}
