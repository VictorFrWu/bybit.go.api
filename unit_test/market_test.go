package bybit_connector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetServerTime(t *testing.T) {
	// Create a mock server that returns a fixed response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"time": "2024-05-20T12:34:56Z"}`))
		if err != nil {
			fmt.Sprint(err.Error())
			return
		}
	}))
	defer mockServer.Close()

	// Adjust the base URL of the BybitClientRequest to use the mock server
	client := BybitClientRequest{
		httpClient: http.DefaultClient,
		url:        mockServer.URL, // Use mock server URL
	}

	// Call the method
	response, err := client.GetServerTime(context.Background())
	if err != nil {
		t.Errorf("Error when calling GetServerTime: %v", err)
	}

	// Check if the response is as expected
	expectedTime := "2024-05-20T12:34:56Z"
	if response.Time != expectedTime {
		t.Errorf("Expected time %v, got %v", expectedTime, response.Time)
	}
}
