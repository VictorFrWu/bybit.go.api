package bybit_connector

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	handlers "github.com/wuhewuhe/bybit.go.api/handlers"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// TimeInForceType define time in force type of order
type TimeInForceType string

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

type doFunc func(req *http.Request) (*http.Response, error)

type ClientOption func(*Client)

func WithDebug(debug bool) ClientOption {
	return func(c *Client) {
		c.Debug = debug
	}
}

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Bybit.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

type ServerResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo struct{}    `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

// NewBybitHttpClient NewClient Create client function for initialising new Bybit client
func NewBybitHttpClient(apiKey string, secretKey string, options ...ClientOption) *Client {
	c := &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    "https://api.bybit.com",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, Name, log.LstdFlags),
	}

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	if bodyString != "" {
		header.Set("Content-Type", "application/json")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeSigned {
		header.Set(signTypeKey, "2")
		header.Set(apiRequestKey, c.APIKey)
		header.Set(timestampKey, fmt.Sprintf("%d", currentTimestamp()))
		if r.recvWindow == "" {
			header.Set(recvWindowKey, "5000")
		} else {
			header.Set(recvWindowKey, r.recvWindow)
		}

		var signatureBase string
		if r.method == "GET" {
			signatureBase = fmt.Sprintf("%d%s%s%s", FormatTimestamp(time.Now()), c.APIKey, r.recvWindow, queryString)
		} else if r.method == "POST" {
			signatureBase = fmt.Sprintf("%d%s%s%s", FormatTimestamp(time.Now()), c.APIKey, r.recvWindow, bodyString)
		}

		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(signatureBase))
		if err != nil {
			return err
		}
		signatureValue := fmt.Sprintf("%x", mac.Sum(nil))
		header.Set(signatureKey, signatureValue)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)
	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		var (
			apiErr = new(handlers.APIError)
		)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

// NewServerTimeService Market Endpoints
func (c *Client) NewServerTimeService() *ServerTime {
	return &ServerTime{c: c}
}
