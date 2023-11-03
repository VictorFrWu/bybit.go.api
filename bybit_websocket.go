package bybit_connector

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

type MessageHandler func(message string) error

func (b *WebSocket) handleIncomingMessages() {
	for {
		_, message, err := b.conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading:", err)
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

func (b *WebSocket) SetMessageHandler(handler MessageHandler) {
	b.onMessage = handler
}

type WebSocket struct {
	conn      *websocket.Conn
	url       string
	apiKey    string
	apiSecret string
	onMessage MessageHandler
}

func NewBybitPrivateWebSocket(url, apiKey, apiSecret string, handler MessageHandler) *WebSocket {
	return &WebSocket{
		url:       url,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		onMessage: handler,
	}
}

func NewBybitPublicWebSocket(url string, handler MessageHandler) *WebSocket {
	return &WebSocket{
		url:       url,
		onMessage: handler,
	}
}

func (b *WebSocket) Connect(args []string) error {
	var err error
	b.conn, _, err = websocket.DefaultDialer.Dial(b.url, nil)
	if err != nil {
		return err
	}

	if b.requiresAuthentication() {
		if err = b.sendAuth(); err != nil {
			return err
		}
	}

	go b.handleIncomingMessages()

	Ping(b)

	return b.sendSubscription(args)
}

func Ping(b *WebSocket) {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		defer ticker.Stop() // Ensure the ticker is stopped when this goroutine ends
		for {
			select {
			case <-ticker.C: // Wait until the ticker sends a signal
				if err := b.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					fmt.Println("Failed to send ping:", err)
				}
			}
		}
	}()
}

func (b *WebSocket) Disconnect() error {
	return b.conn.Close()
}

func (b *WebSocket) Send(message string) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (b *WebSocket) requiresAuthentication() bool {
	return b.url == WEBSOCKET_PRIVATE_MAINNET ||
		b.url == WEBSOCKET_PRIVATE_TESTNET ||
		b.url == V3_CONTRACT_PRIVATE ||
		b.url == V3_UNIFIED_PRIVATE ||
		b.url == V3_SPOT_PRIVATE
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
		"req_id": uuid.New(), // You would need to implement or find a package for generating GUID
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
