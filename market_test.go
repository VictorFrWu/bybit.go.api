package bybit_connector

import (
	"github.com/stretchr/testify/suite"
	"github.com/wuhewuhe/bybit.go.api/models"
	"testing"
)

type marketTestSuite struct {
	baseTestSuite
}

func TestMarket(t *testing.T) {
	suite.Run(t, new(marketTestSuite))
}

func (s *marketTestSuite) TestServerTime() {
	data := []byte(`{
						"retCode": 0,
						"retMsg": "OK",
						"result": {
							"timeSecond": "1699287551",
							"timeNano": "1699287551919622633"
						},
						"retExtInfo": {},
						"time": 1699287551919
					}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	serverTime, err := s.client.NewMarketInfoServiceNoParams().GetServerTime(newContext())

	e1 := &models.ServerTimeResult{
		TimeSecond: "1699287551",
		TimeNano:   "1699287551919622633",
	}
	s.r().NoError(err)
	s.assertServerTimeEqual(e1, serverTime)
}

func (s *marketTestSuite) assertServerTimeEqual(expected *models.ServerTimeResult, actual *ServerResponse) {
	// Assert that the actual result is not nil and is a map
	r := s.r()
	actualResult, ok := actual.Result.(map[string]interface{})
	if !ok {
		r.FailNow("Actual result is not a map", "Actual Result: %#v", actual.Result)
		return
	}

	// Cast the map to the expected struct type
	var actualStruct models.ServerTimeResult
	timeSecond, okSecond := actualResult["timeSecond"].(string)
	timeNano, okNano := actualResult["timeNano"].(string)
	if !okSecond || !okNano {
		r.FailNow("Failed to cast actual result fields to string", "Actual Result: %#v", actual.Result)
		return
	}
	actualStruct.TimeSecond = timeSecond
	actualStruct.TimeNano = timeNano

	// Compare the fields
	r.Equal(expected.TimeSecond, actualStruct.TimeSecond, "TimeSecond field is not equal")
	r.Equal(expected.TimeNano, actualStruct.TimeNano, "TimeNano field is not equal")
}
