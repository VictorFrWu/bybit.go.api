package bybit_connector

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type baseTestSuite struct {
	suite.Suite
	client    *mockedClient
	apiKey    string
	apiSecret string
	baseURL   string
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	s.apiKey = "dummyAPIKey"
	s.apiSecret = "dummyApiSecret"
	s.baseURL = "https://dummyapi.com"
	s.client = newMockedClient(s.apiKey, s.apiSecret, s.baseURL)
}

func (s *baseTestSuite) mockDo(data []byte, err error, statusCode ...int) {
	s.client.Client.do = s.client.do
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	s.client.On("do", anyHTTPRequest()).Return(newHTTPResponse(data, code), err)
}

func (s *baseTestSuite) assertDo() {
	s.client.AssertCalled(s.T(), "do", anyHTTPRequest())
}

func (s *baseTestSuite) assertReq(f func(r *request)) {
	s.client.assertReq = f
}

func (s *baseTestSuite) assertRequestEqual(e, a *request) {
	s.assertURLValuesEqual(e.query, a.query)
}

func (s *baseTestSuite) assertURLValuesEqual(e, a url.Values) {
	var eKeys, aKeys []string
	for k := range e {
		eKeys = append(eKeys, k)
	}
	for k := range a {
		aKeys = append(aKeys, k)
	}
	r := s.r()
	r.Len(aKeys, len(eKeys))
	for k := range a {
		switch k {
		case timestampKey, signatureKey:
			r.NotEmpty(a.Get(k))
			continue
		}
		r.Equal(e[k], a[k], k)
	}
}

func anythingOfType(t string) mock.AnythingOfTypeArgument {
	return mock.AnythingOfType(t)
}

func newContext() context.Context {
	return context.Background()
}

func anyHTTPRequest() mock.AnythingOfTypeArgument {
	return anythingOfType("*http.Request")
}

func newHTTPResponse(data []byte, statusCode int) *http.Response {
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(data)),
		StatusCode: statusCode,
	}
}

func newRequest() *request {
	r := &request{
		query: url.Values{},
	}
	return r
}

func newSignedRequest() *request {
	return newRequest().setParams(params{
		timestampKey:  "",
		signatureKey:  "",
		apiRequestKey: "",
		recvWindowKey: "5000",
		signTypeKey:   "2",
	})
}

type assertReqFunc func(r *request)

type mockedClient struct {
	mock.Mock
	*Client
	assertReq assertReqFunc
}

func newMockedClient(apiKey, apiSecret, baseURL string) *mockedClient {
	m := new(mockedClient)
	m.Client = NewBybitHttpClient(apiKey, apiSecret, WithBaseURL(baseURL))
	return m
}

func (m *mockedClient) do(req *http.Request) (*http.Response, error) {
	if m.assertReq != nil {
		r := newRequest()
		r.query = req.URL.Query()
		if req.Body != nil && req.ContentLength > 0 {
			bs := make([]byte, req.ContentLength)
			_, err := req.Body.Read(bs)
			if err != nil && err != io.EOF {
				return nil, err // Handle read error
			}
			_ = req.Body.Close() // Close the body if we have read from it
			r.body = bytes.NewBuffer(bs)
		}
		m.assertReq(r)
	}
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestFormatTimestamp(t *testing.T) {
	tm, _ := time.Parse("2006-01-02 15:04:05", "2018-06-01 01:01:01")
	assert.Equal(t, int64(1527814861000), FormatTimestamp(tm))
}
