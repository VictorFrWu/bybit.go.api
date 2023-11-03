package handlers

import (
	"errors"
	"fmt"
)

// APIError define API error when response status is 4xx or 5xx
type APIError struct {
	Code    int64  `json:"retCode"`
	Message string `json:"retMsg"`
}

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	var APIError *APIError
	ok := errors.As(e, &APIError)
	return ok
}
